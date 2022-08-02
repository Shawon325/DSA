package main

import "fmt"

func main() {
	numbers := []int{99, 44, 0, 90, 6, 2, 1, 44, 5, 63, 87, 283, 4, 0}

	bubbleSort(numbers)

	fmt.Println(numbers)
}

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if j < len(numbers)-1 && numbers[j] > numbers[j+1] {
				temp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
			}
		}
	}
}
