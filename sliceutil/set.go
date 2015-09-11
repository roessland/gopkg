package sliceutil

// Intersection of two setslices
func Intersect(a, b []int64) []int64 {
	values := make(map[int64]int)
	for _, val := range a {
		values[val]++
	}
	for _, val := range b {
		values[val]++
	}

	intersection := []int64{}
	for value, count := range values {
		if count == 2 {
			intersection = append(intersection, value)
		}
	}
	SortInt64(intersection)
	return intersection
}

// Intersection of two setslices
func IntersectInt(a, b []int) []int {
	values := make(map[int]int)
	for _, val := range a {
		values[val]++
	}
	for _, val := range b {
		values[val]++
	}

	intersection := []int{}
	for value, count := range values {
		if count == 2 {
			intersection = append(intersection, value)
		}
	}
	SortInt(intersection)
	return intersection
}

// Union of two setslices
func Union(a, b []int64) []int64 {
	values := make(map[int64]int)
	for _, val := range a {
		values[val]++
	}
	for _, val := range b {
		values[val]++
	}

	union := []int64{}
	for value, count := range values {
		if count >= 1 {
			union = append(union, value)
		}
	}
	SortInt64(union)
	return union
}
