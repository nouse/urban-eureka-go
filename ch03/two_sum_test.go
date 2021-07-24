package ch03

import (
	"fmt"
)

func ExampleTwoSum()  {
	fmt.Println(TwoSum([]int{2, 7, 10, 0}, 9))
	fmt.Println(TwoSum([]int{2, 3, 5, 0, 1}, 6))
	// Output:
	// [0 1]
	// [2 4]
}
