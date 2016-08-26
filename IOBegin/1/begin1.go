// Дана сторона квадрата a. Найти его периметр P = 4·a.
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x float64
	fmt.Println("Необходимо найти периметр квадрата")
	x = util.Number("Введите значение стороны")
	fmt.Println("Периметр равен:\t", 4*x)
}
