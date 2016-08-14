package stats

import (
	"fmt"
	"math"
	"sort"
)

// MeanFloat32 computes the mean value of float32 elements.
// It uses a bucket size strategy (divide and conquer) to avoid
// overflowing or underflowing values for float32. Therefore
// it builds buckets of BucketSize and computes the mean value
// in each of them.
func MeanFloat32(x []float32) (float64, error) {
	// special cases
	if len(x) == 0 {
		return 0.0, fmt.Errorf("Cannot determine mean of 0 elements")
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
			tmp += float64(x[b*BucketSize+i]) / BucketSize
		}
		data[b] = float32(tmp * BucketSize)
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

// LargestFloat32 computes the maximum value the given elements.
// NaN values will be omitted. Inf values will be considered.
func LargestFloat32(x []float32) (float32, error) {
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

// SmallestFloat32 computes the maximum value the given elements.
// NaN values will be omitted. Inf values will be considered.
func SmallestFloat32(x []float32) (float32, error) {
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

type sortingFloat32 []float32

func (s sortingFloat32) Len() int {
	return len(s)
}

func (s sortingFloat32) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortingFloat32) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// MedianFloat32 computes the median value of the given elements.
// It copies the parameter, sorts it and determines the median.
func MedianFloat32(y []float32) (float64, error) {
	if len(y) == 0 {
		return 0.0, fmt.Errorf("Cannot determine median of 0 elements")
	}
	x := make([]float32, len(y))
	copy(x, y)

	// sorting in O(n * log n)
	sort.Sort(sortingFloat32(x))

	// element selection
	mid := int(len(x) / 2)
	if len(x)%2 == 1 {
		return float64(x[mid]), nil
	} else {
		return (float64(x[mid]) + float64(x[mid-1])) / 2.0, nil
	}
}

// StdevFloat32 computes the population standard deviation of given elements.
// and the mean must be provided as argument. Use MeanFloat32 if unknown.
func StdevFloat32(x []float32, mean float64) (float64, error) {
	var tmp float64

	for _, val := range x {
		v := float64(val)
		tmp += (v - mean) * (v - mean)
	}

	factor := math.Sqrt(1.0 / float64(len(x)))
	return factor * math.Sqrt(tmp), nil
}

// EntropyFloat32 computes the entropy of probabilities given as float32 slice.
func EntropyFloat32(x []float32) (float64, error) {
	var entropy float64
	for _, prob := range x {
		if prob > 0.0 {
			p := float64(prob)
			entropy += p * math.Log2(p)
		}
	}
	return -entropy, nil
}
