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
	Source []byte
	Hash   uint64
}

func RegisterSource(path string, name string, source []byte, hash uint64) error {
	sourceRegistry.Lock()
	defer sourceRegistry.Unlock()
	if sourceRegistry.m == nil {
		sourceRegistry.m = make(map[Reference]SourceHashed)
	}
	r := NewReference(path, name)
	if h, found := sourceRegistry.m[r]; found && h.Hash != hash {
		return kerr.New("ANBTDVMCYE", nil, "Source %s already exists", r.String())
	}
	sourceRegistry.m[r] = SourceHashed{r, source, hash}
	return nil
}

func UnregisterSource(path string, name string) {
	sourceRegistry.Lock()
	defer sourceRegistry.Unlock()
	if sourceRegistry.m == nil {
		return
	}
	delete(sourceRegistry.m, NewReference(path, name))
}

func GetSource(path string, name string) (hashed SourceHashed, found bool) {
	sourceRegistry.RLock()
	defer sourceRegistry.RUnlock()
	hashed, found = sourceRegistry.m[NewReference(path, name)]
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
