package main

import "fmt"

type SignalType int

const (
	RED SignalType = iota
	GREEN
	YELLOW
	UNKNOWN
)

func (s SignalType) String() string {
	switch s {
	case RED:
		return "RED"
	case GREEN:
		return "GREEN"
	case YELLOW:
		return "YELLOW"
	default:
		return "UNKNOWN"
	}
}

type Signal struct {
	signalId          string
	currentSignal     SignalType
	remainingDuration int
	greenDuration     int
	yellowDuration    int
	redDuration       int
	isWorking         bool
	isEmergency       bool
}

func NewSignal(signalId string, greenDuration, yellowDuration, redDuration int) *Signal {
	return &Signal{
		signalId:          signalId,
		currentSignal:     RED,
		remainingDuration: redDuration,
		greenDuration:     greenDuration,
		yellowDuration:    yellowDuration,
		redDuration:       redDuration,
		isWorking:         false,
		isEmergency:       false,
	}
}

func (s *Signal) GetSignalId() string {
	return s.signalId
}

func (s *Signal) GetCurrentSignal() SignalType {
	return s.currentSignal
}

func (s *Signal) GetRemainingDuration() int {
	return s.remainingDuration
}

func (s *Signal) GetGreenDuration() int {
	return s.greenDuration
}

func (s *Signal) GetYellowDuration() int {
	return s.yellowDuration
}

func (s *Signal) GetRedDuration() int {
	return s.redDuration
}

func (s *Signal) GetIsWorking() bool {
	return s.isWorking
}

func (s *Signal) SetSignal(signal SignalType) {
	s.currentSignal = signal
}

func (s *Signal) UpdateSignal(timeElapsed int) {
	if !s.isWorking {
		return
	}
	if s.remainingDuration > timeElapsed {
		s.remainingDuration -= timeElapsed
	} else {
		s.remainingDuration = 0
		s.SwitchSignal()
	}
}

func (s *Signal) SetDuration(greenDuration, yellowDuration, redDuration int) {
	s.greenDuration = greenDuration
	s.yellowDuration = yellowDuration
	s.redDuration = redDuration
	s.remainingDuration = redDuration
}

func (s *Signal) SetWorking(working bool) {
	s.isWorking = working
}

func (s *Signal) Reset() {
	s.currentSignal = RED
	s.remainingDuration = s.redDuration
	s.isWorking = true
	s.isEmergency = false
}

func (s *Signal) SwitchSignal() {
	if !s.isWorking {
		return
	}
	switch s.currentSignal {
	case RED:
		s.currentSignal = GREEN
		s.remainingDuration = s.greenDuration
	case GREEN:
		s.currentSignal = YELLOW
		s.remainingDuration = s.yellowDuration
	case YELLOW:
		s.currentSignal = RED
		s.remainingDuration = s.redDuration
	}
}

func (s *Signal) PrintStatus() {
	fmt.Printf("  Signal[%s] Status: %s, Remaining: %d sec, Working: %t, Emergency: %t\n",
		s.signalId, s.currentSignal, s.remainingDuration, s.isWorking, s.isEmergency)
}

// HandleEmergency sets the signal to GREEN immediately for emergency vehicles
func (s *Signal) HandleEmergency() {
	if !s.isWorking {
		return
	}
	s.isEmergency = true
	s.currentSignal = GREEN
	s.remainingDuration = s.greenDuration
}

// ClearEmergency clears the emergency state
func (s *Signal) ClearEmergency() {
	s.isEmergency = false
}
