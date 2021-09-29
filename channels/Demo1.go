package channels

import (
	"fmt"
	"time"
)

func CiftSayilar(cifSayiCn chan int) { //chan =>channel  int ise toplam işlemi olduğu için
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam += i
		fmt.Println("çift sayı çalışıyor")
		time.Sleep(1 * time.Second)
	}

	cifSayiCn <- toplam //channel değişkenine atama yapılıyor
}

func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam += i
		fmt.Println("tek sayı çalışıyor")
		time.Sleep(1 * time.Second)
	}

	tekSayiCn <- toplam
}
