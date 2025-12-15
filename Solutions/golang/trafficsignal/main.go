package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Traffic Signal Control System Demo ===")
	fmt.Println()

	// Create traffic controller using Singleton pattern
	controller := GetInstance("MainController")

	// Create roads
	road1 := NewRoad("North-South")
	road2 := NewRoad("East-West")

	// Create signals for each road
	signal1 := NewSignal("NS-Signal-1", 30, 5, 25)
	signal2 := NewSignal("EW-Signal-1", 25, 5, 30)

	// Add signals to roads
	road1.AddSignal(signal1)
	road2.AddSignal(signal2)

	// Add roads to controller
	controller.AddRoad(road1)
	controller.AddRoad(road2)

	// Start all signals
	controller.StartAllSignals()
	fmt.Println()

	// Initial status
	controller.PrintStatus()

	// Simulate traffic signal updates
	fmt.Println("\n--- Simulating 10 seconds elapsed ---")
	for _, road := range controller.GetRoads() {
		for _, signal := range road.GetSignals() {
			signal.UpdateSignal(10)
		}
	}
	controller.PrintStatus()

	// Simulate another 20 seconds (should trigger signal switch)
	fmt.Println("\n--- Simulating another 20 seconds elapsed ---")
	for _, road := range controller.GetRoads() {
		for _, signal := range road.GetSignals() {
			signal.UpdateSignal(20)
		}
	}
	controller.PrintStatus()

	// Emergency situation simulation
	fmt.Println("\n--- Emergency Vehicle Approaching on North-South Road! ---")
	controller.HandleEmergency("North-South", "NS-Signal-1")
	controller.PrintStatus()

	// Wait and clear emergency
	time.Sleep(2 * time.Second)
	fmt.Println("\n--- Emergency Cleared ---")
	controller.ClearEmergency("North-South", "NS-Signal-1")
	controller.PrintStatus()

	// Synchronize all signals
	fmt.Println("\n--- Synchronizing All Signals ---")
	controller.SynchronizeRoads()
	controller.PrintStatus()

	// Stop all signals
	fmt.Println("\n--- Stopping All Signals ---")
	controller.StopAllSignals()
	controller.PrintStatus()

	fmt.Println("\n=== Demo Complete ===")
}
