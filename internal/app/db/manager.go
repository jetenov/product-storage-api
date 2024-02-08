package db

import (
	"context"
	"fmt"
	"os"

	"github.com/spaolacci/murmur3"

	"gitlab.dg.ru/platform/database-go/sql"
	"gitlab.dg.ru/platform/scratch/closer"
	"gitlab.dg.ru/platform/scratch/config/secret"
	"gitlab.dg.ru/platform/tracer-go/logger"
)

// ShardManager ...
type ShardManager struct {
	BucketsCount uint64
	Buckets      []Bucket
}

// Bucket ...
type Bucket struct {
	Rng []uint64
	DB  DB
}

// NewManager ...
func NewManager(ctx context.Context, config *Config, cl *closer.Closer, hc healthcheck.Handler) (*ShardManager, error) {
	shardsMap := make(map[string]sql.Balancer, len(config.Shards))
	for _, shard := range config.Shards {
		dsn := secret.GetValue(ctx, shard.DSN).String()
		dsn = os.ExpandEnv(dsn)

		bal, err := NewDB(ctx, dsn, cl, hc, WithMaxOpenConns(shard.MaxOpenConnections),
			WithMaxIdleConns(shard.MaxIdleConnections))
		if err != nil {
			return nil, err
		}

		shardsMap[shard.Name] = bal
	}

	var buckets []Bucket
	for _, desc := range config.Bucket.Desc {
		shard := shardsMap[desc.Shard]
		b := Bucket{
			Rng: desc.Range,
			DB: DB{
				Instance: shard,
				Name:     desc.Shard,
			},
		}
		buckets = append(buckets, b)
	}

	return &ShardManager{
		BucketsCount: config.Bucket.Count,
		Buckets:      buckets,
	}, nil
}

// ShardByName returns shard by key
func (m *ShardManager) ShardByName(name string) (*DB, error) {
	for i := range m.Buckets {
		if m.Buckets[i].DB.Name == name {
			return &m.Buckets[i].DB, nil
		}
	}

	return nil, fmt.Errorf("shard is not found for name %q", name)
}

// ShardByID returns shard by key
func (m *ShardManager) ShardByID(id string) (*DB, error) {
	bucketID := key(id) % m.BucketsCount

	for i := range m.Buckets {
		if bucketID >= m.Buckets[i].Rng[0] && bucketID <= m.Buckets[i].Rng[1] {
			return &m.Buckets[i].DB, nil
		}
	}

	return nil, fmt.Errorf("range is not found for bucket %d", bucketID)
}

// ShardPerIDs returns shard per keys
func (m *ShardManager) ShardPerIDs(ctx context.Context, ids []string) map[string][]string {
	shardPerItems := make(map[string][]string, len(m.Buckets))
	for _, id := range ids {
		shard, err := m.ShardByID(id)
		if err != nil {
			logger.Errorf(ctx, "can't find shard for id %s, err %s", id, err)
			continue
		}

		shardPerItems[shard.Name] = append(shardPerItems[shard.Name], id)
	}

	return shardPerItems
}

func key(id string) uint64 {
	return murmur3.Sum64([]byte(id))
}
