package arrays

import "fmt"

func Demo2() {
	sayilar := [5]int{20, 80, 50, 10, 2}
	//fmt.Println(sayilar)
	var max int
	max = sayilar[0]
	for i := 1; i < len(sayilar); i++ {

		if sayilar[i] > max {
			max = sayilar[i]
		}

	}
	fmt.Println("En büyük sayı:", max)
}
