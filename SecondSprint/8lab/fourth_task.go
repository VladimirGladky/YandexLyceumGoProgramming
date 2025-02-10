package main

import (
	"sort"
	"strings"
	"testing"
)

func TestAreAnagrams(t *testing.T) {
	cases := []struct {
		name string
		str  string
		str1 string
		want bool
	}{{name: "casual", str: "abc", str1: "cba", want: true},
		{name: "test1", str: "abcdefgh", str1: "hgfedcba", want: true},
		{name: "test2", str: "478235111111", str1: "38532874", want: false},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := AreAnagrams(tc.str, tc.str1)
			if got != tc.want {
				t.Errorf("answer== %v; want== %v", got, tc.want)
			}
		})
	}
}

func AreAnagrams(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if len(str1) != len(str2) {
		return false
	}

	// Convert strings to slices of runes for sorting
	r1 := []rune(str1)
	r2 := []rune(str2)

	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})

	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	return string(r1) == string(r2)
}
