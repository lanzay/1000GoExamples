package main

import "testing"

func TestFooN(t *testing.T) {
	var tests1 = []struct {
		inputN int
		want   float64
	}{
		{4, 33},
	}
	for _, test := range tests1 {
		if got := foo1(test.inputN); got != test.want {
			t.Errorf("foo1(%v) = %v", test.inputN, got)
		}
	}
	var tests2 = []struct {
		inputN int
		want   float64
	}{
		{4, 24},
	}
	for _, test := range tests2 {
		if got := foo2(test.inputN); got != test.want {
			t.Errorf("foo2(%v) = %v", test.inputN, got)
		}
	}
}
