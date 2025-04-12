package parkinglot

type ParkingLot struct {
	levels []*Level
}

var instance *ParkingLot

func NewParkingLot(numLevels int) *ParkingLot {
	if instance != nil {
		return instance
	}

	instance = &ParkingLot{
		levels: make([]*Level, 0, numLevels),
	}

	return instance
}

func (p *ParkingLot) AddLevel(level *Level) {
	p.levels = append(p.levels, level)
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	for _, level := range p.levels {
		if level.ParkVehicle(vehicle) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) UnparkVehicle(spotNumber int) bool {
	for _, level := range p.levels {
		if level.UnparkVehicle(spotNumber) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) DisplayAvailability() {
	for _, level := range p.levels {
		level.DisplayAvailability()
	}
}