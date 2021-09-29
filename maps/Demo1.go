package maps

import "fmt"

func Demo1() {
	//key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"

	sozluk2 := map[string]string{"glass": "bardak", "red": "kırmızı"}
	fmt.Println(sozluk)
	fmt.Println(sozluk2)
}
