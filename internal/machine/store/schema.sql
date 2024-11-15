-- cluster table stores the key-value pairs of the cluster configuration.
CREATE TABLE cluster
(
    key   TEXT NOT NULL PRIMARY KEY,
    value ANY
);

-- machines table stores the basic information of the machines in the cluster.
CREATE TABLE machines
(
    id   TEXT NOT NULL PRIMARY KEY,
    name TEXT AS (json_extract(info, '$.name')),
    -- info is a JSON-serialized MachineInfo protobuf message.
    info TEXT NOT NULL DEFAULT '{}' CHECK (json_valid(info))
);

-- containers table stores the Uncloud-managed Docker containers created in the cluster.
CREATE TABLE containers
(
    id           TEXT NOT NULL PRIMARY KEY,
    -- container is a JSON-serialized Docker container.Summary struct.
    container    TEXT NOT NULL DEFAULT '{}' CHECK (json_valid(container)),
    machine_id   TEXT NOT NULL DEFAULT '',
    service_id   TEXT AS (json_extract(container, '$.Labels."uncloud.service.id"')),
    service_name TEXT AS (json_extract(container, '$.Labels."uncloud.service.name"'))
);

CREATE INDEX idx_machines_name ON machines (name);

CREATE INDEX idx_containers_machine_id ON containers (machine_id);
CREATE INDEX idx_containers_service_id ON containers (service_id);
CREATE INDEX idx_containers_service_name ON containers (service_name);
