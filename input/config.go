package input

type ParsingConfig struct {
	IgnoreLines    []string
	CheckNbVars    bool
	CheckNbClauses bool
}

func NewParsingConfig() *ParsingConfig {
	pc := new(ParsingConfig)
	pc.IgnoreLines = make([]string, 0, 2)
	pc.IgnoreLines = append(pc.IgnoreLines, "c")
	pc.IgnoreLines = append(pc.IgnoreLines, "%")
	return pc
}
