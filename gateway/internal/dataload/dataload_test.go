package dataload

import (
	"testing"
)

func TestSplitCoordinate(t *testing.T) {

	tests := []struct {
		name    string
		input   float64
		intPart int32
		frPart  int32
	}{
		{
			name:    "test1",
			input:   54.37,
			intPart: 54,
			frPart:  3700000,
		},
		{
			name:    "test2",
			input:   55.5136433,
			intPart: 55,
			frPart:  5136433,
		},
		{
			name:    "test2",
			input:   51.123456789,
			intPart: 51,
			frPart:  1234568,
		},
	}

	for _, tc := range tests {
		gotInt, gotFr := splitCoordinate(tc.input)
		if gotInt != tc.intPart || gotFr != tc.frPart {
			t.Fatalf("%s: expected: %v.%v, got: %v.%v", tc.name, tc.intPart, tc.frPart, gotInt, gotFr)
		}
	}

}
