package parkinglot

// VehicleType 定义车辆类型
type VehicleType int

const (
	CAR VehicleType = iota
	MOTORCYCLE
	TRUCK
)

// Vehicle 接口定义车辆的基本行为
type Vehicle interface {
	GetLicensePlate() string
	GetType() VehicleType
}

// BaseVehicle 提供Vehicle接口的基本实现
type BaseVehicle struct {
	licensePlate string
	vehicleType  VehicleType
}

func (v *BaseVehicle) GetLicensePlate() string {
	return v.licensePlate
}

func (v *BaseVehicle) GetType() VehicleType {
	return v.vehicleType
}