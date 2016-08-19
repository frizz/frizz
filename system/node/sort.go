package node

import (
	"sort"

	"github.com/davelondon/sorter"
)

func SortNodeMap(in map[string]*Node) (out []*Node) {
	names := map[*Node]string{}
	for name, node := range in {
		out = append(out, node)
		names[node] = name
	}
	sort.Sort(sorter.New(
		len(out),
		func(i, j int) { out[i], out[j] = out[j], out[i] },
		func(i, j int) bool { return names[out[i]] < names[out[j]] },
	))
	return
}
