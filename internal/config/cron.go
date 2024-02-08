package config

// Strictly-typed config keys.
const (
	BatchSize       = configKey("batch_size")
	SaveBatchSize   = configKey("save_batch_size")
	CephImagePrefix = configKey("ceph_image_prefix")
)
