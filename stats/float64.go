package stats

import (
	"fmt"
	"math"
	"sort"
)

const BucketSize = 512

// MeanFloat64 computes the mean value of float64 elements.
// It uses a bucket size strategy (divide and conquer) to avoid
// overflowing or underflowing values for float64. Therefore
// it builds buckets of BucketSize and computes the mean value
// in each of them.
func MeanFloat64(x []float64) (float64, error) {
	// special cases
	if len(x) == 0 {
		return 0.0, fmt.Errorf("Cannot determine mean of 0 elements")
	}
	if len(x) == 1 {
		return x[0], nil
	}

	// setup
	var tmp float64
	buckets := len(x) / BucketSize
	trailer := len(x) % BucketSize
	data := make([]float64, 1+buckets)

	// divide
	for b := 0; b < buckets; b++ {
		tmp = 0.0
		for i := 0; i < BucketSize; i++ {
			tmp += x[b*BucketSize+i] / BucketSize
		}
		data[b] = tmp * BucketSize
	}
	for i := 0; i < trailer; i++ {
		data[buckets] += x[buckets*BucketSize+i] / float64(trailer)
	}
	data[buckets] = data[buckets] * float64(trailer)

	// conquer
	var result float64
	for b := 0; b < buckets+1; b++ {
		result += data[b]
	}
	result /= float64(len(x))

	return result, nil
}

// LargestFloat64 computes the maximum value the given elements.
// NaN values will be omitted. Inf values will be considered.
func LargestFloat64(x []float64) (float64, error) {
	// special cases
	if len(x) == 0 {
		return 0.0, fmt.Errorf("Cannot determine largest value of 0 elements")
	}
	if len(x) == 1 {
		return x[0], nil
	}

	largest := x[0]
	for _, val := range x {
		if val > largest {
			largest = val
		}
	}

	return largest, nil
}

// SmallestFloat64 computes the maximum value the given elements.
// NaN values will be omitted. Inf values will be considered.
func SmallestFloat64(x []float64) (float64, error) {
	// special cases
	if len(x) == 0 {
		return 0.0, fmt.Errorf("Cannot determine largest value of 0 elements")
	}
	if len(x) == 1 {
		return x[0], nil
	}

	smallest := x[0]
	for _, val := range x {
		if val < smallest {
			smallest = val
		}
	}

	return smallest, nil
}

// MedianFloat64 computes the median value of the given elements.
// It copies the parameter, sorts it and determines the median.
func MedianFloat64(y []float64) (float64, error) {
	if len(y) == 0 {
		return 0.0, fmt.Errorf("Cannot determine median of 0 elements")
	}
	x := make([]float64, len(y))
	copy(x, y)

	// sorting in O(n * log n)
	sort.Float64s(x)

	// element selection
	mid := int(len(x) / 2)
	if len(x)%2 == 1 {
		return x[mid], nil
	} else {
		return (x[mid] + x[mid-1]) / 2.0, nil
	}
}

// StdevFloat64 computes the population standard deviation of given elements
// and the mean must be provided as argument. Use MeanFloat64 if unknown.
func StdevFloat64(x []float64, mean float64) (float64, error) {
	var tmp float64
	for _, val := range x {
		tmp += (val - mean) * (val - mean)
	}

	return math.Sqrt(tmp), nil
}

// EntropyFloat64 computes the entropy of probabilities given as float64 slice.
func EntropyFloat64(x []float64) (float64, error) {
	var entropy float64
	for _, p := range x {
		if p > 0.0 {
			entropy += p * math.Log2(p)
		}
	}
	return -entropy, nil
}
