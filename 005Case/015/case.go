// Мастям игральных карт присвоены порядковые номера: 1 — пики,
// 2 — трефы, 3 — бубны, 4 — червы. Достоинству карт, старших десятки,
// присвоены номера: 11 — валет, 12 — дама, 13 — король, 14 — туз. Даны
// два целых числа: N — достоинство (6 ≤ N ≤ 14) и M — масть карты
// (1 ≤ M ≤ 4). Вывести название соответствующей карты вида «шестерка бу-
// бен», «дама червей», «туз треф» и т. п.
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		m, n   uint
		m1, n1 string
	)

	for m < 1 || m > 4 {
		m = ioutil.UInteger("масть")
	}
	for n <= 6 || n >= 14 {
		n = ioutil.UInteger("достоинство")
	}

	switch {
	case m == 1: // пики
		m1 = "пик"
	case m == 2: // трефы
		m1 = "треф"
	case m == 3: // бубны
		m1 = "бубен"
	case m == 4: // червы
		m1 = "червей"
	}

	switch {
	case n == 6: // шестерка
		n1 = "шестерка"
	case n == 7: // семерка
		n1 = "семерка"
	case n == 8: // восьмерка
		n1 = "восьмерка"
	case n == 9: // девятка
		n1 = "девятка"
	case n == 10: // десятка
		n1 = "десятка"
	case n == 11: // валет
		n1 = "валет"
	case n == 12: // дама
		n1 = "дама"
	case n == 13: // король
		n1 = "король"
	case n == 14: // туз
		n1 = "туз"
	}

	fmt.Printf("%v %v\n", n1, m1)
}
