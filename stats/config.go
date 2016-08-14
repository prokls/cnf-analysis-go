package stats

type FeatureConfig struct {
	Hashes   bool
	FullPath bool
}

func NewFeatureConfig() *FeatureConfig {
	return new(FeatureConfig)
}
