package main

import "testing"

func TestReverseString(t *testing.T) {
	cases := []struct {
		name     string
		str      string
		expected string
	}{{name: "casual", str: "abc", expected: "cba"},
		{name: "test1", str: "abcdefgh", expected: "hgfedcba"},
		{name: "test2", str: "47823583", expected: "38532874"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := ReverseString(tc.str)
			if got != tc.expected {
				t.Errorf("answer== %v; want== %v", got, tc.expected)
			}
		})
	}
}

func ReverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
