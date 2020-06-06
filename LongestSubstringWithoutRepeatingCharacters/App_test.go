package LongestSubstringWithoutRepeatingCharacters

import (
	"strconv"
	"testing"
)

type InputCase struct {
	s        string
	expected int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s: "bbbbb", expected: 1, comment: "OnlyOneRuneCase"},
		{s: "abcabcbb", expected: 3, comment: "TestCaseOne"},
		{s: "pwwkew", expected: 3, comment: "TestCaseTwo"},
		{s: "", expected: 0, comment: "EmptyCase"},
		{s: "pabcdefpghij", expected: 11, comment: "OneDuplicateCase"},
		{s: "pabcdefpaghij", expected: 11, comment: "MultipleDuplicateCase"},
		{s: "pabbcdefpghij", expected: 10, comment: "DoubleCase"},
		{s: "pabcdefpghii", expected: 10, comment: "DoubleEndCase"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testApp(element.s, element.expected, index, t) })
	}
}

func testApp(s string, expected int, index int, t *testing.T) {
	actual := lengthOfLongestSubstring(s)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: [" + strconv.Itoa(expected) +
			"] Actual: [" + strconv.Itoa(actual) + "]")
	}
}

func BenchmarkApp1(b *testing.B) {
	benchmarkSets := []InputCase{
		{s: "bbbbb", expected: 1, comment: "OnlyOneRuneCase"},
		{s: "abcabcbb", expected: 3, comment: "TestCaseOne"},
		{s: "pwwkew", expected: 3, comment: "TestCaseTwo"},
		{s: "", expected: 0, comment: "EmptyCase"},
		{s: "pabcdefpghij", expected: 11, comment: "OneDuplicateCase"},
		{s: "pabcdefpaghij", expected: 11, comment: "MultipleDuplicateCase"},
		{s: "pabbcdefpghij", expected: 10, comment: "DoubleCase"},
		{s: "pabcdefpghii", expected: 10, comment: "DoubleEndCase"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark BruteForce: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.s, b) })
	}
}

func benchmarkApp(s string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = lengthOfLongestSubstring(s)
	}
}
