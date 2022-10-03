package producer

type Producer interface {
	Produce(event interface{}) error
}
