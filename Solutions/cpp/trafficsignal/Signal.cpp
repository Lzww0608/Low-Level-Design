#include "Singal.h"

Signal::Signal(const std::string& id, int green, int yellow, int red)
    : signalId(id), greenDuration(green), yellowDuration(yellow), redDuration(red), currentSignal(SignalType::RED), isWorking_(true) {}

Signal::~Signal() {}

std::string Signal::getSignalId() const {
    return signalId;
}

SignalType Signal::getCurrentSignal() const {
    return currentSignal;
}

int Signal::getRemainingDuration() const {
    return remaingDuration;
}

bool Signal::isWorking() const {
    return isWorking_;
}

void Signal::setSignal(SignalType signal) {
    currentSignal = signal;
    switch (signal) {
        case SignalType::GREEN:
            currentSignal = SignalType::GREEN;
            remaingDuration = greenDuration;
            break;
        case SignalType::YELLOW:
            currentSignal = SignalType::YELLOW;
            remaingDuration = yellowDuration;
            break;
        case SignalType::RED:
            currentSignal = SignalType::RED;
            remaingDuration = redDuration;
            break;
        default:
            remaingDuration = 0;
            break;
    }
}

void Signal::updateSignal(int timeElapsed) {
    if (!isWorking_) return;
    if (remaingDuration > timeElapsed) {
        remaingDuration -= timeElapsed;
    } else {
        remaingDuration = 0;
        switchSignal();
    }
}

void Signal::setDuration(int green, int yellow, int red) {
    greenDuration = green;
    yellowDuration = yellow;
    redDuration = red;
    // 若当前信号为某色，重置剩余时间
    switch (currentSignal) {
        case SignalType::GREEN:
            remaingDuration = greenDuration;
            break;
        case SignalType::YELLOW:
            remaingDuration = yellowDuration;
            break;
        case SignalType::RED:
            remaingDuration = redDuration;
            break;
        default:
            remaingDuration = 0;
            break;
    }
}

void Signal::setWorking(bool working) {
    isWorking_ = working;
}

void Signal::reset() {
    currentSignal = SignalType::RED;
    remaingDuration = redDuration;
    isWorking_ = true;
}

void Signal::switchSignal() {
    if (!isWorking_) return;
    switch (currentSignal) {
        case SignalType::RED:
            currentSignal = SignalType::GREEN;
            remaingDuration = greenDuration;
            break;
        case SignalType::GREEN:
            currentSignal = SignalType::YELLOW;
            remaingDuration = yellowDuration;
            break;
        case SignalType::YELLOW:
            currentSignal = SignalType::RED;
            remaingDuration = redDuration;
            break;
        default:
            currentSignal = SignalType::RED;
            remaingDuration = redDuration;
            break;
    }
}

void Signal::printStatus() const {
    std::string signalStr;
    switch (currentSignal) {
        case SignalType::RED:
            signalStr = "RED";
            break;
        case SignalType::GREEN:
            signalStr = "GREEN";
            break;
        case SignalType::YELLOW:
            signalStr = "YELLOW";
            break;
        default:
            signalStr = "UNKNOWN";
            break;
    }
    std::cout << "Signal[" << signalId << "] Status: " << signalStr
              << ", Remaining: " << remaingDuration
              << ", Working: " << (isWorking_ ? "Yes" : "No") << std::endl;
}
