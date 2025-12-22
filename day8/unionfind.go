package main

import "fmt"

type UnionFind[T comparable] struct {
	Parent map[*T]*T
	Rank   map[*T]int
	Size   map[*T]int
	Count  int
}

func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		Parent: make(map[*T]*T),
		Rank:   make(map[*T]int),
		Size:   make(map[*T]int),
		Count:  0,
	}
}

func (uf *UnionFind[T]) Find(item *T) *T {
	if _, exists := uf.Parent[item]; !exists {
		uf.Parent[item] = item
		uf.Rank[item] = 0
		uf.Size[item] = 1
		uf.Count++
	}

	if uf.Parent[item] != item {
		uf.Parent[item] = uf.Find(uf.Parent[item])
	}

	return uf.Parent[item]
}

func (uf *UnionFind[T]) Union(item1, item2 *T) bool {
	root1 := uf.Find(item1)
	root2 := uf.Find(item2)

	if root1 != root2 {
		if uf.Rank[root1] > uf.Rank[root2] {
			uf.Parent[root2] = root1
			uf.Size[root1] += uf.Size[root2]
		} else {
			uf.Parent[root1] = root2
			uf.Size[root2] += uf.Size[root1]
			if uf.Rank[root1] == uf.Rank[root2] {
				uf.Rank[root2]++
			}
		}
		uf.Count--
		return true
	} else {
		return false
	}
}

func (uf *UnionFind[T]) ToString() string {
	result := ""
	for item, parent := range uf.Parent {
		if item == parent && uf.Size[item] > 1 {
			result += fmt.Sprintf("Item: %v, Size: %d\n",
				*item, uf.Size[item])
		}

	}
	return result
}
