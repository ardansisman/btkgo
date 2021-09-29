package interfaces

import "fmt"

type mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type creditCalculator interface {
	calculate() float64
}

func (m mortgage) calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c car) calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func calculateMonthlyPayment(credits []creditCalculator) float64 {
	var paymentTotal float64

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].calculate()
	}
	return paymentTotal
}

func Demo2() {
	credit1 := mortgage{rate: 10, creditPaymentTotal: 100000, address: "Ankara"}
	credit2 := mortgage{rate: 12, creditPaymentTotal: 50000, address: "İstanbul"}
	credit3 := car{rate: 15, creditPaymentTotal: 60000, carInfo: "İstanbul"}

	credits := []creditCalculator{credit1, credit2, credit3}

	total := calculateMonthlyPayment(credits)
	fmt.Println("Toplam Ödeme:", total)
}
