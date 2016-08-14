package stats

import (
	"fmt"
	"math"
)

// MeanUint32 computes the mean value of uint32 elements.
// It uses a bucket size strategy (divide and conquer) to avoid
// overflowing or underflowing values for float64. Therefore
// it builds buckets of BucketSize and computes the mean value
// in each of them.
func MeanUint32(x []uint32) (float64, error) {
	// special cases
	if len(x) == 0 {
		return 0, fmt.Errorf("Cannot determine mean of 0 elements")
	}
	if len(x) == 1 {
		return float64(x[0]), nil
	}

	// setup
	var tmp float64
	buckets := len(x) / BucketSize
	trailer := len(x) % BucketSize
	data := make([]float32, 1+buckets)

	// divide
	for b := 0; b < buckets; b++ {
		tmp = 0.0
		for i := 0; i < BucketSize; i++ {
			tmp += float64(x[b*BucketSize+i]) / float64(BucketSize)
		}
		data[b] = float32(tmp) * BucketSize
	}
	for i := 0; i < trailer; i++ {
		data[buckets] += float32(x[buckets*BucketSize+i]) / float32(trailer)
	}
	data[buckets] = data[buckets] * float32(trailer)

	// conquer
	var result float64
	for b := 0; b < buckets+1; b++ {
		result += float64(data[b])
	}
	result /= float64(len(x))

	return result, nil
}

// StdevUint32 computes the population standard deviation of given elements
// and the mean must be provided as argument. Use MeanUint32 if unknown.
func StdevUint32(x []uint32, mean float64) (float64, error) {
	var tmp float64
	for _, val := range x {
		v := float64(val)
		tmp += (v - mean) * (v - mean)
	}

	factor := math.Sqrt(1.0 / float64(len(x)))
	return factor * math.Sqrt(tmp), nil
}
