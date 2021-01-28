package main

import (
	"fmt"
	"github.com/rprajapati0067/datastructures-algorithims-using-go/warmups"
)

func main() {
	// Minimum number in array
	fmt.Printf("Minumum Number: %d \n", warmups.MinNumber([]int{2, 1, 3, 6, 4, 8}))

	// Maximum number in array
	fmt.Printf("Maximum Number: %d \n", warmups.MaxNumber([]int{2, 1, 3, 6, 4, 8}))

	//  Minimum and Maximum number in array
	fmt.Printf("Minimum and Maximum: %s \n", warmups.MinMaxNumber([]int{2, 1, 3, 6, 4, 8}))

	//  Sum of two number equal to k
	fmt.Printf("Pairs: %v \n", warmups.SumOfTwoNumbers([]int{2, 1, 3, 6, 4, 8}, 9))
	//  sort color Leet code
	warmups.SortColors([]int{2, 0, 2, 1, 1, 1, 0, 2, 2})

	//  Arrange zero one and two
	result := warmups.ArrangeZeroOneTwo([]int{2, 0, 2, 1, 0, 1, 0})
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
}
