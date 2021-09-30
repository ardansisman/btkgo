package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Ardan"
	fmt.Println(s.HasPrefix(isim, "Ar"))
	fmt.Println(s.HasSuffix(isim, "an"))

	harfler := []string{"a", "r", "d", "a", "n"}
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", -1))
	fmt.Println(s.Replace(sonuc, "*", "+", 1))
	fmt.Println(s.Replace(sonuc, "*", "+", 2))

	fmt.Println(s.Split(sonuc, "*"))
	fmt.Println(s.Split(sonuc, "-"))

	fmt.Println(s.Repeat(sonuc, 5))

}
