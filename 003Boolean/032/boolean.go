// Даны целые числа a , b , c , являющиеся сторонами некоторого тре-
// угольника. Проверить истинность высказывания: «Треугольник со сторо-
// нами a , b , c является прямоугольным»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var a, b, c uint
	var x bool
	a = util.UInteger("число a")
	b = util.UInteger("число b")
	c = util.UInteger("число c")
	x = (c*c == (a*a + b*b)) || (a*a == (c*c + b*b)) || (b*b == (a*a + c*c))
	fmt.Println(x)
}