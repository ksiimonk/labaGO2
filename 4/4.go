package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Ввод строки
	fmt.Print("Введите  строку:")
	var str string
	fmt.Scan(&str)

	// Вывод
	fmt.Println(utf8.RuneCountInString(str))
}
