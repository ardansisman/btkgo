package string_functions

import (
	"fmt"
	s "strings" //alias strings yerine s kullanacağız
)

func Demo1() {
	isim := "ardan"
	fmt.Println(s.Count(isim, "a"))
	fmt.Println(s.Contains(isim, "a"))
	fmt.Println(s.Index(isim, "d"))
	fmt.Println(s.Index(isim, "z"))

}
