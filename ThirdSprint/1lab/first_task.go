package main

type MyConstraint interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Sum[T MyConstraint](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}
