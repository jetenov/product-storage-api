package loader

import (
	"context"
	"strings"
	"time"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/producer"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/config"
	kafkachecker "gitlab.dg.ru/platform/healthcheck-go/checker/kafka"
	"gitlab.dg.ru/platform/scratch/closer"
)

// InitKafkaHealthcheck ...
func InitKafkaHealthcheck(ctx context.Context, hc healthcheck.Handler) {
	hc.AddReadinessCheck(
		kafkachecker.CheckerName,
		kafkachecker.DialCheck(
			strings.Fields(config.GetValue(ctx, config.KafkaBrokerList).String()),
			1500*time.Millisecond,
		),
	)
}

// InitKafkaProducer ...
func InitKafkaProducer(ctx context.Context, cl *closer.Closer) (*producer.Producer, error) {
	brokers := strings.Fields(config.GetValue(ctx, config.KafkaBrokerList).String())
	pr, err := producer.NewProducer(ctx, brokers)
	if err != nil {
		return nil, err
	}

	if cl == nil {
		closer.Add(pr.Close)
	} else {
		cl.Add(pr.Close)
	}

	return pr, nil
}
