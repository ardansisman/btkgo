package pointer_test

import "fmt"

func Demo2(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("Demodaki ilk sayı", sayilar[0])
}
