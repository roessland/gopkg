package mathutil

import "fmt"

var partitionsCache = map[int][][]int{
	1: {{1}},
}

// PartitionsInt finds the integer partitions
// of a positive integer n.
// E.g. 4 can be partitioned into:

//     4 =
//	    4,
//      3+1,
//      2+2,
//      2+1+1,
//      1+1+1+1.
func PartitionsInt(n int) [][]int {
	var recurse func(int) [][]int
	recurse = func(n int) [][]int {
		cached, ok := partitionsCache[n]
		if ok {
			return cached
		}
		res := [][]int{{n}}

		for i := 1; i < n; i++ {
			a := n - i
			rs := recurse(i)
			for _, r := range rs {
				if r[0] <= a {
					res = append(res, append([]int{a}, r...))
				}
			}
		}
		partitionsCache[n] = res
		return res
	}
	recurse(n)
	return partitionsCache[n]
}

// PartitionIntToMultisetIntInt converts a partition result into a multiset
func PartitionIntToMultisetIntInt(partition []int) MultisetIntInt {
	ms := MultisetIntInt{}
	for _, part := range partition {
		ms[part]++
	}
	return ms
}

// // PartitionsIntToMultisetsIntInt converts a list of
// partition results into a list of multisets
func PartitionsIntToMultisetsIntInt(partitions [][]int) []MultisetIntInt {
	multisets := []MultisetIntInt{}
	for _, partition := range partitions {
		multisets = append(multisets, PartitionIntToMultisetIntInt(partition))
	}
	return multisets
}

type MultisetIntInt map[int]int

func (ms MultisetIntInt) Equal(other MultisetIntInt) bool {
	for key, multiplicity := range ms {
		if other[key] != multiplicity {
			return false
		}
	}
	for key, multiplicity := range other {
		if ms[key] != multiplicity {
			return false
		}
	}
	return true
}

func (ms MultisetIntInt) String() string {
	return fmt.Sprintf("%v", map[int]int(ms))
}
