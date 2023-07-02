package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	ar := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ar[i])
	}

	zero := 0
	product := 1
	for i := 0; i < n; i++ {
		if ar[i] == 0 {
			zero++
		} else {
			product *= ar[i]
		}
	}

	for i := 0; i < n; i++ {
		if ar[i] == 0 && zero == 1 {
			ar[i] = product
		} else if ar[i] == 0 && zero > 1 {
			ar[i] = 0
		} else if ar[i] != 0 && zero > 0 {
			ar[i] = 0
		} else {
			if ar[i] == 0 {
				ar[i] = product
			} else {
				ar[i] = product / ar[i]
			}
		}
	}

	fmt.Println(ar)
}
