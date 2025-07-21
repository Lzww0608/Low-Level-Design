#ifndef TRAFFIC_SYSTEM_H
#define TRAFFIC_SYSTEM_H

#include "Intersection.h"

#include <string>
#include <vector>

class TrafficSystem {
public:
    TrafficSystem(const std::string& id);
    ~TrafficSystem();

    std::string getSystemId() const;
    bool isWorking() const;
    void setWorking(bool working);

    Intersection* createIntersection();
    void removeIntersection(const std::string& intersectionId);
    Signal* addSignal(const std::string& intersectionId,
                     int greenDuration = 30, int yellowDuration = 5, int redDuration = 30);
    void removeSignal(const std::string& intersectionId, const std::string& signalId);
    void updateSystem(int timeElapsed);
    void setIntersectionStatus(const std::string& intersectionId, bool operational);
    void synchronizeIntersection(const std::string& intersectionId);
    void displaySystemStatus() const;
    
private:
    std::string systemId;
    std::vector<Intersection*> intersections;
    int intersectionCount;
    int signalCount;
    bool isWorking_;

    Intersection* findIntersection(const std::string& id) const;
    std::string generateIntersectionId();
    std::string generateSignalId();
};

#endif