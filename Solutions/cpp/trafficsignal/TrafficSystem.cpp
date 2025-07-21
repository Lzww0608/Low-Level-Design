#include "TrafficSystem.h"
#include "Intersection.h"
#include "Signal.h"
#include <iostream>
#include <algorithm>
#include <sstream>

TrafficSystem::TrafficSystem(const std::string& id)
    : systemId(id), intersectionCount(0), signalCount(0), isWorking_(true) {}

TrafficSystem::~TrafficSystem() {
    for (Intersection* intersection : intersections) {
        delete intersection;
    }
    intersections.clear();
}

std::string TrafficSystem::getSystemId() const {
    return systemId;
}

bool TrafficSystem::isWorking() const {
    return isWorking_;
}

void TrafficSystem::setWorking(bool working) {
    isWorking_ = working;
    // 同步所有路口的工作状态
    for (Intersection* intersection : intersections) {
        if (intersection) {
            intersection->setWorking(working);
        }
    }
}

Intersection* TrafficSystem::createIntersection() {
    std::string id = generateIntersectionId();
    Intersection* intersection = new Intersection(id);
    intersections.push_back(intersection);
    ++intersectionCount;
    return intersection;
}

void TrafficSystem::removeIntersection(const std::string& intersectionId) {
    auto it = std::remove_if(intersections.begin(), intersections.end(),
        [&](Intersection* intersection) {
            if (intersection && intersection->getIntersectionId() == intersectionId) {
                delete intersection;
                --intersectionCount;
                return true;
            }
            return false;
        });
    intersections.erase(it, intersections.end());
}

Signal* TrafficSystem::addSignal(const std::string& intersectionId,
                                 int greenDuration, int yellowDuration, int redDuration) {
    Intersection* intersection = findIntersection(intersectionId);
    if (!intersection) return nullptr;
    std::string signalId = generateSignalId();
    Signal* signal = new Signal(signalId, greenDuration, yellowDuration, redDuration);
    intersection->addSignal(signal);
    ++signalCount;
    return signal;
}

void TrafficSystem::removeSignal(const std::string& intersectionId, const std::string& signalId) {
    Intersection* intersection = findIntersection(intersectionId);
    if (intersection) {
        intersection->removeSignal(signalId);
        --signalCount;
    }
}

void TrafficSystem::updateSystem(int timeElapsed) {
    for (Intersection* intersection : intersections) {
        if (intersection && intersection->isWorking()) {
            // 遍历每个路口的信号灯并更新
            // Intersection没有暴露信号灯列表，只能通过信号灯id逐个更新
            // 这里假设Intersection内部有updateSignal方法能批量处理
            // 实际上应在Intersection中实现批量更新所有信号灯的方法
            // 这里简单处理：同步所有信号灯
            intersection->synchronizeSignals();
        }
    }
}

void TrafficSystem::setIntersectionStatus(const std::string& intersectionId, bool operational) {
    Intersection* intersection = findIntersection(intersectionId);
    if (intersection) {
        intersection->setWorking(operational);
    }
}

void TrafficSystem::synchronizeIntersection(const std::string& intersectionId) {
    Intersection* intersection = findIntersection(intersectionId);
    if (intersection) {
        intersection->synchronizeSignals();
    }
}

void TrafficSystem::displaySystemStatus() const {
    std::cout << "TrafficSystem[" << systemId << "] Status: "
              << (isWorking_ ? "Working" : "Not Working") << std::endl;
    std::cout << "Intersections: " << intersectionCount << ", Signals: " << signalCount << std::endl;
    for (const Intersection* intersection : intersections) {
        if (intersection) {
            intersection->printStatus();
        }
    }
}

Intersection* TrafficSystem::findIntersection(const std::string& id) const {
    for (Intersection* intersection : intersections) {
        if (intersection && intersection->getIntersectionId() == id) {
            return intersection;
        }
    }
    return nullptr;
}

std::string TrafficSystem::generateIntersectionId() {
    std::ostringstream oss;
    oss << "I" << (intersectionCount + 1);
    return oss.str();
}

std::string TrafficSystem::generateSignalId() {
    std::ostringstream oss;
    oss << "S" << (signalCount + 1);
    return oss.str();
}
