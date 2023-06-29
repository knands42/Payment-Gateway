package kafka

import (
	"context"
	"github.com/caiofernandes00/payment-gateway/adapter/presenter"
	tracer_adapter "github.com/caiofernandes00/payment-gateway/adapter/trace"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	ConfigMap *ckafka.ConfigMap
	Presenter presenter.Presenter
	otel      tracer_adapter.TraceClosure
}

func NewKafkaProducer(configMap *ckafka.ConfigMap, presenter presenter.Presenter, otel tracer_adapter.TraceClosure) *Producer {
	return &Producer{
		ConfigMap: configMap,
		Presenter: presenter,
		otel:      otel,
	}
}

func (p *Producer) Publish(ctx context.Context, msg interface{}, key []byte, topic string) (err error) {
	var producer *ckafka.Producer
	ctx = p.otel(ctx, "kafka-new-producer", func(ctx context.Context) {
		producer, err = ckafka.NewProducer(p.ConfigMap)
		defer producer.Close()
	})
	if err != nil {
		return
	}

	ctx = p.otel(ctx, "kafka-producer-bind", func(ctx context.Context) {
		err = p.Presenter.Bind(msg)
	})
	if err != nil {
		return
	}

	var presenterMsg []byte
	ctx = p.otel(ctx, "kafka-producer-show", func(ctx context.Context) {
		presenterMsg, err = p.Presenter.Show()
	})
	if err != nil {
		return
	}

	ctx = p.otel(ctx, "kafka-producer-produce", func(ctx context.Context) {
		message := &ckafka.Message{
			TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
			Value:          presenterMsg,
			Key:            key,
		}

		err = producer.Produce(message, nil)
	})

	return
}
