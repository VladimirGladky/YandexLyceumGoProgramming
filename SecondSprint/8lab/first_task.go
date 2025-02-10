package main

import (
	"slices"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	cases := []struct {
		name  string
		start []int
		want  []int
	}{{name: "casual", start: []int{5, 2, 3}, want: []int{2, 3, 5}},
		{name: "test1", start: []int{6, 5, 4}, want: []int{4, 5, 6}},
		{name: "test2", start: []int{10, 30, 0, 20}, want: []int{0, 10, 20, 30}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			SortIntegers(tc.start)
			for i := 0; i < len(tc.start); i++ {
				if tc.start[i] != tc.want[i] {
					t.Errorf("answer== %v; want== %v", tc.start, tc.want)
				}
			}
		})
	}
}

func SortIntegers(numbers []int) {
	slices.Sort(numbers)
}
