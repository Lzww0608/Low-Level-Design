#ifndef INTERSECTION_H
#define INTERSECTION_H

#include "Signal.h"

#include <string>
#include <vector>

class Intersection {
public:
    Intersection(const std::string& id);
    ~Intersection();

    std::string getIntersectionId() const;
    bool isWorking() const;
    void setWorking(bool working);

    void addSignal(Signal* signal);
    void removeSignal(const std::string& id);
    void updateSignal(const std::string& id, int timeElapsed);
    void synchronizeSignals();
    void printStatus() const;

private:
    std::string intersectionId;
    std::vector<Signal*> signals;
    bool isWorking_;

    Signal* findSignal(const std::string& id);
};

#endif