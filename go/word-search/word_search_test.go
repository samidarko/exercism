package wordsearch

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := Solve(tc.words, tc.puzzle)
			switch {
			case tc.expectError:
				if err == nil {
					t.Fatalf("Solve(%v,%v) expected error, got:%v", tc.words, tc.puzzle, actual)
				}
			case err != nil:
				t.Fatalf("Solve(%v,%v) returned error: %v, want:%v", tc.words, tc.puzzle, err, tc.expected)
			case !reflect.DeepEqual(actual, tc.expected):
				t.Fatalf("Solve(%v,%v) = %v, want:%v", tc.words, tc.puzzle, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Solve(tc.words, tc.puzzle)
		}
	}
}

func TestGetDiagonalsTopLeftBottomRight(t *testing.T) {
	puzzle := []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"}
	actual := getDiagonalsTopLeftBottomRight(puzzle)
	expected := []string{"c", "jl", "aao", "sllj", "wcxau", "rorhyr", "pilepce", "obxcebar", "ciwiqaulm", "javallurmt", "emoseimyp", "fdkqlrgi", "bcprhpr", "lirorc", "pmjus", "egsa", "ptm", "rc", "e"}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected: %v\n\tGot: %v", expected, actual)
	}

	actual = getDiagonalsTopLeftBottomRight([]string{"s", "u", "r", "a", "b", "c", "t"})
	expected = []string{"t", "c", "b", "a", "r", "u", "s", "", "", "", "", "", ""}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected: %v\n\tGot: %v", expected, actual)
	}

}

func TestGetDiagonalsTopRightBottomLeft(t *testing.T) {
	puzzle := []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"}
	actual := getDiagonalsTopRightBottomLeft(puzzle)
	expected := []string{"j", "ec", "oaf", "sllj", "wcxau", "rorhyr", "pilepce", "obxcebar", "ciwiqaulm", "javallurmt", "emoseimyp", "fdkqlrgi", "bcprhpr", "lirorc", "pmjus", "egsa", "ptm", "rc", "e"}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected: %v\n\tGot: %v", expected, actual)
	}
	//
	//actual = getDiagonalsTopRightBottomLeft([]string{"s", "u", "r", "a", "b", "c", "t"})
	//expected = []string{"t", "c", "b", "a", "r", "u", "s", "", "", "", "", "", ""}
	//if !reflect.DeepEqual(actual, expected) {
	//	t.Fatalf("Expected: %v\n\tGot: %v", expected, actual)
	//}

}
