package warmups

import "fmt"

/*
 * Minimum number in array, input: [2, 1, 3, 6, 4, 8]
 */
func MinNumber(arr []int) int {
	min := -1
	if arr != nil {
		min = arr[0]
		for _, v := range arr {
			if min > v {
				min = v
			}
		}
	}
	return min
}

/*
 * Maximum number in array, input: [2, 1, 3, 6, 4, 8]
 */
func MaxNumber(arr []int) int {
	max := -1
	if arr != nil {
		max = arr[0]
		for _, v := range arr {
			if max < v {
				max = v
			}
		}
	}
	return max
}

/*
 * Minimum and Maximum number in array(Single Iteration), input: [2, 1, 3, 6, 4,
 * 8]
 */
func MinMaxNumber(arr []int) string {
	min := -1
	max := -1
	if arr != nil {
		min = arr[0]
		max = arr[1]
		for _, v := range arr {
			if max < v {
				max = v
			} else if min > v {
				min = v
			}
		}
	}

	return fmt.Sprintf("%s %d %s %d %s", "[Min: ", min, ",Max: ", max, "]")
}

/*
 * Sum of two number equal to k, input: [2, 1, 3, 6, 4, 8]
 */
func SumOfTwoNumbers(arr []int, k int) []string {

	var result []string

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == k {
				pair := fmt.Sprintf("(%d, %d)", arr[i], arr[j])
				result = append(result, pair)
			}
		}
	}

	return result
}
