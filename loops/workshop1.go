package loops

import "fmt"

func Demo2() {
	var sayi int = 10
	tahminEdilenSayi := 0
	tahminSayisi := 0
	fmt.Println("Bir sayı tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)

	for sayi != tahminEdilenSayi {
		if sayi > tahminEdilenSayi {
			fmt.Println("Daha büyük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi++
		}
		if tahminEdilenSayi > sayi {
			fmt.Println("Daha küçük bir sayı giriniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi++

		}
		if sayi == tahminEdilenSayi {
			fmt.Println("Tebriklerrrrr Sayıyı buldunuz")
			tahminSayisi++
			fmt.Print("Tahmin sayısı ", tahminSayisi)

		}
	}

}
