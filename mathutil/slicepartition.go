package mathutil

// SlicePartitions find all partitions of s, where each part has positive length.
// A slice of length n has 2^(n-1) partitions.
// For example, the partitions of [1 2 3 4] are:
//
//	((1,2,3,4))
//  ((1,2,3), (4))
//  ((1,2), (3,4))
//  ((1,2), (3), (4))
//  ((1), (2,3,4))
//  ((1), (2,3), (4))
//  ((1), (2), (3,4))
//  ((1), (2), (3), (4))
//
// So these are the unique ways to split the list into parts. Concatenating any partition of a sequence will result
// in the original list again.
func SlicePartitions(slice []int) [][][]int {
	if len(slice) == 0 {
		return [][][]int{}
	}
	partitions := make([][][]int, int(Pow(2, int64(len(slice)-1))))
	compositions := Compositions(len(slice))
	for i, comp := range compositions {
		partitions[i] = getSlicePartitionFromComposition(slice, comp)
	}
	return partitions
}

func getSlicePartitionFromComposition(sequence []int, comp []int) [][]int {
	var partition [][]int
	i := 0
	for _, partitionSize := range comp {
		partition = append(partition, sequence[i:i+partitionSize])
		i += partitionSize
	}
	return partition
}
