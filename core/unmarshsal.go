package core

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
)

type HostList []Container
type InstanceList []Container
type Containers []Container
type AppNameToInstanceList map[string][]int64
type AppNameNotToInstanceList map[string][]int64

type ClusterInfo struct {
	Hosts                 HostList
	Instances             InstanceList
	AppNameToInstances    AppNameToInstanceList
	AppNameNotToInstances AppNameNotToInstanceList
}

const (
	INSTANCE_TYPE    = "instance"
	HOST_TYPE        = "host"
	LOCAL_READ_DIR   = "dev.jsonl"
	REMOTE_READ_DIR  = "/home/admin/workspace/job/input/test.jsonl"
	LOCAL_WRITE_DIR  = "result.jsonl"
	REMOTE_WRITE_DIR = "/home/admin/workspace/job/output/data/results.jsonl"
)

type Container struct {
	ContainerType               string  `json:"type"`
	HostId                      int64   `json:"host_id"`
	HostCpu                     int64   `json:"host_cpu"`
	HostMemory                  int64   `json:"host_memory"`
	HostDisk                    int64   `json:"host_disk"`
	InstanceId                  int64   `json:"instance_id"`
	InstanceAppName             string  `json:"instance_app_name"`
	InstanceCpu0                float64 `json:"instance_cpu_0"`
	InstanceMemory0             float64 `json:"instance_memory_0"`
	InstanceDisk0               float64 `json:"instance_disk_0"`
	InstanceCpu1                float64 `json:"instance_cpu_1"`
	InstanceMemory1             float64 `json:"instance_memory_1"`
	InstanceDisk1               float64 `json:"instance_disk_1"`
	InstanceCpu2                float64 `json:"instance_cpu_2"`
	InstanceMemory2             float64 `json:"instance_memory_2"`
	InstanceDisk2               float64 `json:"instance_disk_2"`
	InstanceCpu3                float64 `json:"instance_cpu_3"`
	InstanceMemory3             float64 `json:"instance_memory_3"`
	InstanceDisk3               float64 `json:"instance_disk_3"`
	InstanceAntiAffinityAppName string  `json:"instance_anti_affinity_app_name"`
	Id                          int64   `json:"id"`
}

func initContainer() *Containers {
	fd, err := os.OpenFile(LOCAL_READ_DIR, os.O_RDONLY, 0755)
	defer fd.Close()
	if err != nil {
		panic(err)
	}
	containers := make(Containers, 0)
	br := bufio.NewReader(fd)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		var container Container
		err := json.Unmarshal(line[:len(line)], &container)
		if err != nil {
			panic(err)
		}
		containers = append(containers, container)
	}
	return &containers
}

func GetClusterInfo() *ClusterInfo {
	containers := initContainer()

	hosts := make(HostList, 0)
	instances := make(InstanceList, 0)
	appNameToInstanceList := make(map[string][]int64)
	appNameNotToInstanceList := make(map[string][]int64)

	for _, container := range *containers {
		if strings.EqualFold(container.ContainerType, INSTANCE_TYPE) {
			instances = append(instances, container)
		} else if strings.EqualFold(container.ContainerType, HOST_TYPE) {
			hosts = append(hosts, container)
		}
		appNameToInstanceList[container.InstanceAppName] = append(appNameToInstanceList[container.InstanceAppName], container.InstanceId)
		appNameNotToInstanceList[container.InstanceAntiAffinityAppName] = append(appNameNotToInstanceList[container.InstanceAntiAffinityAppName], container.InstanceId)
	}

	return &ClusterInfo{
		Hosts:                 hosts,
		Instances:             instances,
		AppNameToInstances:    appNameToInstanceList,
		AppNameNotToInstances: appNameNotToInstanceList,
	}
}

func WriteClusterInfo(res *[]Res) {
	fd, err := os.OpenFile(LOCAL_WRITE_DIR, os.O_CREATE|os.O_WRONLY, 0755)
	defer fd.Close()
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(fd)
	for _, r := range *res {
		err = encoder.Encode(r)
	}
	if err != nil {
		panic(err)
	}
}
