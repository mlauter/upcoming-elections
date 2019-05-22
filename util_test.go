package main

import (
	"testing"
)

type stringInSliceInput struct {
	s   string
	arr []string
}

func TestStringInSlice(t *testing.T) {
	tests := []struct {
		Input  *stringInSliceInput
		Expect bool
	}{
		{
			&stringInSliceInput{s: "a", arr: []string{"a", "b", "c"}},
			true,
		},
		{
			&stringInSliceInput{s: "CA", arr: []string{"CA", "NY", "MA"}},
			true,
		},
		{
			&stringInSliceInput{s: "CA", arr: []string{}},
			false,
		},
		{
			&stringInSliceInput{s: "CA", arr: []string{"a", "b"}},
			false,
		},
		{
			&stringInSliceInput{s: "", arr: []string{"a", "b"}},
			false,
		},
	}

	for _, tc := range tests {
		actual := stringInSlice(tc.Input.s, tc.Input.arr)
		if actual != tc.Expect {
			t.Errorf("expected %t, got %t", tc.Expect, actual)
		}
	}
}
