package TwoSum

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	nums     []int
	target   int
	expected []int
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{nums: []int{6, 2, 3}, target: 9, expected: []int{0, 2}, comment: "NormalCase1"},
		{nums: []int{6, 2, 3}, target: 8, expected: []int{0, 1}, comment: "NormalCase2"},
		{nums: []int{6, 3, 2}, target: 8, expected: []int{0, 2}, comment: "NormalCase3"},
		{nums: []int{5, 3, 2, 1, 4}, target: 4, expected: []int{1, 3}, comment: "BiggerInput"},
		{nums: []int{3, 2, 1, 4}, target: 6, expected: []int{1, 3}, comment: "NoSumWithItself"},
		{nums: []int{3, 0, 2, 1, 3}, target: 3, expected: []int{0, 1}, comment: "ZeroCalculation"},
		{nums: []int{2, 3, 4}, target: 8, expected: nil, comment: "NoSumEqualTarget"},
		{nums: []int{2}, target: 8, expected: nil, comment: "NumsInsufficient"},
	}

	for index, element := range testSets {
		t.Run("Test BruteForce: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testAppBruteForce(element.nums, element.target, element.expected, index, t) })
		t.Run("Test Map: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testAppMap(element.nums, element.target, element.expected, index, t) })
	}
}

func testAppBruteForce(nums []int, target int, expected []int, index int, t *testing.T) {
	actual, _ := twoSumBruteForce(nums, target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: [" + strconv.Itoa(expected[0]) + "," + strconv.Itoa(expected[1]) +
			"] Actual: [" + strconv.Itoa(actual[0]) + "," + strconv.Itoa(actual[1]) + "]")
	}
}

func testAppMap(nums []int, target int, expected []int, index int, t *testing.T) {
	actual, _ := twoSumMap(nums, target)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: [" + strconv.Itoa(expected[0]) + "," + strconv.Itoa(expected[1]) +
			"] Actual: [" + strconv.Itoa(actual[0]) + "," + strconv.Itoa(actual[1]) + "]")
	}
}

func BenchmarkApp1(b *testing.B) {
	benchmarkSets := []InputCase{
		{nums: []int{1, 2, 3}, target: 3, comment: "BestCase 3"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 3, comment: "BestCase 10"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, target: 3, comment: "BestCase 20"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 31, 33, 34, 35, 36, 37, 38, 39, 40}, target: 3, comment: "BestCase 40"},
		{nums: []int{1, 2, 3}, target: 5, comment: "WorstCase 3"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 19, comment: "WorstCase 10"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, target: 39, comment: "WorstCase 20"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 31, 33, 34, 35, 36, 37, 38, 39, 40}, target: 39, comment: "WorstCase 40"},
		{nums: []int{1, 2, 3}, target: 99, comment: "NonExisting 3"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 99, comment: "NonExisting 10"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, target: 99, comment: "NonExisting 20"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 31, 33, 34, 35, 36, 37, 38, 39, 40}, target: 99, comment: "NonExisting 40"},
		{nums: []int{1, 2, 3, 4, 5}, target: 7, comment: "RandomCase 5"},
		{nums: []int{1, 2, 3, 4, 5}, target: 8, comment: "RandomCase 5"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 10, comment: "RandomCase 10"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 13, comment: "RandomCase 10"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 11, comment: "RandomCase 10"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, target: 22, comment: "RandomCase 20"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, target: 25, comment: "RandomCase 20"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, target: 16, comment: "RandomCase 20"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, target: 32, comment: "RandomCase 20"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 31, 33, 34, 35, 36, 37, 38, 39, 40}, target: 44, comment: "RandomCase 40"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 31, 33, 34, 35, 36, 37, 38, 39, 40}, target: 67, comment: "RandomCase 40"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 31, 33, 34, 35, 36, 37, 38, 39, 40}, target: 18, comment: "RandomCase 40"},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 31, 33, 34, 35, 36, 37, 38, 39, 40}, target: 57, comment: "RandomCase 40"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark BruteForce: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppBruteForce(element.nums, element.target, b) })
		b.Run("Benchmark Map: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppMap(element.nums, element.target, b) })
	}
}

func benchmarkAppBruteForce(nums []int, target int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = twoSumBruteForce(nums, target)
	}
}

func benchmarkAppMap(nums []int, target int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = twoSumMap(nums, target)
	}
}
