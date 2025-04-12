package parkinglot

type Level struct {
	floor        int
	parkingSpots []*ParkingSpot
}

func NewLevel(floor, numSpots int) *Level {
	level := &Level{
		floor:        floor,
		parkingSpots: make([]*ParkingSpot, numSpots),
	}

	// 平均分配不同类型的停车位：一半为CAR，四分之一为MOTORCYCLE，四分之一为TRUCK
	for i := 0; i < numSpots; i++ {
		var spotType VehicleType
		if i < numSpots/2 {
			spotType = CAR
		} else if i < numSpots*3/4 {
			spotType = MOTORCYCLE
		} else {
			spotType = TRUCK
		}
		level.parkingSpots[i] = NewParkingSpot(i, spotType)
	}

	return level
}

func (l *Level) ParkVehicle(vehicle Vehicle) bool {
	for i, spot := range l.parkingSpots {
		if spot.IsAvailable() && spot.GetVehicleType() == vehicle.GetType() {
			return l.parkingSpots[i].ParkVehicle(vehicle)
		}
	}

	return false
}

func (l *Level) UnparkVehicle(spotNumber int) bool {
	// 如果传入的是车辆在该层的相对位置
	if spotNumber >= 0 && spotNumber < len(l.parkingSpots) {
		return l.parkingSpots[spotNumber].UnparkVehicle()
	}
	
	// 如果传入的是绝对位置，查找对应的停车位
	for i, spot := range l.parkingSpots {
		if spot.GetSpotNumber() == spotNumber && !spot.IsAvailable() {
			return l.parkingSpots[i].UnparkVehicle()
		}
	}
	
	return false
}

func (l *Level) GetAvailableSpots() int {
	count := 0
	for _, spot := range l.parkingSpots {
		if spot.IsAvailable() {
			count++
		}
	}
	return count
}

func (l *Level) GetVehicleCount() int {
	count := 0
	for _, spot := range l.parkingSpots {
		if !spot.IsAvailable() {
			count++
		}
	}
	return count
}

func (l *Level) GetFloor() int {
	return l.floor
}

func (l *Level) DisplayAvailability() {
	for _, spot := range l.parkingSpots {
		status := "Available"
		if !spot.IsAvailable() {
			status = "Occupied"
		}
		println("Level:", l.floor, "Spot:", spot.GetSpotNumber(), "Status:", status, "Type:", spot.GetVehicleType())
	}
}
