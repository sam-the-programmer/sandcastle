package parse_test

import (
	"castle/parse"
	"testing"
)

func TestRejoinStringArgs(t *testing.T) {
	testCases := []struct {
		desc  string
		input []string
		want  []string
	}{
		{
			"NoChange",
			[]string{"These", "are", "separate."},
			[]string{"These", "are", "separate."},
		},
		{
			"SingleString",
			[]string{"This", "is", "\"one", "string\""},
			[]string{"This", "is", "\"one string\""},
		},
		{
			"MultipleStrings",
			[]string{"This", "is", "\"one", "string\"", "and", "\"another", "string\""},
			[]string{"This", "is", "\"one string\"", "and", "\"another string\""},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			reconstructed := parse.RejoinStringArgs(tC.input)
			if len(reconstructed) != len(tC.want) {
				t.Errorf("Reconstructed length (%d) does not match want length (%d).", len(reconstructed), len(tC.want))
			}

			for i, arg := range reconstructed {
				if arg != tC.want[i] {
					t.Errorf("Reconstructed arg (%s) does not match want arg (%s).", arg, tC.want[i])
				}
			}
		})
	}
}

func TestHasPrefixSlice(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		prefixes [3]string
		want     bool
	}{
		{
			desc:     "NoPrefix",
			input:    "This is a string.",
			prefixes: parse.STRING_STARTERS,
			want:     false,
		},
		{
			desc:     "SinglePrefix",
			input:    "`This is a string.",
			prefixes: parse.STRING_STARTERS,
			want:     true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if parse.HasPrefixSlice(tC.input, tC.prefixes) != tC.want {
				t.Errorf("HasPrefixSlice returned %t, want %t.", parse.HasPrefixSlice(tC.input, tC.prefixes), tC.want)
			}
		})
	}
}
