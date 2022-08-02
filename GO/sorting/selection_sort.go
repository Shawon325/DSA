package main

import "fmt"

func main() {
	numbers := []int{99, 44, 0, 90, 6, 2, 1, 44, 5, 63, 87, 283, 4, 0}

	selectionSort(numbers)

	fmt.Println(numbers)
}

func selectionSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		min := i
		temp := numbers[i]

		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[min] {
				min = j
			}
		}

		numbers[i] = numbers[min]
		numbers[min] = temp
	}
}
