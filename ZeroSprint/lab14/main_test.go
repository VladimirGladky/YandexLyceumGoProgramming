package main

import (
	"errors"
	"testing"
)

func TestSum(t *testing.T) {
	res := Sum(1, 1)
	if res != (2) {
		t.Error("Wrong answer")
	}
}

func TestMultiply(t *testing.T) {
	res := Multiply(2, 3)
	if res != 6 {
		t.Error("Wrong answer")
	}
}

func TestDeleteVowels(t *testing.T) {
	cases := []struct {
		name  string
		value string
		want  string
	}{
		{name: "oneVowels", value: "ddddaddd", want: "ddddddd"},
		{name: "allVowels", value: "aaaa", want: ""},
		{name: "noVowels", value: "dddd", want: "dddd"},
		{name: "noVowels", value: "oddddo", want: "dddd"},
		{name: "noVowels", value: "iddddi", want: "dddd"},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := DeleteVowels(tc.value)
			if got != tc.want {
				t.Errorf("Sum(%v) = %v; want %v", tc.value, got, tc.want)
			}
		})
	}
}

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
		err      error
	}{
		{input: []byte("Hello, 世界"), expected: 9, err: nil},
		{input: []byte{0x80}, expected: 0, err: ErrInvalidUTF8},
		{input: []byte{}, expected: 0, err: nil},
		{input: []byte("你"), expected: 1, err: nil},
	}

	for _, tt := range tests {
		got, err := GetUTFLength(tt.input)
		if got != tt.expected {
			t.Errorf("GetUTFLength() got = %v, expected = %v", got, tt.expected)
		}
		if !errors.Is(err, tt.err) {
			t.Errorf("GetUTFLength() error = %v, expected = %v", err, tt.err)
		}
	}
}
