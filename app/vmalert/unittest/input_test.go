package unittest

import (
	"reflect"
	"testing"
)

func TestParseInputValue(t *testing.T) {
	testCases := []struct {
		input  string
		exp    []sequenceValue
		failed bool
	}{
		{
			"",
			nil,
			true,
		},
		{
			"testfailed",
			nil,
			true,
		},
		{
			"-4",
			[]sequenceValue{{Value: -4}},
			false,
		},
		{
			"_",
			[]sequenceValue{{Omitted: true}},
			false,
		},
		{
			"-4x1",
			[]sequenceValue{{Value: -4}, {Value: -4}},
			false,
		},
		{
			"_x1",
			[]sequenceValue{{Omitted: true}},
			false,
		},
		{
			"1+1x4",
			[]sequenceValue{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
			false,
		},
		{
			"2-1x4",
			[]sequenceValue{{Value: 2}, {Value: 1}, {Value: 0}, {Value: -1}, {Value: -2}},
			false,
		},
		{
			"1+1x1 _ -4 3+20x1",
			[]sequenceValue{{Value: 1}, {Value: 2}, {Omitted: true}, {Value: -4}, {Value: 3}, {Value: 23}},
			false,
		},
	}

	for _, tc := range testCases {
		output, err := parseInputValue(tc.input, true)
		if err != nil != tc.failed {
			t.Fatalf("failed to parse %s, expect %t, got %t", tc.input, tc.failed, err != nil)
		}
		if !reflect.DeepEqual(tc.exp, output) {
			t.Fatalf("expect %v, got %v", tc.exp, output)
		}
	}
}
