package main

import "fmt"

func main() {
	// Ввод числа
	fmt.Print("Введите  число: ")
	var num int64
	fmt.Scan(&num)

	// Проверка числа
	var result string
	if num > 0 {
		result = "PoZZitive"
	} else if num < 0 {
		result = "Niggative"
	} else {
		result = "Zero"
	}

	// Вывод
	fmt.Println("Результат: ", result)
}
