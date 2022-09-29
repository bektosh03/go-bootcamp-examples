package producer

type Producer interface {
	Produce(event RegisteredEvent) error
}
