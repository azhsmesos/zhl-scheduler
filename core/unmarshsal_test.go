package core

import "testing"

func TestWriteClusterInfo(t *testing.T) {
	res := make(map[string]int64)
	res2 := make(map[string]int64)
	res["host_id"] = 159
	res["instance_id"] = 83383
	res2["host_id"] = 159
	res2["instance_id"] = 39134
	WriteClusterInfo(&[]Res{
		res,
		res2,
	})
}
