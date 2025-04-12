package parkinglot

func NewMotorcycle(licensePlate string) *BaseVehicle {
	return &BaseVehicle {
		licensePlate: licensePlate,
		vehicleType:  MOTORCYCLE,
	}
}