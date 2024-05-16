package main

import "fmt"

func findFactors(number int) []int {
	result := make([]int, 0)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	a := findFactors(3419110721)
	fmt.Println(a)
	a = findFactors(341)
	fmt.Println(a)
}
