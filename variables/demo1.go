package variables

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya"
	fmt.Println(metin)
	fmt.Println(metin)

	var kdv int = 10
	fmt.Print(kdv)

	var kdv2 float32 = 0.1
	fmt.Println(kdv2)

	kdv3 := 25.2
	fmt.Println(kdv3)
	fmt.Printf("Veri Tipi: %T\n", kdv3)

	sonuc := false

	metin1 := "Ardan"
	metin2 := "Ardan"

	sonuc = metin1 == metin2

	fmt.Println(sonuc)
}
