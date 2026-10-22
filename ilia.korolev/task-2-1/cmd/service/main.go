package main

import (
	"fmt"
)

func main() {
	var (
		N, employees, T int
		operator        string
	)

	if _, err := fmt.Scan(&N); err != nil {
		fmt.Println("Invalid input N")

		return
	}

	if !(N >= 1 && N <= 1000) {
		fmt.Println("Invalid input N")

		return
	}

	for range N {
		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Println("Invalid input K")

			return
		}

		maxK := 30
		minK := 15

		for range employees {
			if _, err := fmt.Scan(&operator, &T); err != nil {
				fmt.Println("Invalid input")

				return
			}

			if operator == ">=" {
				if T > minK {
					minK = T
				}
			}

			if operator == "<=" {
				if T < maxK {
					maxK = T
				}
			}

			if minK <= maxK {
				fmt.Println(minK)
			} else {
				fmt.Println("-1")
			}
		}
	}
}
