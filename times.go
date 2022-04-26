package main

import (
	"fmt"
	"strconv"
)

/*
Task: Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34) that the given integers can be assembled into and return the total number of valid times.
You can not use the same number twice.
Times such as 34:12 and 12:60 are not valid.
Provided integers can not be negative.
Notes: Input integers can not be negative.
Input integers can yeald 0 possible valid combinations.
Example:
	Input: [1, 2, 3, 4]
	Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41", "21:34", "21:43"]
	Return: 10
*/

func possibleTimes(digits []int) int {
	// Your code here

	// --- input length must be 4 integers
	if len(digits) != 4 {
		fmt.Println("input length must be 4")
		return 0
	}

	// --- input can't be negative
	for i, v := range digits {
		if v < 0 {
			msg := fmt.Sprintf("input index : [%v] can't be negative (given value : %v)", i, v)
			fmt.Println(msg)
			return 0
		}
	}

	// --- prepare var to catch all combinations
	var combinations []string

	// --- loop 1, define 1st digit
	for i1, v1 := range digits {
		// --- not valid if more than 2
		if v1 > 2 {
			continue
		}

		// --- first digit is valid
		temp1 := fmt.Sprintf("%v", v1)

		// --- loop 2, define 2nd digit
		for i2, v2 := range digits {
			// --- continue if it's the same index with loop 1
			if i2 == i1 {
				continue
			}

			// --- found (probably) valid 2nd digit
			temp2 := fmt.Sprintf("%s%v", temp1, v2)

			// --- check combination between 1st and 2nd digit not more than 23 (ignore error because it should be valid integer)
			x, _ := strconv.Atoi(temp2)
			if x > 23 {
				continue
			}

			// --- loop 3, define 3rd digit
			for i3, v3 := range digits {
				// --- continue if it's the same index with loop 1 or loop 2
				if i3 == i1 || i3 == i2 {
					continue
				}

				// --- continue if it's more than 5 as it is not valid third digit of 24hour format
				if v3 > 5 {
					continue
				}

				// --- found valid 3rd digit
				temp3 := fmt.Sprintf("%s:%v", temp2, v3)

				// --- loop 4, define 4th digit
				for i4, v4 := range digits {
					// --- continue if it's the same index with loop 1 or loop 2 or loop 3
					if i4 == i1 || i4 == i2 || i4 == i3 {
						continue
					}

					// --- continue if it's more than 9 as it is not valid fourth digit of 24hour format
					if v4 > 9 {
						continue
					}

					// --- found valid 4th digit
					temp4 := fmt.Sprintf("%s%v", temp3, v4)

					// --- check if it's duplicate data or not
					if !array_contains(combinations, temp4) {
						combinations = append(combinations, temp4)
					}

				}
			}
		}

	}

	// --- checkers
	// for _, v := range combinations {
	// 	fmt.Println(v)
	// }

	return len(combinations)
}

func array_contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	// Example test cases.
	fmt.Println(possibleTimes([]int{1, 2, 3, 4}))
	// fmt.Println(possibleTimes([]int{9, 1, 2, 0}))
	// fmt.Println(possibleTimes([]int{2, 2, 1, 9}))
	// fmt.Println(possibleTimes([]int{0, -1, 2, 2}))
	// fmt.Println(possibleTimes([]int{0, -1, 2}))
}
