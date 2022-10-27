package core

type Res map[string]int64

type Calculate interface {
	Calculate(clusterInfo *ClusterInfo) []*Res
	Write(res *[]Res)
}
