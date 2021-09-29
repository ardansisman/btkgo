package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "Aksaray"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])

	}
	for _, sehir := range sehirler { //_ yerine i yazılabilir i indexi verir mesele şehir = İstanbul ise i=1 olur
		fmt.Println(sehir)
	}
	for i, sehir := range sehirler {
		fmt.Print(i)
		fmt.Println(sehir)

	}
}
