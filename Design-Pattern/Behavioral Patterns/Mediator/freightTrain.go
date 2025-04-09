package main

import "fmt"

type FreightTrain struct {
	mediator Mediator
}

func (f *FreightTrain) arrive() {	
	if !f.mediator.canArrive(f) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (f *FreightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	f.mediator.notifyAboutDeparture()
}

func (f *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted, arriving")
	f.arrive()
}
