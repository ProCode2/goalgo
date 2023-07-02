package main

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ar[i])
	}

	maxi := math.MinInt
	sum := 0
	for i := 0; i < n; i++ {
		sum += ar[i]
		if sum < 0 {
			sum = 0
		}
		maxi = max(maxi, sum)
	}
	fmt.Printf("%d is the max sum", maxi)
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}
