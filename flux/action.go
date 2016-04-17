package flux

type ActionInterface interface{}

type Payload struct {
	Action  ActionInterface
	WaitFor func(stores ...StoreInterface)
	Done    chan struct{}
}
