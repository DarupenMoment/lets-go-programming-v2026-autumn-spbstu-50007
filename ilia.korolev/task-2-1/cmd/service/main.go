package main

import "fmt"

func main() {
	var N, K, T int
	var op string
	if _, err := fmt.Scan(&N); err != nil {
		fmt.Println("Invalid input N")
		return
	}
	if !(N >= 1 && N <= 1000) {
		fmt.Println("Invalid input N")
		return
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&K)
		max_K := 30
		min_K := 15
		for j := 0; j < K; j++ {
			fmt.Scan(&op, &T)
			if op == ">=" {
				if T > min_K {
					min_K = T
				}
			}
			if op == "<=" {
				if T < max_K {
					max_K = T
				}
			}
			if min_K <= max_K {
				fmt.Println(min_K)
			} else {
				fmt.Println("-1")
			}

		}
	}

}
