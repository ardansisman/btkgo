package structs

import "fmt"

func Demo1() {
	fmt.Println(product{1, "Lenovo", 5200})

	fmt.Println(product{id: 1, name: "Lenovo"})
}

type product struct {
	id        int
	name      string
	unitPrice float64
}
