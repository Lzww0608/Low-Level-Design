#ifndef PARKINGLOT_H
#define PARKINGLOT_H

#include "ParkingSpot.h"
#include <vector>
#include <unordered_map>
#include <string>

class ParkingLot {
private:
    std::vector<ParkingSpot*> spots;
    std::unordered_map<std::string, ParkingSpot*> occupiedSpots;
    int capacity;
    int availableSpots;

public:
    ParkingLot(int numCompact, int numRegular, int numLarge);
    ~ParkingLot();

    int getCapacity() const;
    int getAvailableSpots() const;

    bool parkVehicle(Vehicle* vehicle);
    Vehicle* removeVehicle(const std::string& licensePlate);
    ParkingSpot* findVehicle(const std::string& licensePlate) const;

    void displayInfo() const;
    void displayOccupancy() const;

private:
    ParkingSpot* findAvailableSpot(const Vehicle* vehicle) const;
};

#endif
