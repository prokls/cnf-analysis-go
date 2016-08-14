package output

type Stats struct {
	CNFHash   string   `json:"@cnfhash"`
	Filename  string   `json:"@filename"`
	MD5Sum    string   `json:"@md5sum"`
	SHA1Sum   string   `json:"@sha1sum"`
	Timestamp string   `json:"@timestamp"`
	Version   string   `json:"@version"`
	Fts       Features `json:"featuring"`
}

func NewStats() *Stats {
	return new(Stats)
}

type Features struct {
	ClauseVariablesSdMean                        float64 `json:"clause_variables_sd_mean"`
	ClausesCount                                 uint32  `json:"clauses_count"`
	ClausesLengthLargest                         uint16  `json:"clauses_length_largest"`
	ClausesLengthMean                            float64 `json:"clauses_length_mean"`
	ClausesLengthMedian                          float64 `json:"clauses_length_median"`
	ClausesLengthSd                              float64 `json:"clauses_length_sd"`
	ClausesLengthSmallest                        uint16  `json:"clauses_length_smallest"`
	ConnectedLiteralComponentsCount              uint16  `json:"connected_literal_components_count"`
	ConnectedVariableComponentsCount             uint16  `json:"connected_variable_components_count"`
	DefiniteClausesCount                         uint32  `json:"definite_clauses_count"`
	ExistentialLiteralsCount                     uint32  `json:"existential_literals_count"`
	ExistentialPositiveLiteralsCount             uint32  `json:"existential_positive_literals_count"`
	FalseTrivial                                 bool    `json:"false_trivial"`
	GoalClausesCount                             uint32  `json:"goal_clauses_count"`
	LiteralsCount                                uint64  `json:"literals_count"`
	LiteralsFrequency0To5                        uint32  `json:"literals_frequency_0_to_5"`
	LiteralsFrequency5To10                       uint32  `json:"literals_frequency_5_to_10"`
	LiteralsFrequency10To15                      uint32  `json:"literals_frequency_10_to_15"`
	LiteralsFrequency15To20                      uint32  `json:"literals_frequency_15_to_20"`
	LiteralsFrequency20To25                      uint32  `json:"literals_frequency_20_to_25"`
	LiteralsFrequency25To30                      uint32  `json:"literals_frequency_25_to_30"`
	LiteralsFrequency30To35                      uint32  `json:"literals_frequency_30_to_35"`
	LiteralsFrequency35To40                      uint32  `json:"literals_frequency_35_to_40"`
	LiteralsFrequency40To45                      uint32  `json:"literals_frequency_40_to_45"`
	LiteralsFrequency45To50                      uint32  `json:"literals_frequency_45_to_50"`
	LiteralsFrequency50To55                      uint32  `json:"literals_frequency_50_to_55"`
	LiteralsFrequency55To60                      uint32  `json:"literals_frequency_55_to_60"`
	LiteralsFrequency60To65                      uint32  `json:"literals_frequency_60_to_65"`
	LiteralsFrequency65To70                      uint32  `json:"literals_frequency_65_to_70"`
	LiteralsFrequency70To75                      uint32  `json:"literals_frequency_70_to_75"`
	LiteralsFrequency75To80                      uint32  `json:"literals_frequency_75_to_80"`
	LiteralsFrequency80To85                      uint32  `json:"literals_frequency_80_to_85"`
	LiteralsFrequency85To90                      uint32  `json:"literals_frequency_85_to_90"`
	LiteralsFrequency90To95                      uint32  `json:"literals_frequency_90_to_95"`
	LiteralsFrequency95To100                     uint32  `json:"literals_frequency_95_to_100"`
	LiteralsFrequencyEntropy                     float64 `json:"literals_frequency_entropy"`
	LiteralsFrequencyLargest                     float64 `json:"literals_frequency_largest"`
	LiteralsFrequencyMean                        float64 `json:"literals_frequency_mean"`
	LiteralsFrequencyMedian                      float64 `json:"literals_frequency_median"`
	LiteralsFrequencySd                          float64 `json:"literals_frequency_sd"`
	LiteralsFrequencySmallest                    float64 `json:"literals_frequency_smallest"`
	LiteralsOccurenceOneCount                    uint64  `json:"literals_occurence_one_count"`
	NbClauses                                    uint32  `json:"nbclauses"`
	NbVars                                       uint32  `json:"nbvars"`
	NegativeLiteralsInClauseLargest              uint16  `json:"negative_literals_in_clause_largest"`
	NegativeLiteralsInClauseMean                 float64 `json:"negative_literals_in_clause_mean"`
	NegativeLiteralsInClauseSmallest             uint16  `json:"negative_literals_in_clause_smallest"`
	NegativeUnitClauseCount                      uint32  `json:"negative_unit_clause_count"`
	PositiveLiteralsCount                        uint32  `json:"positive_literals_count"`
	PositiveLiteralsInClauseLargest              uint16  `json:"positive_literals_in_clause_largest"`
	PositiveLiteralsInClauseMean                 float64 `json:"positive_literals_in_clause_mean"`
	PositiveLiteralsInClauseMedian               float32 `json:"positive_literals_in_clause_median"`
	PositiveLiteralsInClauseSd                   float64 `json:"positive_literals_in_clause_sd"`
	PositiveLiteralsInClauseSmallest             uint16  `json:"positive_literals_in_clause_smallest"`
	PositiveNegativeLiteralsInClauseRatioEntropy float64 `json:"positive_negative_literals_in_clause_ratio_entropy"`
	PositiveNegativeLiteralsInClauseRatioStdev   float64 `json:"positive_negative_literals_in_clause_ratio_stdev"`
	PositiveNegativeLiteralsInClauseRatioMean    float64 `json:"positive_negative_literals_in_clause_ratio_mean"`
	PositiveUnitClauseCount                      uint32  `json:"positive_unit_clause_count"`
	TautologicalLiteralsCount                    uint16  `json:"tautological_literals_count"`
	TrueTrivial                                  bool    `json:"true_trivial"`
	TwoLiteralsClauseCount                       uint32  `json:"two_literals_clause_count"`
	VariablesFrequency0To5                       uint32  `json:"variables_frequency_0_to_5"`
	VariablesFrequency5To10                      uint32  `json:"variables_frequency_5_to_10"`
	VariablesFrequency10To15                     uint32  `json:"variables_frequency_10_to_15"`
	VariablesFrequency15To20                     uint32  `json:"variables_frequency_15_to_20"`
	VariablesFrequency20To25                     uint32  `json:"variables_frequency_20_to_25"`
	VariablesFrequency25To30                     uint32  `json:"variables_frequency_25_to_30"`
	VariablesFrequency30To35                     uint32  `json:"variables_frequency_30_to_35"`
	VariablesFrequency35To40                     uint32  `json:"variables_frequency_35_to_40"`
	VariablesFrequency40To45                     uint32  `json:"variables_frequency_40_to_45"`
	VariablesFrequency45To50                     uint32  `json:"variables_frequency_45_to_50"`
	VariablesFrequency50To55                     uint32  `json:"variables_frequency_50_to_55"`
	VariablesFrequency55To60                     uint32  `json:"variables_frequency_55_to_60"`
	VariablesFrequency60To65                     uint32  `json:"variables_frequency_60_to_65"`
	VariablesFrequency65To70                     uint32  `json:"variables_frequency_65_to_70"`
	VariablesFrequency70To75                     uint32  `json:"variables_frequency_70_to_75"`
	VariablesFrequency75To80                     uint32  `json:"variables_frequency_75_to_80"`
	VariablesFrequency80To85                     uint32  `json:"variables_frequency_80_to_85"`
	VariablesFrequency85To90                     uint32  `json:"variables_frequency_85_to_90"`
	VariablesFrequency90To95                     uint32  `json:"variables_frequency_90_to_95"`
	VariablesFrequency95To100                    uint32  `json:"variables_frequency_95_to_100"`
	VariablesFrequencyEntropy                    float64 `json:"variables_frequency_entropy"`
	VariablesFrequencyLargest                    float64 `json:"variables_frequency_largest"`
	VariablesFrequencyMean                       float64 `json:"variables_frequency_mean"`
	VariablesFrequencyMedian                     float64 `json:"variables_frequency_median"`
	VariablesFrequencySd                         float64 `json:"variables_frequency_sd"`
	VariablesFrequencySmallest                   float64 `json:"variables_frequency_smallest"`
	VariablesLargest                             uint32  `json:"variables_largest"`
	VariablesSmallest                            uint32  `json:"variables_smallest"`
	VariablesUsedCount                           uint32  `json:"variables_used_count"`
}

func NewFeatures() *Features {
	return new(Features)
}
