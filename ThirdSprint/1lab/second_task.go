package main

func Filter[T any](nums []T, f func(T) bool) []T {
	var filtered []T
	for _, n := range nums {
		if f(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}
