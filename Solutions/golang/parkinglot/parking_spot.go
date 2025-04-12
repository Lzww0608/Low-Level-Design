package parkinglot

type ParkingSpot struct {
	spotNumber    int
	vehicleType   VehicleType
	parkedVehicle Vehicle
}

func NewParkingSpot(spotNumber int, vehicleType VehicleType) *ParkingSpot {
	return &ParkingSpot{
		spotNumber:  spotNumber,
		vehicleType: vehicleType,
	}
}

func (s *ParkingSpot) IsAvailable() bool {
	return s.parkedVehicle == nil
}

func (s *ParkingSpot) ParkVehicle(vehicle Vehicle) bool {
	if !s.IsAvailable() {
		return false
	}
	
	// 规则：
	// 1. 摩托车可以停在任何类型的停车位
	// 2. 汽车只能停在汽车停车位
	// 3. 卡车只能停在卡车停车位
	
	vType := vehicle.GetType()
	switch vType {
	case MOTORCYCLE:
		// 摩托车可以停在任何地方
		s.parkedVehicle = vehicle
		return true
	case CAR:
		// 汽车只能停在汽车或更大的车位
		if s.vehicleType == CAR || s.vehicleType == TRUCK {
			s.parkedVehicle = vehicle
			return true
		}
	case TRUCK:
		// 卡车只能停在卡车车位
		if s.vehicleType == TRUCK {
			s.parkedVehicle = vehicle
			return true
		}
	}
	
	return false
}

func (s *ParkingSpot) UnparkVehicle() bool {
	if s.IsAvailable() {
		return false
	}
	s.parkedVehicle = nil
	return true
}

func (s *ParkingSpot) GetSpotNumber() int {
	return s.spotNumber
}

func (s *ParkingSpot) GetVehicleType() VehicleType {
	return s.vehicleType
}

func (s *ParkingSpot) GetParkedVehicle() Vehicle {
	return s.parkedVehicle
}
