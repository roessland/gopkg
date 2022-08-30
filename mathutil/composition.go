package mathutil

var compositionsCache map[int][][]int = map[int][][]int{
	0: nil,
}

// Compositions finds all A-restricted compositions of a positive integer.
// https://en.wikipedia.org/wiki/Composition_(combinatorics)
//
// The sixteen compositions (there are 2^(n-1) compositions of n) of 5 are:
//   5
//   4 + 1
//   3 + 2
//   3 + 1 + 1
//   2 + 3
//   2 + 2 + 1
//   2 + 1 + 2
//   2 + 1 + 1 + 1
//   1 + 4
//   1 + 3 + 1
//   1 + 2 + 2
//   1 + 2 + 1 + 1
//   1 + 1 + 3
//   1 + 1 + 2 + 1
//   1 + 1 + 1 + 2
//   1 + 1 + 1 + 1 + 1.
//
// Compare this with the seven _partitions_ of 5:
//   5
//   4 + 1
//   3 + 2
//   3 + 1 + 1
//   2 + 2 + 1
//   2 + 1 + 1 + 1
//   1 + 1 + 1 + 1 + 1.
//
func Compositions(n int) [][]int {
	cached, ok := compositionsCache[n]
	if ok {
		return cached
	}
	compositionsCache[n] = compositions(nil, []int{}, n)
	return compositionsCache[n]
}

// [] (k=2)
//   2 (i=2, k=0)
//   1 (i=1, k=1)
//     1 1 (i=1, k=0)
// [] (k=5)
// 5 (k=0)
// 4 (k=1)
//	4 1 (k=0)
// 3 (k=2)
//	3 2 (k=0)
//  3 1 (k=1)
//    3 1 1 (k=0)
// 2 (k=3)
//  2 3 (k=0)
//	2 2 (k=1)
//	  2 2 1 (k=0)
//	2 1 (k=2)
//    2 1 2 (k=0)
//    2 1 1 (k=1)
//	    2 1 1 1 (k=0)
// 1 (k=4)
//  1 4 (k=0)
//  1 3 (k=1)
//    1 3 1 (k=0)
//  1 2 (k=2)
//    1 2 2 (k=0)
//    1 2 1 (k=1)
//      1 2 1 1 (k=0)
//  1 1 (k=3)
//    1 1 3 (k=0)
//    1 1 2 (k=1)
//      1 1 2 1 (k=0)
//    1 1 1 (k=2)
//      1 1 1 2 (k=0)
//      1 1 1 1 (k=1)
//        1 1 1 1 1 (k=0)
// basecase: k=0 -> add to output
// otherwise: loop i from k to 0
func compositions(comps [][]int, prev []int, k int) [][]int {
	if k == 0 {
		comps = append(comps, prev)
	}
	prevCopy := make([]int, len(prev))
	copy(prevCopy, prev)
	for i := k; i > 0; i-- {
		comps = compositions(comps, append(prevCopy, i), k-i)
	}
	return comps
}
