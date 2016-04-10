package flux // import "kego.io/editor/client/flux"

import "reflect"

func Changed(stores ...StoreInterface) chan struct{} {
	out := make(chan struct{})
	cases := make([]reflect.SelectCase, len(stores))
	for i, store := range stores {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(store.WatchAll())}
	}
	go func() {
		for {
			reflect.Select(cases)
			out <- struct{}{}
		}
	}()
	return out
}
