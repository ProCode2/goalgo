package main

import "fmt"

func main() {

	var n int

	fmt.Scanf("%d", &n)

	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ar[i])
	}

	exists := make(map[int]bool)

	for i := 0; i < n; i++ {
		if _, ok := exists[ar[i]]; ok {
			fmt.Printf("%d exists at least twice.\n", ar[i])

			return
		}

		exists[ar[i]] = true
	}

	fmt.Println("No values exists twice.")

}
