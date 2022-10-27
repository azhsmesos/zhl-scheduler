package main

import (
	"fmt"
	"zhl-scheduler/core"
)

func main() {
	Load()
}

func Load() {
	clusterInfo := core.GetClusterInfo()
	fmt.Println(len(clusterInfo.Hosts))
	fmt.Println(len(clusterInfo.Instances))
}
