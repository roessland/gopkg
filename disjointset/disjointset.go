package disjointset

// For usage, see the tests

type DisjointSet struct {
	rank   []int
	parent []int
	Count  int // Number of connected components
}

func Make(size int) *DisjointSet {
	var ds DisjointSet
	ds.Count = size
	ds.rank = make([]int, size)
	ds.parent = make([]int, size)
	for i := 0; i < size; i++ {
		ds.parent[i] = i
	}
	return &ds
}

func (ds *DisjointSet) Find(i int) int {
	if ds.parent[i] != i {
		ds.parent[i] = ds.Find(ds.parent[i])
	}
	return ds.parent[i]
}

func (ds *DisjointSet) Connected(x, y int) bool {
	return ds.Find(x) == ds.Find(y)
}

func (ds *DisjointSet) Union(x, y int) {
	xRoot := ds.Find(x)
	yRoot := ds.Find(y)
	if xRoot != yRoot {
		if ds.rank[xRoot] < ds.rank[yRoot] {
			ds.parent[xRoot] = yRoot
		} else if ds.rank[xRoot] > ds.rank[yRoot] {
			ds.parent[yRoot] = xRoot
		} else {
			ds.parent[yRoot] = xRoot
			ds.rank[xRoot]++
		}
		ds.Count--
	}
}
