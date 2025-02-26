package client

import (
	"fmt"
	"github.com/distribution/reference"
	"strings"
	"uncloud/internal/api"
	"uncloud/internal/secret"
)

// ServiceSpecResolver transforms user-provided service specs into deployment-ready form.
type ServiceSpecResolver struct {
	ClusterDomain string
}

func NewServiceSpecResolver(clusterDomain string) *ServiceSpecResolver {
	return &ServiceSpecResolver{ClusterDomain: clusterDomain}
}

// Resolve transforms a service spec into its fully resolved form ready for deployment.
func (r *ServiceSpecResolver) Resolve(spec *api.ServiceSpec) error {
	if err := spec.Validate(); err != nil {
		return fmt.Errorf("invalid service spec: %w", err)
	}

	steps := []func(*api.ServiceSpec) error{
		r.resolveServiceName,
		r.expandIngressPorts,
	}

	for _, step := range steps {
		if err := step(spec); err != nil {
			return err
		}
	}

	return nil
}

func (r *ServiceSpecResolver) resolveServiceName(spec *api.ServiceSpec) error {
	if spec.Name != "" {
		return nil
	}

	// Generate a random service name from the image when not provided.
	img, err := reference.ParseDockerRef(spec.Container.Image)
	if err != nil {
		return fmt.Errorf("invalid image: %w", err)
	}
	// Get the image name without the repository and tag/digest parts.
	imageName := reference.FamiliarName(img)
	// Get the last part of the image name (path), e.g. "nginx" from "bitnami/nginx".
	if i := strings.LastIndex(imageName, "/"); i != -1 {
		imageName = imageName[i+1:]
	}
	// Append a random suffix to the image name to generate an optimistically unique service name.
	suffix, err := secret.RandomAlphaNumeric(4)
	if err != nil {
		return fmt.Errorf("generate random suffix: %w", err)
	}
	spec.Name = fmt.Sprintf("%s-%s", imageName, suffix)

	return nil
}

// expandIngressPorts processes ingress ports in a service spec by:
// 1. Setting a default hostname (service-name.cluster-domain) for ports without a hostname.
// 2. Duplicating a port with a cluster domain hostname for ports with external domains.
// This ensures every ingress port is accessible via the cluster domain, while preserving any custom domains specified
// by the user.
func (r *ServiceSpecResolver) expandIngressPorts(spec *api.ServiceSpec) error {
	for i, port := range spec.Ports {
		if port.Mode != api.PortModeIngress {
			continue
		}

		if port.Hostname == "" {
			if r.ClusterDomain == "" {
				return fmt.Errorf("cluster domain must be reserved to generate hostname for ingress port: %d/%s",
					port.ContainerPort, port.Protocol)
			}
			// Assign the default hostname (service-name.cluster-domain).
			spec.Ports[i].Hostname = fmt.Sprintf("%s.%s", spec.Name, r.ClusterDomain)
		} else {
			if r.ClusterDomain == "" {
				// When no cluster domain is reserved, use only the provided hostname.
				continue
			}

			if strings.HasSuffix(port.Hostname, "."+r.ClusterDomain) {
				// If the hostname is already a cluster subdomain, use as is.
				continue
			}
			// For external domains, duplicate the port with a service-name.cluster-domain hostname so the service
			// can be accessed via both hostnames.
			newPort := port
			newPort.Hostname = fmt.Sprintf("%s.%s", spec.Name, r.ClusterDomain)
			spec.Ports = append(spec.Ports, newPort)
		}
	}

	return nil
}
