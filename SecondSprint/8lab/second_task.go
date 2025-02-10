package main

import "testing"

func TestContains(t *testing.T) {
	cases := []struct {
		name     string
		arr      []int
		elem     int
		expected bool
	}{{name: "casual", arr: []int{5, 2, 3}, elem: 2, expected: true},
		{name: "test1", arr: []int{6, 5, 4}, elem: 3, expected: false},
		{name: "test2", arr: []int{10, 30, 0, 20}, elem: 10, expected: true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := Contains(tc.arr, tc.elem)
			if got != tc.expected {
				t.Errorf("answer== %v; want== %v", got, tc.expected)
			}
		})
	}
}

func Contains(numbers []int, target int) bool {
	for _, num := range numbers {
		if num == target {
			return true
		}
	}
	return false
}
