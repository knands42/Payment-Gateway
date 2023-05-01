package broker

import "context"

type ProducerInterface interface {
	Publish(ctx context.Context, msg interface{}, key []byte, topic string) error
}
