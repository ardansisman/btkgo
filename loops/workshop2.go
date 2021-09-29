package loops

import "fmt"

func Demo3() {
	a := 0
	fmt.Println("Bir sayı giriniz:")
	fmt.Scanln(&a)
	var cntrl int
	for i := 2; i < a; i++ {
		if a%i == 0 {
			cntrl++
		}
	}
	if cntrl == 0 {
		fmt.Println("Girdiğiniz sayı asaldır.")
	}

}
