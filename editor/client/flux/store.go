package flux

type StoreInterface interface {
	Handle(payload *Payload) (finished bool)
}
