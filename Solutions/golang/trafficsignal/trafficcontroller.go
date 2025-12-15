package main

import (
	"fmt"
	"sync"
)

type TrafficController struct {
	controllerId string
	roads        []*Road
	roadCount    int
	mu           sync.Mutex
}

var (
	instance *TrafficController
	once     sync.Once
)

// GetInstance returns the singleton instance of TrafficController
func GetInstance(controllerId string) *TrafficController {
	once.Do(func() {
		instance = &TrafficController{
			controllerId: controllerId,
			roads:        make([]*Road, 0),
			roadCount:    0,
		}
	})
	return instance
}

// NewTrafficController creates a new TrafficController (non-singleton)
func NewTrafficController(controllerId string) *TrafficController {
	return &TrafficController{
		controllerId: controllerId,
		roads:        make([]*Road, 0),
		roadCount:    0,
	}
}

func (t *TrafficController) GetControllerId() string {
	return t.controllerId
}

func (t *TrafficController) GetRoads() []*Road {
	return t.roads
}

func (t *TrafficController) GetRoadCount() int {
	return t.roadCount
}

func (t *TrafficController) AddRoad(road *Road) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.roads = append(t.roads, road)
	t.roadCount++
}

func (t *TrafficController) RemoveRoad(roadId string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for i, road := range t.roads {
		if road.GetRoadId() == roadId {
			t.roads = append(t.roads[:i], t.roads[i+1:]...)
			t.roadCount--
		}
	}
}

func (t *TrafficController) UpdateRoad(roadId string, signalId string, timeElapsed int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, road := range t.roads {
		if road.GetRoadId() == roadId {
			road.UpdateSignal(signalId, timeElapsed)
		}
	}
}

func (t *TrafficController) SynchronizeRoads() {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, road := range t.roads {
		road.SynchronizeSignals()
	}
}

func (t *TrafficController) PrintStatus() {
	t.mu.Lock()
	defer t.mu.Unlock()
	fmt.Printf("\n=== Traffic Controller[%s] Status ===\n", t.controllerId)
	for _, road := range t.roads {
		road.PrintStatus()
	}
	fmt.Println("=====================================")
}

// HandleEmergency handles emergency on a specific road and signal
func (t *TrafficController) HandleEmergency(roadId string, signalId string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, road := range t.roads {
		if road.GetRoadId() == roadId {
			road.HandleEmergency(signalId)
			return
		}
	}
}

// ClearEmergency clears emergency state
func (t *TrafficController) ClearEmergency(roadId string, signalId string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, road := range t.roads {
		if road.GetRoadId() == roadId {
			road.ClearEmergency(signalId)
			return
		}
	}
}

// StartAllSignals starts all signals in all roads
func (t *TrafficController) StartAllSignals() {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, road := range t.roads {
		for _, signal := range road.GetSignals() {
			signal.SetWorking(true)
		}
	}
	fmt.Println("All traffic signals started")
}

// StopAllSignals stops all signals in all roads
func (t *TrafficController) StopAllSignals() {
	t.mu.Lock()
	defer t.mu.Unlock()
	for _, road := range t.roads {
		for _, signal := range road.GetSignals() {
			signal.SetWorking(false)
		}
	}
	fmt.Println("All traffic signals stopped")
}
