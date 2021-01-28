package main

import (
	"fmt"
	"github.com/rprajapati0067/datastructures-algorithims-using-go/arrays"
)

func main() {

	// Arrange a greater to right and smaller in left of the given number k
	result := arrays.Arrange([]int{12, 2, 3, 1, 8, 123, 7, 11, 5, 4, 40, 100}, 7)
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	// replace-elements-with-greatest-element-on-right
	result1 := arrays.ReplaceElementsAtRight([]int{17, 18, 5, 4, 6, 1})
	for _, v := range result1 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	// replace-elements-with-greatest-element-on-left
	result3 := arrays.ReplaceElementsAtLeft([]int{2, 6, 1, 3, 5, 4})
	for _, v := range result3 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	// replace-elements-with-greatest-element-on-right-including-current-index
	result4 := arrays.ReplaceElementsAtRightWithCurrentIndexAsWell([]int{2, 6, 1, 3, 5, 4})
	for _, v := range result4 {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	// replace-elements-with-greatest-element-on-left-including-current-index
	result5 := arrays.ReplaceElementsAtLeftWithCurrentIndexAsWell([]int{2, 6, 1, 3, 5, 4})
	for _, v := range result5 {
		fmt.Printf("%d ", v)
	}
}
