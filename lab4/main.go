package main

import (
	"fmt"
)

func insert_sort(mass []int) {

	for unsort_side_index := 1; unsort_side_index < len(mass); unsort_side_index++ {
		x := mass[unsort_side_index]
		sort_side_index := unsort_side_index
		for ; sort_side_index >= 1 && mass[sort_side_index-1] > x; sort_side_index-- {
			mass[sort_side_index] = mass[sort_side_index-1]
		}
		mass[sort_side_index] = x
	}
}
func main() {
	var arr_size int
	fmt.Print("Введите размер массива ")
	fmt.Scanln(&arr_size)
	var mass []int
	mass = make([]int, arr_size)
	for i := 0; i < arr_size; i++ {
		if i == 0 {
			fmt.Print("Введите целое число ")
		} else {
			fmt.Print("Еще ")
		}
		fmt.Scanln(&mass[i])
	}
	fmt.Print("неотсортированный массив - ")
	fmt.Println(mass)
	insert_sort(mass)
	fmt.Print("отсортированный массив - ")
	fmt.Println(mass)
}
