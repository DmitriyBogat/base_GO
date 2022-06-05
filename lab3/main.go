package main

import (
	"fmt"
	"math"
	"os"
)

func factorial(num float64) float64 { ///// Для ДЗ1
	var (
		result float64
	)
	result = 1
	for i := 1.0; i < num; i++ {
		result = result * (i + 1)
	}
	return result // можно было конечно стянуть рекурсию из первой строки поиска google))
} // вообще искал встроенную функцию в модуле math, но разрабы языка походу решили, что
//она нафиг никому не нужна))

func IS_Simple_Number(num int) bool { ///// Для ДЗ2
	var i int
	if num%2 == 0 {
		return false
	}
	i = 3
	for true {
		if num%i == 0 && i != num {
			return false
		}
		i += 2
		if i*i >= num || i >= num {
			break
		}
	}
	return true
}
func main() {
	var a, b, res, c float64
	var op, dz_num string
	var sumple_num int
	fmt.Print("Введите dz ID (dz_1, dz_2): ")
	fmt.Scanln(&dz_num)
	switch dz_num {
	case "dz_1":
		{
			fmt.Print("Введите арифметическую операцию (+, -, *, /,!,pow): ")
			fmt.Scanln(&op)
			if op == "!" {
				fmt.Print("Введите число: ")
				fmt.Scanln(&c)
			}
			{
				fmt.Print("Введите первое число: ")
				fmt.Scanln(&a)
				fmt.Print("Введите второе число: ")
				fmt.Scanln(&b)
			}
			switch op {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			case "!":
				res = factorial(c)
			case "pow":
				res = math.Pow(a, b)
			default:
				fmt.Println("Операция выбрана неверно")
				os.Exit(1)
			}
			fmt.Printf("Результат выполнения операции: %.2f\n", res)
		}
	case "dz_2":
		{
			fmt.Print("Введите число: ")
			fmt.Scanln(&sumple_num)
			if sumple_num == 2 {
				fmt.Println("Простые числа:")
				fmt.Println(2)
			} else {
				fmt.Println("Простые числа:")
				fmt.Println(2)
				for i := 3; i <= sumple_num; i++ {
					if IS_Simple_Number(i) {
						fmt.Println(i)
					}
				}
			}

		}
	}

}
