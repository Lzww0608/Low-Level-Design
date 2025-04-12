package parkinglot

import (
	"testing"
)

// 测试创建停车场
func TestNewParkingLot(t *testing.T) {
	parkingLot := NewParkingLot(3)
	if parkingLot == nil {
		t.Error("停车场创建失败")
	}
	
	// 测试单例模式
	anotherParkingLot := NewParkingLot(5)
	if parkingLot != anotherParkingLot {
		t.Error("单例模式实现失败")
	}
}

// 测试创建停车层
func TestNewLevel(t *testing.T) {
	level := NewLevel(1, 10)
	if level == nil {
		t.Error("停车层创建失败")
	}
	
	if level.GetFloor() != 1 {
		t.Errorf("楼层号错误，期望：1，实际：%d", level.GetFloor())
	}
	
	if len(level.parkingSpots) != 10 {
		t.Errorf("停车位数量错误，期望：10，实际：%d", len(level.parkingSpots))
	}
}

// 模拟车辆类型
type MockCar struct {
	BaseVehicle
}

func NewMockCar(licensePlate string) *MockCar {
	return &MockCar{
		BaseVehicle: BaseVehicle{
			licensePlate: licensePlate,
			vehicleType:  CAR,
		},
	}
}

type MockMotorcycle struct {
	BaseVehicle
}

func NewMockMotorcycle(licensePlate string) *MockMotorcycle {
	return &MockMotorcycle{
		BaseVehicle: BaseVehicle{
			licensePlate: licensePlate,
			vehicleType:  MOTORCYCLE,
		},
	}
}

// 测试停车和取车功能
func TestParkAndUnparkVehicle(t *testing.T) {
	parkingLot := NewParkingLot(1)
	level := NewLevel(1, 10)
	parkingLot.AddLevel(level)
	
	car := NewMockCar("京A12345")
	motorcycle := NewMockMotorcycle("京B12345")
	
	// 测试停车
	if !parkingLot.ParkVehicle(car) {
		t.Error("汽车停车失败")
	}
	
	if !parkingLot.ParkVehicle(motorcycle) {
		t.Error("摩托车停车失败")
	}
	
	// 验证车辆数量
	if level.GetVehicleCount() != 2 {
		t.Errorf("车辆数量错误，期望：2，实际：%d", level.GetVehicleCount())
	}
	
	// 测试取车
	carSpot := 0
	if !parkingLot.UnparkVehicle(carSpot) {
		t.Error("取车失败")
	}
	
	// 验证车辆数量
	if level.GetVehicleCount() != 1 {
		t.Errorf("取车后车辆数量错误，期望：1，实际：%d", level.GetVehicleCount())
	}
}

// 测试停车场可用性
func TestParkingLotAvailability(t *testing.T) {
	// 重置单例，避免之前测试的影响
	instance = nil
	
	// 创建一个只有CAR类型停车位的停车场
	parkingLot := NewParkingLot(1)
	level := &Level{
		floor:        1,
		parkingSpots: make([]*ParkingSpot, 3),
	}
	
	// 创建3个CAR类型的停车位
	for i := 0; i < 3; i++ {
		level.parkingSpots[i] = NewParkingSpot(i, CAR)
	}
	
	parkingLot.AddLevel(level)
	
	initialAvailable := level.GetAvailableSpots()
	if initialAvailable != 3 {
		t.Errorf("初始可用停车位数量错误，期望：3，实际：%d", initialAvailable)
	}
	
	// 停三辆车
	car1 := NewMockCar("京A11111")
	car2 := NewMockCar("京A22222")
	car3 := NewMockCar("京A33333")
	
	parkingLot.ParkVehicle(car1)
	parkingLot.ParkVehicle(car2)
	parkingLot.ParkVehicle(car3)
	
	// 验证没有可用停车位
	if level.GetAvailableSpots() != 0 {
		t.Errorf("停满车后可用停车位数量错误，期望：0，实际：%d", level.GetAvailableSpots())
	}
	
	// 尝试再停一辆车，应该失败
	car4 := NewMockCar("京A44444")
	if parkingLot.ParkVehicle(car4) {
		t.Error("停车场已满，但仍然可以停车")
	}
	
	// 取一辆车，然后再停车
	parkingLot.UnparkVehicle(0)
	if level.GetAvailableSpots() != 1 {
		t.Errorf("取车后可用停车位数量错误，期望：1，实际：%d", level.GetAvailableSpots())
	}
	
	if !parkingLot.ParkVehicle(car4) {
		t.Error("有可用停车位，但停车失败")
	}
}
