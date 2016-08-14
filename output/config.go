package output

const (
	JSONFormat int = iota
)

type OutputConfig struct {
	Format int
}

func NewOutputConfig() *OutputConfig {
	oc := new(OutputConfig)
	oc.Format = JSONFormat
	return oc
}
