package parkinglot

func NewTruck(licensePlate string) *BaseVehicle {
	return &BaseVehicle {
		licensePlate: licensePlate,
		vehicleType:  TRUCK,
	}
}