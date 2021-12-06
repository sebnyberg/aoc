package ax

import "constraints"

type DSU[T constraints.Integer] struct {
	parent []T
}

func NewDSU[T constraints.Integer](n T) DSU[T] {
	dsu := DSU[T]{
		parent: make([]T, n+1),
	}
	for i := range dsu.parent {
		dsu.parent[i] = T(i)
	}
	return dsu
}

func (d *DSU[T]) Find(a T) T {
	if d.parent[a] != a {
		root := d.Find(d.parent[a])
		d.parent[a] = root // Path compression
	}
	return d.parent[a]
}

func (d *DSU[T]) Union(a, b T) {
	rootA, rootB := d.Find(a), d.Find(b)
	if rootA != rootB {
		d.parent[rootB] = rootA
	}
}

type DSUMap[T comparable] struct {
	parent map[T]T
}

func NewDSUMap[T comparable](vals []T) DSUMap[T] {
	n := len(vals)
	dsu := DSUMap[T]{
		parent: make(map[T]T, n+1),
	}
	for _, val := range vals {
		dsu.parent[val] = val
	}
	return dsu
}

func (d *DSUMap[T]) Find(a T) T {
	if d.parent[a] != a {
		root := d.Find(d.parent[a])
		d.parent[a] = root // Path compression
	}
	return d.parent[a]
}

func (d *DSUMap[T]) Union(a, b T) {
	rootA, rootB := d.Find(a), d.Find(b)
	if rootA != rootB {
		d.parent[rootB] = rootA
	}
}
