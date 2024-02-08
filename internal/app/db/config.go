package db

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Config ...
type Config struct {
	Shards []InstanceConfig `json:"shards"`
	Bucket BucketConfig     `json:"buckets"`
}

// InstanceConfig ...
type InstanceConfig struct {
	Name               string `json:"name"`
	DSN                string `json:"dsn"`
	MaxOpenConnections int    `json:"max_open_connections"`
	MaxIdleConnections int    `json:"max_idle_connections"`
}

// BucketConfig ...
type BucketConfig struct {
	Count uint64              `json:"count"`
	Desc  []BucketsDescConfig `json:"desc"`
}

// BucketsDescConfig ...
type BucketsDescConfig struct {
	Range []uint64 `json:"range"`
	Shard string   `json:"shard"`
}

// ParseConfig ...
func ParseConfig(data []byte) (*Config, error) {
	if len(data) == 0 {
		return nil, errors.New("config doesn't set")
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("can't unmarshal config: %w", err)
	}

	if len(config.Shards) == 0 {
		return nil, errors.New("number of shards can't be 0")
	}

	if config.Bucket.Count < 2 {
		return nil, errors.New("buckets count should be at least 2")
	}

	if len(config.Bucket.Desc) == 0 {
		return nil, errors.New("buckets description list is empty")
	}

	for i, desc := range config.Bucket.Desc {
		if len(desc.Range) != 2 {
			return nil, errors.New("unknown range configuration")
		}

		if desc.Range[0] >= desc.Range[1] {
			return nil, errors.New("first value of the range is less/eq than the second")
		}

		if i == 0 && desc.Range[0] != 0 {
			return nil, errors.New("range doesn't start from zero")
		}

		if i == len(config.Bucket.Desc)-1 && desc.Range[1] != config.Bucket.Count-1 {
			return nil, errors.New("range doesn't cover all buckets")
		}

		if i != 0 {
			if desc.Range[0] != (config.Bucket.Desc[i-1].Range[1] + 1) {
				return nil, errors.New("range doesn't cover all Buckets")
			}
		}
		if desc.Shard == "" {
			return nil, errors.New("shard name is empty")
		}
	}

	return &config, nil
}
