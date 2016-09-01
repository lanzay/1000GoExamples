// С начала суток прошло N секунд ( N — целое). Найти количество се-
// кунд, прошедших с начала последнего часа
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x uint
	for x == 0 {
		x = util.UInteger("прошло N секунд ( N — целое)")
	}
	x = x % 3600
	fmt.Printf("количество секунд, прошедших с начала последнего часа: %v\n", x)
}
