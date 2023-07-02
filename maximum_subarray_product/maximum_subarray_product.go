package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ar[i])
	}

	max_prod := 1
	min_prod := 1
	res := math.MinInt

	for i := 0; i < n; i++ {
		tmp := ar[i] * max_prod
		max_prod = max([]int{tmp, ar[i] * min_prod, ar[i]})
		min_prod = min([]int{tmp, ar[i] * min_prod, ar[i]})
		res = max([]int{res, max_prod})
	}
	fmt.Printf("%d is the max subarray product", res)
}
func min(values []int) int {
	min := math.MaxInt
	n := len(values)
	for i := 0; i < n; i++ {
		if values[i] < min {
			min = values[i]
		}
	}

	return min
}

func max(values []int) int {
	max := math.MinInt
	n := len(values)
	for i := 0; i < n; i++ {
		if values[i] > max {
			max = values[i]
		}
	}

	return max
}
