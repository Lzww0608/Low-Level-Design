package main

import "fmt"

type Cashier struct {
	next Department
}


func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment already done")
		return
	}
	fmt.Println("Cashier getting money from patient")
	p.paymentDone = true
	fmt.Println("Payment done")
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
