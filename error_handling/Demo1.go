package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt") //dosya bulunursa f ye atanıyor , err'ye nil atanıyor
	if err != nil {
		fmt.Println("Dosya bulunamadı.")
	} else {
		fmt.Println("Dosya bulundu.Dosya adı:", f.Name())
	}

}
