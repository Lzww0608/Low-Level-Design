package main

import (
	"fmt"
	"sync"
)

type Road struct {
	roadId      string
	signals     []*Signal
	signalCount int
	mu          sync.Mutex
}

func NewRoad(roadId string) *Road {
	return &Road{
		roadId:      roadId,
		signals:     make([]*Signal, 0),
		signalCount: 0,
	}
}

func (r *Road) GetRoadId() string {
	return r.roadId
}

func (r *Road) GetSignals() []*Signal {
	return r.signals
}

func (r *Road) GetSignalCount() int {
	return r.signalCount
}

func (r *Road) AddSignal(signal *Signal) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.signals = append(r.signals, signal)
	r.signalCount++
}

func (r *Road) RemoveSignal(signalId string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, signal := range r.signals {
		if signal.GetSignalId() == signalId {
			r.signals = append(r.signals[:i], r.signals[i+1:]...)
			r.signalCount--
		}
	}
}

func (r *Road) UpdateSignal(signalId string, timeElapsed int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, signal := range r.signals {
		if signal.GetSignalId() == signalId {
			signal.UpdateSignal(timeElapsed)
		}
	}
}

func (r *Road) SynchronizeSignals() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, signal := range r.signals {
		signal.SwitchSignal()
	}
}

func (r *Road) PrintStatus() {
	fmt.Printf("Road[%s] SignalCount: %d\n", r.roadId, r.signalCount)
	for _, signal := range r.signals {
		signal.PrintStatus()
	}
}

// HandleEmergency handles emergency situations on this road
func (r *Road) HandleEmergency(signalId string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, signal := range r.signals {
		if signal.GetSignalId() == signalId {
			signal.HandleEmergency()
			fmt.Printf("Emergency handled on Road[%s], Signal[%s]\n", r.roadId, signalId)
		}
	}
}

// ClearEmergency clears emergency state
func (r *Road) ClearEmergency(signalId string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, signal := range r.signals {
		if signal.GetSignalId() == signalId {
			signal.ClearEmergency()
		}
	}
}
