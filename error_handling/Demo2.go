package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	if tahmin < 0 || tahmin > 100 {
		return "", errors.New("Girilen sayı 1-100 aralığında olmalıdır")
	}

	return "Bildiniz", nil
}

func Demo2() {

	fmt.Println(TahminEt(50))
	fmt.Println(TahminEt(101))
	fmt.Println(TahminEt(-1))
}
