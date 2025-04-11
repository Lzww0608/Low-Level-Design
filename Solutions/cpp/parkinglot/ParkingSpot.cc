#include "ParkingSpot.h"
#include <iostream>



ParkingSpot::ParkingSpot(int spotNumber, SpotType type): spotNumber(spotNumber), type(type), available(true) {}


int ParkingSpot::getSpotNumber() const {
    return spotNumber;
}

SpotType ParkingSpot::getType() const {
    return type;
}

Vehicle* ParkingSpot::getVehicle() const {
    return vehicle;
}

bool ParkingSpot::isAvailable() const {
    return available;
}


bool ParkingSpot::canFitVehicle(const Vehicle* vehicle) const {
    if (vehicle == nullptr) return false;

    switch (vehicle->getType()) {
        case VehicleType::MOTORCYCLE:
            return true; // can park in any spot
        case VehicleType::CAR:
            return type != SpotType::COMPACT;
        case VehicleType::TRUCK:
        case VehicleType::BUS:
            return type == SpotType::LARGE;
    }

    return false;
}


bool ParkingSpot::parkVehicle(Vehicle* vehicle) {
    if (vehicle == nullptr || !canFitVehicle(vehicle)) return false;

    this->vehicle = vehicle;
    this->available = false;
    return true;
}


Vehicle* ParkingSpot::removeVehicle() {
    if (vehicle == nullptr) return nullptr;

    Vehicle* tmp = vehicle;
    vehicle = nullptr;
    available = true;
    return tmp;
}



void ParkingSpot::displayInfo() const {
    std::cout << "Spot " << spotNumber << " (";
    switch (type) {
        case SpotType::COMPACT: std::cout << "Compact"; break;
        case SpotType::REGULAR: std::cout << "Regular"; break;
        case SpotType::LARGE: std::cout << "Large"; break;
    }
    std::cout << "): " << (available ? "Available" : "Occupied");
    if (vehicle) {
        std::cout << " by ";
        vehicle->displayInfo();
    } else {
        std::cout << std::endl;
    }
}
