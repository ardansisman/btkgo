package conditionals

import "fmt"

func Demo3() {
	var sayi1, sayi2, sayi3 int = 62, 12, 39

	buyukSayi := sayi1
	if sayi2 > buyukSayi {
		buyukSayi = sayi2
	}
	if sayi3 > buyukSayi {
		fmt.Print("En büyük sayi ", sayi3)
	} else {
		fmt.Print("En büyük sayi ", buyukSayi)

	}

}
