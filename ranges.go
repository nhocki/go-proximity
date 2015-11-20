package proximity

import (
	"sort"

	geohash "github.com/corsc/go-geohash"
)

const maxBitDepth = 52

// Int64arr is an array of int64 elements that can be sorted using the `sort`
// package.
type Int64arr []int64

func (a Int64arr) Len() int           { return len(a) }
func (a Int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int64arr) Less(i, j int) bool { return a[i] < a[j] }

// Search looks for element `x` in array `a` using binary search.
// Uses `sort.Search` to find the index where `x` is or could be inserted
// to keep the array sorted.
//
// The array must be sorted before using `Search`
func (a Int64arr) Search(x int64) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

// Include returns if `x` is included in the array `a` or not using binary search.
//
// The array must be sorted before using `Include`
func (a Int64arr) Include(x int64) bool {
	index := a.Search(x)
	return index < len(a) && a[index] == x
}

// intervalsFromRadius Get the intervals to look into given a point & a radius.
func intervalsFromRadius(lat, lng, radius float64) []Int64arr {
	radiusBits := geohash.FindBitDepth(radius)
	hash := geohash.EncodeInt(lat, lng, radiusBits)

	neighbors := buildBoxSet(hash, radiusBits)
	ranges := rangesFromBoxSet(neighbors)
	increaseRangeBitDepth(ranges, radiusBits)
	return ranges
}

// Get the neighbors of the bounding box for geohash and sort them
func buildBoxSet(hash, bitDepth int64) Int64arr {
	neighbors := Int64arr(geohash.NeighborsInt(hash, bitDepth))
	sort.Sort(neighbors)
	return neighbors
}

// If we have consecutive neighbors, ignore those and get bigger ranges. This
// will result in fewer Redis queries.
func rangesFromBoxSet(neighbors Int64arr) []Int64arr {
	var ranges []Int64arr
	for i := 0; i < len(neighbors); i++ {
		lowerBound := neighbors[i]
		upperBound := lowerBound + 1
		for neighbors.Include(upperBound) {
			upperBound++
			i++
		}
		ranges = append(ranges, Int64arr{lowerBound, upperBound})
	}

	return ranges
}

// Convert the neighbors from radius' bit depth to maximum bit depth
func increaseRangeBitDepth(ranges []Int64arr, radiusBitDepth int64) {
	bitDiff := uint(maxBitDepth - radiusBitDepth)
	for _, val := range ranges {
		val[0] = leftShift(val[0], bitDiff)
		val[1] = leftShift(val[1], bitDiff)
	}
}

// value * 2^bits
func leftShift(value int64, bits uint) int64 {
	return value * (1 << bits)
}
