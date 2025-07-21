#include "Intersection.h"
#include <iostream>
#include <algorithm>

Intersection::Intersection(const std::string& id)
    : intersectionId(id), isWorking_(true) {}

Intersection::~Intersection() {
    // 释放信号灯对象
    for (Signal* signal : signals) {
        delete signal;
    }
    signals.clear();
}

std::string Intersection::getIntersectionId() const {
    return intersectionId;
}

bool Intersection::isWorking() const {
    return isWorking_;
}

void Intersection::setWorking(bool working) {
    isWorking_ = working;
    // 同步所有信号灯的工作状态
    for (Signal* signal : signals) {
        if (signal) {
            signal->setWorking(working);
        }
    }
}

void Intersection::addSignal(Signal* signal) {
    if (!signal) return;
    // 避免重复添加
    if (findSignal(signal->getSignalId()) == nullptr) {
        signals.push_back(signal);
    }
}

void Intersection::removeSignal(const std::string& id) {
    auto it = std::remove_if(signals.begin(), signals.end(),
        [&](Signal* s) {
            if (s && s->getSignalId() == id) {
                delete s;
                return true;
            }
            return false;
        });
    signals.erase(it, signals.end());
}

void Intersection::updateSignal(const std::string& id, int timeElapsed) {
    Signal* signal = findSignal(id);
    if (signal) {
        signal->updateSignal(timeElapsed);
    }
}

void Intersection::synchronizeSignals() {
    // 使所有信号灯切换到同一状态（以第一个信号灯为基准）
    if (signals.empty()) return;
    SignalType baseType = signals[0]->getCurrentSignal();
    for (Signal* signal : signals) {
        if (signal) {
            signal->setSignal(baseType);
        }
    }
}

void Intersection::printStatus() const {
    std::cout << "Intersection[" << intersectionId << "] Status: "
              << (isWorking_ ? "Working" : "Not Working") << std::endl;
    for (const Signal* signal : signals) {
        if (signal) {
            signal->printStatus();
        }
    }
}

Signal* Intersection::findSignal(const std::string& id) {
    for (Signal* signal : signals) {
        if (signal && signal->getSignalId() == id) {
            return signal;
        }
    }
    return nullptr;
}
