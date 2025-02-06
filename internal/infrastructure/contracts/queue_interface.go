package contracts

type IQueue interface {
	Publish(queue string, message []byte) error
	Consume(queue string, worker func([]byte)) error
}
