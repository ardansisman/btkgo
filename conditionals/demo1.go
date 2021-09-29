package conditionals

import "fmt"

func Demo1() {
	hesap := 1000
	cekilmekIstenen := 900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz")
	}
	if cekilmekIstenen <= hesap {
		fmt.Println("Paranız hazırlanıyor")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Println("Hesaptaki para :" + fmt.Sprintf("%v", hesap))
	//fmt.Printf("Hesaptaki para : %v", hesap)

}
