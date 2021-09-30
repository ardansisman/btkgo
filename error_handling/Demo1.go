package error_handling

import (
	"fmt"
	"os"
)

//type assertion
func Demo1() {
	f, err := os.Open("error_handling/demo1.txt") //dosya bulunursa f ye atanıyor , err'ye nil atanıyor
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok { //path errorse ok true oluyor
			fmt.Println("Dosya bulunamadı : ", pErr.Path)
			return
		}
		fmt.Println("Dosya bulunamadı.")
		return
	}
	fmt.Println("Dosya bulundu.Dosya adı:", f.Name())
}
