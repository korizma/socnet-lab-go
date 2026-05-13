package graph

import "errors"

var ErrSliceFind = errors.New("Not found in slice")

func CompareNodes(node1 Node, node2 Node) bool {
	return node1.id == node2.id
}

func CompareEdges(edge1 Edge, edge2 Edge) bool {
	return edge1.id == edge2.id
}

func RemoveFromSlice[T any](slice []T, element T, equal func(T, T) bool) []T {
	for i, _ := range slice {

		if equal(element, slice[i]) {

			if i == len(slice)-1 {
				return slice[:len(slice)-1]
			}

			slice[i] = slice[len(slice)-1]
			return slice[:len(slice)-1]
		}
	}
	return slice
}

func FindInSlice[T any](slice []T, element T, equal func(T, T) bool) (int, error) {
	for i, val := range slice {

		if equal(val, element) {
			return i, nil
		}
	}
	return -1, ErrSliceFind
}

func CopySlice[T any](slice []T) []T {
	copySlice := make([]T, len(slice))
	copy(copySlice, slice)
	return copySlice
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	copyMap := make(map[K]V, len(m))
	for key, value := range m {
		copyMap[key] = value
	}
	return copyMap
}
