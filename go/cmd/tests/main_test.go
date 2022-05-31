package main

import (
	"testing"
)

func FuzzPrintAny(f *testing.F) {
	strTestCases := []string{"asdasd", "asdw", "", ""}
	for i := range strTestCases {
		f.Add(strTestCases[i])
	}

	f.Fuzz(func(t *testing.T, in string) {
		PrintAny(in)
	})
}
