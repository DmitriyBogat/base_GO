package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a, b int
		c    float64
	)
	//задание 1
	fmt.Print("Введите длину и ширину прямоугольника:")
	fmt.Scanln(&a, &b)
	fmt.Print("Площадь прямоугольника равна: ")
	fmt.Println(a * b)
	//задание 2
	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&c)
	fmt.Print("Диаметр окружности равен: ")
	fmt.Println(2 * math.Sqrt(c/math.Pi))
	fmt.Print("Длина окружности равна: ")
	fmt.Println(2 * math.Pi * math.Sqrt(c/math.Pi))
	//задание 3
	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&a)
	fmt.Print(a / 100)
	b = a % 100
	fmt.Println(" сотен в сием числе")
	fmt.Print(b / 10)
	fmt.Println(" десятков в сием числе")
	fmt.Print(a % 10)
	fmt.Println(" единиц в сием числе")

}
