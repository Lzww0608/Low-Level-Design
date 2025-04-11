#include "Vehicle.h"
#include <iostream>


Vehicle::Vehicle(std::string licensePlate, VehicleType type, std::string color)
    : licensePlate(licensePlate), type(type), color(color) {}


std::string Vehicle::getLicensePlate() const {
    return licensePlate;
}

VehicleType Vehicle::getType() const {
    return type;
}

std::string Vehicle::getColor() const {
    return color;
}

void Vehicle::displayInfo() const {
    std::cout << "Vehicle: " << color << " ";
    switch (type) {
        case VehicleType::CAR:
            std::cout << "Car";
            break;
        case VehicleType::MOTORCYCLE:
            std::cout << "Motorcycle";
            break;
        case VehicleType::TRUCK:
            std::cout << "Truck";
            break;
        case VehicleType::BUS:
            std::cout << "Bus";
            break;
        default:
            std::cout << "Unknown";
            break;
    }
    std::cout << " (License Plate: " << licensePlate << ")" << std::endl;
}