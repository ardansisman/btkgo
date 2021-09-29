package defer_statement

import "fmt"

type product struct {
	name      string
	unitPrice float64
}

func (p product) Save() {
	fmt.Println("Ürün kaydedildi:", p.name)
	defer Log()
}

func Log() {
	fmt.Println("Loglama işlemi yapıldı")
}

func Demo3() {
	p := product{name: "Laptop", unitPrice: 5300}
	defer p.Save()                            //Ezmeden önce defer stack'e p.Save fonksiyonu için p'yi Laptop olarak atıyo. Dolayısıyla aşağıda ezmiş olsak bile Laptop kaydedilir
	p = product{name: "Mouse", unitPrice: 45} //İlk kaydı eziyoruz

	fmt.Println("İşlem başarılı")
	fmt.Println(p.name)
}
