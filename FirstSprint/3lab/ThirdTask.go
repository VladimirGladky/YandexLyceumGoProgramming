package main

import "strings"

type UpperWriter struct {
	UpperString string
}

func (u *UpperWriter) Write(str []byte) (int, error) {
	u.UpperString = strings.ToUpper(string(str))
	return len(str), nil
}
