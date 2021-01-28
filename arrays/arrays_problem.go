package arrays

/*
 * Arrange a greater to right and smaller in left of the given number k,input: [
 * 12, 2, 3, 1, 8, 123, 7, 11, 5, 4, 40, 100 ]
 *
 */
func Arrange(arr []int, k int) []int {

	first := 0
	last := len(arr) - 2

	for i := 0; i < len(arr); i++ {
		if arr[i] == k {
			temp := arr[len(arr)-1]
			arr[len(arr)-1] = k
			arr[i] = temp
		}

	}
	for first < last {
		if arr[first] > k {
			current := arr[first]
			arr[first] = arr[last]
			arr[last] = current
			last--
		} else {
			first++
		}

	}
	temp := arr[first]
	arr[first] = arr[len(arr)-1]
	arr[len(arr)-1] = temp

	return arr
}

/*
 * https://leetcode.com/problems/replace-elements-with-greatest-element-on-right
 * -side/,input: [ 17,18,5,4,6,1 ]
 *
 */
func ReplaceElementsAtRight(arr []int) []int {
	max := arr[len(arr)-1]

	arr1 := make([]int, len(arr))
	arr1[len(arr)-1] = -1
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr1[i] {
			max = arr[i]
			arr1[i-1] = max
		} else {
			max = arr1[i]
			arr1[i-1] = max
		}
	}
	return arr1
}

/*
 * replace-elements-with-greatest-element-on-left, input: [ 17,18,5,4,6,1 ]
 *
 */

func ReplaceElementsAtLeft(arr []int) []int {
	max := arr[0]
	arr1 := make([]int, len(arr))
	arr1[0] = -1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr1[i] {
			max = arr[i]
			arr1[i+1] = max
		} else {
			max = arr1[i]
			arr1[i+1] = max
		}
	}
	return arr1
}

/*
 *
 * replace-elements-with-greatest-element-on-right-including-current-index,
 * input: [ 2, 6, 1, 3, 5, 4], output: [6, 6, 5, 5, 5, 4]
 *
 */

func ReplaceElementsAtRightWithCurrentIndexAsWell(arr []int) []int {
	max := arr[len(arr)-1]
	arr1 := make([]int, len(arr))
	arr1[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {

		if arr[i] > arr1[i+1] {
			max = arr[i]
			arr1[i] = max
		} else {
			max = arr1[i+1]
			arr1[i] = max
		}
	}
	return arr1
}

/*
 *
 * replace-elements-with-greatest-element-on-left-including-current-index,
 * input: [ 2 6 1 3 5 4 ], output: [2 6 6 6 6 6]
 *
 */

func ReplaceElementsAtLeftWithCurrentIndexAsWell(arr []int) []int {
	max := arr[0]
	arr1 := make([]int, len(arr))
	arr1[0] = arr[0]
	for i := 0; i < len(arr); i++ {

		if i < len(arr)-1 {

			if arr[i+1] > arr1[i] {
				max = arr[i+1]
				arr1[i+1] = max
			} else {
				max = arr1[i]
				arr1[i+1] = max
			}
		}
	}
	return arr1
}
