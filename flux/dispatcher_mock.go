package flux

type DispatcherMock struct {
	Log []ActionInterface
}

func (d *DispatcherMock) Dispatch(action ActionInterface) chan struct{} {
	d.Log = append(d.Log, action)
	done := make(chan struct{}, 1)
	close(done)
	return done
}
