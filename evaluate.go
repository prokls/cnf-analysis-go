package main

import (
	"github.com/prokls/cnf-analysis-go/output"
	"github.com/prokls/cnf-analysis-go/sat"
	"github.com/prokls/cnf-analysis-go/stats"
)

// the evaluated features are classified by their memory consumption

func evaluateConstant(cnf *sat.CNF, feat *output.Features, fconf *stats.FeatureConfig) error {
	// evaluating features: {ClausesCount, DefiniteClausesCount, GoalClausesCount,
	//   FalseTrivial, LiteralsCount, NbClauses, NbVars, NegativeUnitClauseCount,
	//   PositiveLiteralsCount, PositiveUnitClauseCount, TautologicalLiteralsCount,
	//   TrueTrivial, TwoLiteralsClauseCount, VariablesLargest, VariablesSmallest}
	var pos uint16
	var neg uint16
	var length uint16
	var maxlength uint16

	feat.NbClauses = uint32(cnf.NbClauses)
	feat.NbVars = uint32(cnf.NbVars)
	feat.TrueTrivial = true
	feat.FalseTrivial = true
	initSmallest := false

	for _, lit := range cnf.Lits {
		if lit == 0 {
			if pos == 1 {
				feat.DefiniteClausesCount += 1
			} else if pos == 0 {
				feat.GoalClausesCount += 1
				feat.TrueTrivial = false
			}
			if neg == 0 {
				feat.FalseTrivial = false
			}
			if length == 1 && pos > 0 {
				feat.PositiveUnitClauseCount += 1
			}
			if length == 1 && neg > 0 {
				feat.NegativeUnitClauseCount += 1
			}
			if length == 2 {
				feat.TwoLiteralsClauseCount += 1
			}
			if length > maxlength {
				maxlength = length
			}
			feat.ClausesCount += 1
			length = 0
			pos = 0
			neg = 0
		} else if lit > 0 {
			pos += 1
			length += 1
			feat.LiteralsCount += 1
			feat.PositiveLiteralsCount += 1
			if uint32(lit) > feat.VariablesLargest {
				feat.VariablesLargest = uint32(lit)
			}
			if !initSmallest || uint32(lit) < feat.VariablesSmallest {
				feat.VariablesSmallest = uint32(lit)
				initSmallest = true
			}
		} else {
			neg += 1
			length += 1
			feat.LiteralsCount += 1
			if uint32(-lit) > feat.VariablesLargest {
				feat.VariablesLargest = uint32(-lit)
			}
			if !initSmallest || uint32(-lit) < feat.VariablesSmallest {
				feat.VariablesSmallest = uint32(-lit)
				initSmallest = true
			}
		}
	}

	clause := make([]sat.Lit, maxlength)
	ci := 0
	skip := false

	for _, lit := range cnf.Lits {
		if skip && lit != 0 {
			continue
		}
		if lit == 0 {
			// reset
			for i := uint16(0); i < maxlength; i++ {
				clause[i] = 0
			}
			ci = 0
			skip = false
		} else {
			// does clause contain -lit?
			for i := 0; i < ci; i++ {
				if clause[i] == -lit {
					feat.TautologicalLiteralsCount += 1
					skip = true
				}
			}
			// store
			clause[ci] = lit
			ci += 1
		}
	}

	return nil
}

func posEquiv(lit int32, nbvars int) int32 {
	if lit < 0 {
		return int32(nbvars) + (-lit) - 1
	} else {
		return lit - 1
	}
}
func negEquiv(lit int32, nbvars int) int32 {
	if lit >= int32(nbvars) {
		return -(lit + 1 - int32(nbvars))
	} else {
		return lit + 1
	}
}

func evaluateOccurence(cnf *sat.CNF, feat *output.Features, fconf *stats.FeatureConfig) error {
	var err error
	freq := make([]float32, 2*cnf.NbVars)
	lowLit := int32(-cnf.NbVars)
	lowVar := int32(1)
	high := int32(cnf.NbVars)

	// retrieve occurence list
	for _, lit := range cnf.Lits {
		if lit != 0 {
			freq[posEquiv(int32(lit), cnf.NbVars)] += 1
		}
	}

	// determine existential literals
	for lit := lowLit; lit <= high; lit++ {
		if lit == 0 {
			continue
		}
		if freq[posEquiv(lit, cnf.NbVars)] == 1 && freq[posEquiv(-lit, cnf.NbVars)] == 0 {
			feat.ExistentialLiteralsCount += 1
			if lit > 0 {
				feat.ExistentialPositiveLiteralsCount += 1
			}
		}
	}

	// count variables used
	for lit := lowVar; lit <= high; lit++ {
		if lit == 0 {
			continue
		}
		if freq[posEquiv(lit, cnf.NbVars)] > 0.5 || freq[posEquiv(-lit, cnf.NbVars)] > 0.5 {
			feat.VariablesUsedCount += 1
		}
	}

	// frequency = occurences / nbclauses
	nbc32 := float32(cnf.NbClauses)
	for lit := lowLit; lit <= high; lit++ {
		if lit == 0 {
			continue
		}
		index := posEquiv(lit, cnf.NbVars)
		if 0.5 < freq[index] && freq[index] < 1.5 {
			feat.LiteralsOccurenceOneCount += 1
		}
		freq[index] /= nbc32
		if freq[index] >= 1.0 {
			freq[index] = 1.0
		}
	}

	// write frequency
	var freqFeats [20]*uint32 = [20]*uint32{
		&feat.LiteralsFrequency0To5, &feat.LiteralsFrequency5To10,
		&feat.LiteralsFrequency10To15, &feat.LiteralsFrequency15To20,
		&feat.LiteralsFrequency20To25, &feat.LiteralsFrequency25To30,
		&feat.LiteralsFrequency30To35, &feat.LiteralsFrequency35To40,
		&feat.LiteralsFrequency40To45, &feat.LiteralsFrequency45To50,
		&feat.LiteralsFrequency50To55, &feat.LiteralsFrequency55To60,
		&feat.LiteralsFrequency60To65, &feat.LiteralsFrequency65To70,
		&feat.LiteralsFrequency70To75, &feat.LiteralsFrequency75To80,
		&feat.LiteralsFrequency80To85, &feat.LiteralsFrequency85To90,
		&feat.LiteralsFrequency90To95, &feat.LiteralsFrequency95To100,
	}
	for lit := lowLit; lit <= high; lit++ {
		if lit == 0 {
			continue
		}
		class := int((100.0 * freq[posEquiv(lit, cnf.NbVars)]) / 5.0)
		if class == 20 {
			class = 19
		}
		*freqFeats[class] += 1
	}

	// min, max, mean, median, sd, entropy
	mean, err := stats.MeanFloat32(freq)
	if err != nil {
		return err
	}
	feat.LiteralsFrequencyEntropy, err = stats.EntropyFloat32(freq)
	if err != nil {
		return err
	}
	largest, err := stats.LargestFloat32(freq)
	if err != nil {
		return err
	}
	feat.LiteralsFrequencyLargest = float64(largest)
	feat.LiteralsFrequencyMean = mean
	feat.LiteralsFrequencyMedian, err = stats.MedianFloat32(freq)
	if err != nil {
		return err
	}
	feat.LiteralsFrequencySd, err = stats.StdevFloat32(freq, mean)
	if err != nil {
		return err
	}
	smallest, err := stats.SmallestFloat32(freq)
	if err != nil {
		return err
	}
	feat.LiteralsFrequencySmallest = float64(smallest)

	// frequency = occurences / nbclauses of variables
	var start, end int
	var initStart bool
	for lit := lowVar; lit <= high; lit++ {
		if lit == 0 {
			continue
		}
		p := posEquiv(lit, cnf.NbVars)
		n := posEquiv(-lit, cnf.NbVars)
		freq[p] = freq[p] + freq[n]
		if freq[p] > 1.0 {
			freq[p] = 1.0
		}
		if p > int32(end) {
			end = int(p)
		}
		if !initStart || p < int32(start) {
			start = int(p)
			initStart = true
		}
	}

	// write frequency
	freqFeats = [20]*uint32{
		&feat.VariablesFrequency0To5, &feat.VariablesFrequency5To10,
		&feat.VariablesFrequency10To15, &feat.VariablesFrequency15To20,
		&feat.VariablesFrequency20To25, &feat.VariablesFrequency25To30,
		&feat.VariablesFrequency30To35, &feat.VariablesFrequency35To40,
		&feat.VariablesFrequency40To45, &feat.VariablesFrequency45To50,
		&feat.VariablesFrequency50To55, &feat.VariablesFrequency55To60,
		&feat.VariablesFrequency60To65, &feat.VariablesFrequency65To70,
		&feat.VariablesFrequency70To75, &feat.VariablesFrequency75To80,
		&feat.VariablesFrequency80To85, &feat.VariablesFrequency85To90,
		&feat.VariablesFrequency90To95, &feat.VariablesFrequency95To100,
	}
	for lit := lowVar; lit <= high; lit++ {
		if lit == 0 {
			continue
		}
		class := int(20.0 * freq[posEquiv(int32(lit), cnf.NbVars)])
		if class == 20.0 {
			class = 19
		}
		*freqFeats[class] += 1
	}

	// min, max, mean, median, sd, entropy
	f := freq[start : end+1]
	mean, err = stats.MeanFloat32(f)
	if err != nil {
		return err
	}
	feat.VariablesFrequencyEntropy, err = stats.EntropyFloat32(f)
	if err != nil {
		return err
	}
	large, err := stats.LargestFloat32(f)
	if err != nil {
		return err
	}
	feat.VariablesFrequencyLargest = float64(large)
	feat.VariablesFrequencyMean = mean
	feat.VariablesFrequencyMedian, err = stats.MedianFloat32(f)
	if err != nil {
		return err
	}
	feat.VariablesFrequencySd, err = stats.StdevFloat32(f, mean)
	if err != nil {
		return err
	}
	small, err := stats.SmallestFloat32(f)
	if err != nil {
		return err
	}
	feat.VariablesFrequencySmallest = float64(small)
	return nil
}

func evaluateVarSdPosNeg(cnf *sat.CNF, feat *output.Features, fconf *stats.FeatureConfig) error {
	var err error
	clause := make([]uint32, 0, 64)

	// standard deviation of each clause, mean in CNF
	data := make([]float32, 0, cnf.NbClauses)
	for _, lit := range cnf.Lits {
		if lit > 0 {
			clause = append(clause, uint32(lit))
		} else if lit < 0 {
			clause = append(clause, uint32(-lit))
		} else {
			mean, err := stats.MeanUint32(clause)
			if err != nil {
				return err
			}
			sd, err := stats.StdevUint32(clause, mean)
			if err != nil {
				return err
			}
			clause = clause[:0]
			data = append(data, float32(sd))
		}
	}
	feat.ClauseVariablesSdMean, err = stats.MeanFloat32(data)
	if err != nil {
		return err
	}

	// ratio of positive/negative literals per clause
	var pos, neg, i int
	for _, lit := range cnf.Lits {
		if lit == 0 {
			ratio := float64(pos) / float64(pos+neg)
			data[i] = float32(ratio)
			i += 1
			pos, neg = 0, 0
		} else if lit > 0 {
			pos += 1
		} else if lit < 0 {
			neg += 1
		}
	}
	mean, err := stats.MeanFloat32(data[:i])
	if err != nil {
		return err
	}
	feat.PositiveNegativeLiteralsInClauseRatioEntropy, err = stats.EntropyFloat32(data[:i])
	if err != nil {
		return err
	}
	feat.PositiveNegativeLiteralsInClauseRatioMean = mean
	feat.PositiveNegativeLiteralsInClauseRatioStdev, err = stats.StdevFloat32(data[:i], mean)
	if err != nil {
		return err
	}

	return nil
}

func evaluateClauseLengthPosNeg(cnf *sat.CNF, feat *output.Features, fconf *stats.FeatureConfig) error {
	var err error
	data := make([]uint16, cnf.NbClauses)
	i := 0

	// determine length
	var length uint16
	for _, lit := range cnf.Lits {
		if lit == 0 {
			data[i] = length
			i += 1
			length = 0
		} else {
			length += 1
		}
	}

	// store length features
	feat.ClausesLengthLargest, err = stats.LargestUint16(data)
	if err != nil {
		return err
	}
	feat.ClausesLengthMean, err = stats.MeanUint16(data)
	if err != nil {
		return err
	}
	feat.ClausesLengthMedian, err = stats.MedianUint16(data)
	if err != nil {
		return err
	}
	feat.ClausesLengthSd, err = stats.StdevUint16(data, feat.ClausesLengthMean)
	if err != nil {
		return err
	}
	feat.ClausesLengthSmallest, err = stats.SmallestUint16(data)
	if err != nil {
		return err
	}

	// determine neg literals
	var neg uint16
	i = 0
	for _, lit := range cnf.Lits {
		if lit == 0 {
			data[i] = neg
			i += 1
			neg = 0
		} else if lit < 0 {
			neg += 1
		}
	}

	// store neg-lits features
	feat.NegativeLiteralsInClauseLargest, err = stats.LargestUint16(data)
	if err != nil {
		return err
	}
	feat.NegativeLiteralsInClauseMean, err = stats.MeanUint16(data)
	if err != nil {
		return err
	}
	feat.NegativeLiteralsInClauseSmallest, err = stats.SmallestUint16(data)
	if err != nil {
		return err
	}

	// determine pos literals
	var pos uint16
	i = 0
	for _, lit := range cnf.Lits {
		if lit == 0 {
			data[i] = pos
			i += 1
			pos = 0
		} else if lit > 0 {
			pos += 1
		}
	}

	// store pos-lits features
	mean, err := stats.MeanUint16(data)
	if err != nil {
		return err
	}
	feat.PositiveLiteralsInClauseLargest, err = stats.LargestUint16(data)
	if err != nil {
		return err
	}
	feat.PositiveLiteralsInClauseMean = mean
	med, err := stats.MedianUint16(data)
	if err != nil {
		return err
	}
	feat.PositiveLiteralsInClauseMedian = float32(med)
	feat.PositiveLiteralsInClauseSd, err = stats.StdevUint16(data, mean)
	if err != nil {
		return err
	}
	feat.PositiveLiteralsInClauseSmallest, err = stats.SmallestUint16(data)
	if err != nil {
		return err
	}

	return nil
}

func evaluate(cnf *sat.CNF, feat *output.Features, fconf *stats.FeatureConfig) error {
	var err error

	err = evaluateConstant(cnf, feat, fconf)
	if err != nil {
		return err
	}

	err = evaluateOccurence(cnf, feat, fconf)
	if err != nil {
		return err
	}

	err = evaluateVarSdPosNeg(cnf, feat, fconf)
	if err != nil {
		return err
	}

	err = evaluateClauseLengthPosNeg(cnf, feat, fconf)
	if err != nil {
		return err
	}

	err = stats.EvaluateComponents(cnf, feat, fconf)
	if err != nil {
		return err
	}

	return nil
}
