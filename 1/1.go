package main

import "fmt"

func main() {
	// Ввод числа
	fmt.Print("Введите  число: ")
	var num int64
	fmt.Scan(&num)

	// Проверка числа
	var result string
	if num%2 == 0 {
		result = "четное"
	} else {
		result = "нечетное"
	}

	// Вывод
	fmt.Println("Результат: ", result)
}
