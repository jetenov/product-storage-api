package loader

import (
	"context"
	"fmt"

	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/app/db"
	"gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/config"
	"gitlab.dg.ru/platform/scratch/closer"
)

// InitDBShardManager ...
func InitDBShardManager(ctx context.Context, cl *closer.Closer, hc healthcheck.Handler) (*db.ShardManager, error) {
	configData := config.GetValue(ctx, config.ShardConfig).String()
	shardingCfg, err := db.ParseConfig([]byte(configData))
	if err != nil {
		return nil, fmt.Errorf("unable to parse sharding config: %w", err)
	}

	manager, err := db.NewManager(ctx, shardingCfg, cl, hc)
	if err != nil {
		return nil, fmt.Errorf("unable to create sharding manager: %w", err)
	}

	return manager, nil
}
