package mathutil

func Partition(parts []int64, sum int64, numParts int64, depth int64) [][]int64 {

	if numParts == 0 || len(parts) == 0 {
		return [][]int64(nil)
	}

	if numParts == 1 {
		// The sum must be a part, otherwise return nil
		for i := 0; i < len(parts); i++ {
			if parts[i] == sum {
				return [][]int64{[]int64{sum}}
			}
		}
		return [][]int64(nil)
	}

	withLargestPart := Partition(parts[:len(parts)-1], sum-parts[len(parts)-1], numParts-1, depth+1)
	withoutLargestPart := Partition(parts[:len(parts)-1], sum, numParts, depth+1)

	// Add the missing largest part
	for i, _ := range withLargestPart {
		withLargestPart[i] = append(withLargestPart[i], parts[len(parts)-1])
	}

	both := append(withLargestPart, withoutLargestPart...)
	return both
}
