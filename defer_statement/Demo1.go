package defer_statement

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalıştı")
}

func C() {
	fmt.Println("C fonksiyonu çalıştı")
}

func D() {
	fmt.Println("D fonksiyonu çalıştı")
}

func B() {
	defer A() //başına defer yazdığımız için bu fonksiyonu en üste bile koysak en son çalışır
	defer C()
	defer D()
	fmt.Println("B fonksiyonu çalıştı")

}
