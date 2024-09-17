package main

import "fmt"

func main() {
	// Ввод чисел
	fmt.Print("Введите  числа: ")
	var num1 int
	var num2 int
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	// Вычисление и вывод
	fmt.Printf("Результат: %.2f\n", float64((num1+num2)/2))
}
