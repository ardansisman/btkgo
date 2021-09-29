package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3) //make fonksiyonu ile dizi oluşturulduunda slice olur. Slice'ın farkı yeni eleman eklendiğinde dizinin boyutu otomatik olarak artacak
	fmt.Println(isimler)
	isimler[0] = "ardan"
	isimler[1] = "ardana"
	isimler[2] = "ardanb"
	isimler = append(isimler, "ardanc") //Burada 3 elemanlı slice'a bir eleman dah eklendi
	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
