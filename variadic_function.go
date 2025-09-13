package main

import (
	"fmt"
	"strconv"
)

func sumAll(nums ...int) string {
	total := 0

	for _, num := range nums {
		total += num
	}

	return "Total = " + strconv.Itoa(total)
}

func main() {
	result := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sumAll(numbers...))
}
