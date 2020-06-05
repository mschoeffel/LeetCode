package TwoSum

import (
	"errors"
	"fmt"
	"strconv"
)

func Main() {
	nums := []int{3, 4, 4}
	target := 6

	result, err := twoSumBruteForce(nums, target)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result Brute Force: Number 1: " + strconv.Itoa(result[0]) + " Number 2: " + strconv.Itoa(result[1]))
	}

	result, err = twoSumMap(nums, target)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result Map: Number 1: " + strconv.Itoa(result[0]) + " Number 2: " + strconv.Itoa(result[1]))
	}
}

func twoSumBruteForce(nums []int, target int) ([]int, error) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}, nil
			}
		}
	}
	return nil, errors.New("error: There is no pair that sums up to the given target")
}

func twoSumMap(nums []int, target int) ([]int, error) {
	indices := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := indices[target-num]; ok {
			return []int{j, i}, nil
		}
		indices[num] = i
	}
	return nil, errors.New("error: There is no pair that sums up to the given target")
}
