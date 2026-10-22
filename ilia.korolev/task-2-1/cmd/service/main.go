package main

import "fmt"

func main() {
	var (
		departmens, employees, temperature int
		operator                           string
	)

	if _, err := fmt.Scan(&departmens); err != nil {
		fmt.Println("Invalid input departmens")

		return
	}

	if !(departmens >= 1 && departmens <= 1000) {
		fmt.Println("Invalid input departmens")

		return
	}

	for range departmens {
		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Println("Invalid input K")

			return
		}

		maxK := 30
		minK := 15

		for range employees {
			if _, err := fmt.Scan(&operator, &temperature); err != nil {
				fmt.Println("Invalid input")

				return
			}

			if operator == ">=" {
				minK = max(minK, temperature)
			}

			if operator == "<=" {
				maxK = min(maxK, temperature)
			}

			if minK <= maxK {
				fmt.Println(minK)
			} else {
				fmt.Println("-1")
			}
		}
	}
}
