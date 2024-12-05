package api

import (
	"encoding/json"
	"fmt"
	"uncloud/internal/machine/api/pb"
)

const (
	ServiceModeReplicated = "replicated"
	ServiceModeGlobal     = "global"
)

type ServiceSpec struct {
	Container ContainerSpec
	// Mode is the replication mode of the service. Default is ServiceModeReplicated if empty.
	Mode string
	Name string
}

type ContainerSpec struct {
	Command []string
	Image   string
	// Run a custom init inside the container. If nil, use the daemon's configured settings.
	Init *bool
}

type Service struct {
	ID         string
	Name       string
	Mode       string
	Containers []MachineContainer
}

type MachineContainer struct {
	MachineID string
	Container Container
}

func ServiceFromProto(s *pb.Service) (Service, error) {
	var err error
	containers := make([]MachineContainer, len(s.Containers))
	for i, sc := range s.Containers {
		containers[i], err = machineContainerFromProto(sc)
		if err != nil {
			return Service{}, err
		}
	}

	return Service{
		ID:         s.Id,
		Name:       s.Name,
		Mode:       s.Mode,
		Containers: containers,
	}, nil
}

func machineContainerFromProto(sc *pb.Service_Container) (MachineContainer, error) {
	var c Container
	if err := json.Unmarshal(sc.Container, &c); err != nil {
		return MachineContainer{}, fmt.Errorf("unmarshal container: %w", err)
	}

	return MachineContainer{
		MachineID: sc.MachineId,
		Container: c,
	}, nil
}
