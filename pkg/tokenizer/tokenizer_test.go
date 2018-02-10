package tokenizer

import (
	"reflect"
	"testing"
)

func TestSplitToSlices(t *testing.T) {
	tt := []struct {
		name     string
		str      string
		words    uint
		expected [][]string
	}{
		{
			name:     "An Empty String",
			str:      "",
			words:    3,
			expected: [][]string{},
		},
		{
			name:  "1 Split",
			str:   "alpha bravo charlie delta echo foxtrot golf hotel india",
			words: 1,
			expected: [][]string{
				{
					"alpha", "bravo", "charlie", "delta", "echo", "foxtrot", "golf", "hotel", "india",
				},
			},
		},
		{
			name:  "With extra whitespaces",
			str:   "  alpha      charlie     echo     golf    india",
			words: 1,
			expected: [][]string{
				{
					"alpha", "charlie", "echo", "golf", "india",
				},
			},
		},
		{
			name:  "Evenly-Sized Subsets",
			str:   "alpha bravo charlie delta echo foxtrot golf hotel india",
			words: 3,
			expected: [][]string{
				{"alpha", "bravo", "charlie"},
				{"delta", "echo", "foxtrot"},
				{"golf", "hotel", "india"},
			},
		},
		{
			name:  "Unevenly-Sized Subsets",
			str:   "alpha bravo charlie delta echo foxtrot golf hotel india",
			words: 2,
			expected: [][]string{
				{"alpha", "bravo"},
				{"charlie", "delta"},
				{"echo", "foxtrot"},
				{"golf", "hotel"},
				{"india"},
			},
		},
		{
			name:  "Split Is Greater Than String Word Count",
			str:   "alpha bravo charlie delta echo foxtrot golf hotel india",
			words: 12,
			expected: [][]string{
				{
					"alpha", "bravo", "charlie", "delta", "echo", "foxtrot", "golf", "hotel", "india",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := SplitToSlices(tc.str, tc.words)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("\nSplitToSlices of %s\nshould be:\n\t%v\nbut got:\n\t%v", tc.name, tc.expected, actual)
			}
		})
	}
}

var result [][]string

func benchmarkSplitToSlices(str string, n uint, b *testing.B) {
	var r [][]string

	for i := 0; i < b.N; i++ {
		r = SplitToSlices(str, n)
	}

	result = r
}

func BenchmarkSplitToSlices1(b *testing.B) {
	benchmarkSplitToSlices("alpha bravo charlie delta echo foxtrot golf hotel india", 3, b)
}

func BenchmarkSplitToSlices2(b *testing.B) {
	benchmarkSplitToSlices("alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india alpha bravo charlie delta echo foxtrot golf hotel india", 20, b)
}
