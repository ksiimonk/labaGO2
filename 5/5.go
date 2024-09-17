package main

import (
	"fmt"
)

// Структура Rectangle
type Rectangle struct {
	Length, Width float64
}

// Вычисляем площадь прямоугольника.
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func main() {
	// Создаем экземпляр Rectangle
	rect := Rectangle{Length: 10.0, Width: 5.0}

	// Вычисляем площадь и выводим результат
	area := rect.Area()
	fmt.Printf("Площадь прямоугольника: %.2f\n", area)
}
