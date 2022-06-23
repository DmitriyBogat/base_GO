package main

import (
	"fmt"
)

func fibona_func() func() int64 {
	numbers := make(map[int]int64)
	index := 0
	return func() int64 {
		if index == 0 {
			numbers[index] = 0
			index++
			return 0
		}
		if index == 1 {
			numbers[index] = 1
			index++
			return 1
		}
		number := numbers[index-1] + numbers[index-2]
		numbers[index] = number
		index++
		return number
	}
} // оптимизация здесь в том что записывая числа в мапу мы их повторно не вычисляем в рекурсии.

func main() {
	var fub_number int
	fmt.Print("Введите требуемый номер числа фибоначи: ")
	fmt.Scanln(&fub_number)
	f := fibona_func()
	for i := 1; i < fub_number; i++ {
		f()
	}
	fmt.Println(f())
}
