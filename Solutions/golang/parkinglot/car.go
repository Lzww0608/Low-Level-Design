package parkinglot

func NewCar(licensePlate string) *BaseVehicle {
	return &BaseVehicle {
		licensePlate: licensePlate,
		vehicleType:  CAR,
	}
}