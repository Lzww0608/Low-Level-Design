package main

import "fmt"

type Mac struct {}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lighting connector is plugged into mac machine.")	
}