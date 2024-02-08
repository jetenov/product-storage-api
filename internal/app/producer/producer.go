package producer

import (
	"context"
	"strconv"

	"github.com/Shopify/sarama"
	"google.golang.org/protobuf/proto"

	"gitlab.dg.ru/platform/kafka-client-go/pkg/kafka"
	"gitlab.dg.ru/platform/kafka-client-go/pkg/kafka/client"
	"gitlab.dg.ru/platform/kafka-client-go/pkg/kafka/middleware"
	"gitlab.dg.ru/platform/tracer-go/logger"
)

// KafkaProducer ...
type KafkaProducer interface {
	SendMessage(ctx context.Context, topic, key, message string) error
	SendProtoMessageHash(ctx context.Context, topic string, key int, value proto.Message) error
}

// Producer ...
type Producer struct {
	producer kafka.SyncProducer
}

// NewProducer ...
func NewProducer(ctx context.Context, brokerAddrs []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 10
	config.Version = sarama.V1_1_1_0

	producer, err := client.NewSyncProducer(brokerAddrs, config,
		middleware.SyncProducerMetricMiddleware, middleware.SyncProducerTraceMiddleware)
	if err != nil {
		logger.ErrorKV(ctx, "initializing kafka producer error", "err", err)
		return nil, err
	}

	return &Producer{producer}, nil
}

// SendMessage ...
func (p *Producer) SendMessage(ctx context.Context, topic, key, message string) error {
	producerMessage := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(message),
	}

	_, _, err := p.producer.SendMessageWithContext(ctx, producerMessage)
	return err
}

// SendProtoMessageHash ...
func (p *Producer) SendProtoMessageHash(ctx context.Context, topic string, key int, value proto.Message) error {
	bytesMessage, err := proto.Marshal(value)
	if err != nil {
		//logger.ErrorKV(ctx, "Failed to proto.Marshal message for kafka", "message", value, "topic", topic, "key", key, "err", err)
		return err
	}

	producerMessage := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(strconv.Itoa(key)),
		Value: sarama.ByteEncoder(bytesMessage),
	}

	_, _, err = p.producer.SendMessageWithContext(ctx, producerMessage)
	return err
}

// Close ...
func (p *Producer) Close() error {
	return p.producer.Close()
}
