package system

/*
func GetAllTypesThatImplementInterface(typ *Type) []Hashed {
	registry.RLock()
	defer registry.RUnlock()
	out := SortableHashed{}

	var reflectType reflect.Type
	if typ.Interface {
		// The type provided is an interface type
		rt, ok := typ.Id.GetReflectType()
		if !ok {
			return nil
		}
		reflectType = rt
	} else {
		rt, ok := typ.Id.GetReflectInterface()
		if !ok {
			return nil
		}
		reflectType = rt
	}

	for _, h := range registry.m {
		if t, ok := h.Object.(*Type); ok && t.Implements(reflectType) {
			out = append(out, h)
		}
	}

	sort.Sort(out)
	return []Hashed(out)
}


type SortableHashed []Hashed

func (s SortableHashed) Len() int {
	return len(s)
}
func (s SortableHashed) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortableHashed) Less(i, j int) bool {
	return s[i].Object.GetObject(nil).Id.Value() < s[j].Object.GetObject(nil).Id.Value()
}
*/
