package core

type Res struct {
	result map[string]int64
}

type Calculate interface {
	Calculate(sharingInfo *ClusterInfo) []*Res
}
