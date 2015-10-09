package system

import (
	"sort"
	"sync"

	"kego.io/kerr"
)

var sourceRegistry struct {
	sync.RWMutex
	m map[Reference]SourceHashed
}

type SourceHashed struct {
	Id     Reference
	Type   Reference
	Source []byte
	Hash   uint64
}

func RegisterSource(id Reference, typ Reference, source []byte, hash uint64) error {
	sourceRegistry.Lock()
	defer sourceRegistry.Unlock()
	if sourceRegistry.m == nil {
		sourceRegistry.m = make(map[Reference]SourceHashed)
	}
	if h, found := sourceRegistry.m[id]; found && h.Hash != hash {
		return kerr.New("ANBTDVMCYE", nil, "Source %s already exists", id.String())
	}
	sourceRegistry.m[id] = SourceHashed{id, typ, source, hash}
	return nil
}

func UnregisterSource(id Reference) {
	sourceRegistry.Lock()
	defer sourceRegistry.Unlock()
	if sourceRegistry.m == nil {
		return
	}
	delete(sourceRegistry.m, id)
}

func GetSource(id Reference) (hashed SourceHashed, found bool) {
	sourceRegistry.RLock()
	defer sourceRegistry.RUnlock()
	hashed, found = sourceRegistry.m[id]
	return
}

func GetAllSourceInPackage(path string) []SourceHashed {
	out := SortableSourceHashed{}
	sourceRegistry.RLock()
	defer sourceRegistry.RUnlock()
	for ref, h := range sourceRegistry.m {
		if ref.Package != path {
			continue
		}
		out = append(out, h)
	}
	sort.Sort(out)
	return []SourceHashed(out)
}

type SortableSourceHashed []SourceHashed

func (s SortableSourceHashed) Len() int {
	return len(s)
}
func (s SortableSourceHashed) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortableSourceHashed) Less(i, j int) bool {
	return s[i].Id.Value() < s[j].Id.Value()
}
