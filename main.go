// урок 2 модуль 23 алгоритм сортировки выбором (selection sort)
package main

import "fmt"

const size = 10

func main() {
	a := [size]int64{10, 39, 2, 3, 10, 99, 5, 45, 54, 31}

	fmt.Println(a)

	b := sort(a)
	fmt.Println(b)
}

// напишем сортировку "выбором" (selection sort)
func sort(input [size]int64) [size]int64 {
	for i := 0; i < size; i++ {
		minIndex := i
		for j := i; j < size; j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}
		input[i], input[minIndex] = input[minIndex], input[i]
	}
	return input
}
