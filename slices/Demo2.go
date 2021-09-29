package slices

import "fmt"

func Demo2() {
	sehirler := []string{"Aksaray", "Ankara", "İstanbul"} // == sehirler :=make([]string,3)
	sehirler = append(sehirler, "Sivas")
	fmt.Println(sehirler)

	sehirlerKopya := make([]string, len(sehirler))
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler[2:3]) //2 indisten 3. indise kadarını alır
	fmt.Println(sehirler[2:])  // 2den başlar sonrasını alır
	fmt.Println(sehirler[:2])  //2ye kadar olanı alır

}
