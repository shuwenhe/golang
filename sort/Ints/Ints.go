package Ints

import (
	"fmt"
	"sort"
)

func Ints(nums []int) {
	fmt.Println("nums = ", nums)
	sort.Ints(nums)
	fmt.Println("sort = ", nums)
}
