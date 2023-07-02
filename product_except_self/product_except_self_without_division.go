package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	ar := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ar[i])
	}

	prefix := 1
	postfix := 1

	res := make([]int, n)

	res[0] = 1
	for i := 0; i < n; i++ {
		res[i] = prefix
		prefix *= ar[i]
	}

	for i := n - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= ar[i]
	}

	fmt.Println(res)
}
