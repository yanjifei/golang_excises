package main

import "fmt"

type product struct {
	id       uint32
	price    float32
	quantity uint32
}

type inventory struct {
	listOfPdt []*product
}

func (ivt *inventory) printAll() {
	for _, pdt := range ivt.listOfPdt {
		fmt.Printf("Product: %d, Subtotal: %f\r\n",
			pdt.id, pdt.price*float32(pdt.quantity))
	}
}

func main() {

	pdt1 := product{id: 1, price: 3.5, quantity: 9}

	pdt2 := product{id: 2, price: 4.5, quantity: 5}

	var ivt inventory

	ivt.listOfPdt = append(ivt.listOfPdt, &pdt1)
	ivt.listOfPdt = append(ivt.listOfPdt, &pdt2)

	ivt.printAll()

}
