package main

import (
	"github.com/ardansisman/btkgo/project"
)

func main() {
	//variables.Demo1()

	//conditionals.Demo1()
	//conditionals.Demo3()

	//loops.Demo1()
	//loops.Demo2()
	//loops.Demo3()
	//arrays.Demo2()

	//slices.Demo1()
	//slices.Demo2()

	/*functions.Sum(10, 2)
	  functions.Difference(10, 2)
	  functions.Multiplication(10, 2)
	  functions.Division(10, 2)*/

	/*	a, b, c, d := functions.DotIslem(10, 2) // a, b, c, _ := functions.DotIslem(10, 2)  4 tane returnden ihtiyaç duyulmayana _ koyulabiliyor
		fmt.Println("Toplam:", a)
		fmt.Println("Fark:", b)
		fmt.Println("Carpim:", c)
		fmt.Println("Bölüm:", d)*/

	/*a := functions.ToplaVariadic(3, 6, 5, 8, 5, 6, 9, 5)
	fmt.Println(a)

	sayilar := []int{20, 80, 50, 10, 2}

	toplam := functions.ToplaVariadic(sayilar...) //variadic olduğunu bildirmek için ... koyuluyor
	fmt.Println(toplam)*/

	//maps.Demo1()

	//for_range.Demo2()
	//for_range.Demo3()

	/*sayi := 20 //int bool value type olduğu için pointer kullanmadığımızda değeri değişmiyor. Pointer kullanarak sayi değişkenini adresiyle beraber fonksiyona gönderdik dolayısı ile değeri değişip geldi
	pointer_test.Demo1(&sayi)
	fmt.Println("Maindeki Sayi:", sayi)*/

	/*	sayilar := []int{1, 2, 3} //arraylar reference type olduğu için pointer'a gerek kalmadan değeri değiştirebildik
		pointer_test.Demo2(sayilar)
		fmt.Println("Maindeki ilk sayi:", sayilar[0])*/

	//structs.Demo1()
	//structs.Demo2()

	/*go goroutines.CiftSayilar() //go yazmak asenkron olarak çalıştırmamızı sağlıyor
	go goroutines.TekSayilar()
	time.Sleep(5 * time.Second) //5 saniye bekletiyoruz
	fmt.Println("Main bitti")*/

	/*ciftSayiCn := make(chan int) //make yardımıyla channel değişkeni oluşturuyoruz
	tekSayiCn := make(chan int)
	go channels.CiftSayilar(ciftSayiCn)
	go channels.TekSayilar(tekSayiCn)
	ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	//fmt.Println("Channelden gelen Çift sayı:", ciftSayiToplam)
	//fmt.Println("Channelden gelen Tek sayı:", tekSayiToplam)

	//carpim := ciftSayiCn*tekSayiCn bu sayıları çarpamayız çünkü bunlar channel değişkeni içinde sadece değerler var int değil


	carpim := ciftSayiToplam * tekSayiToplam
	println("Çarpımları:", carpim)*/

	//interfaces.Demo1()
	//interfaces.Demo2()
	//interfaces.Demo3()

	//defer_statement.B()
	//defer_statement.Test()
	//defer_statement.Demo3()

	//error_handling.Demo1()
	//error_handling.Demo2()
	//fmt.Println(error_handling.Demo3(102))

	// string_functions.Demo1()
	//string_functions.Demo2()

	// restful.Demo1()
	//restful.Demo2()

	project.AddProduct()
	//project.GetAllProducts()

}
