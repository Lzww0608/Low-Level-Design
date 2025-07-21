#ifndef TRAFFIC_SIGNAL_H
#define TRAFFIC_SIGNAL_H

#include <string>
#include <iostream>

enum class SignalType {
    RED,
    GREEN,
    YELLOW,
    UNKNOWN
};

class Signal {
public:
    Signal(const std::string& id, int green, int yellow, int red);
    ~Signal();
    
    std::string getSignalId() const;
    SignalType getCurrentSignal() const;
    int getRemainingDuration() const;
    bool isWorking() const;

    void setSignal(SignalType signal);
    void updateSignal(int timeElapsed);
    void setDuration(int green, int yellow, int red);
    void setWorking(bool working);
    void reset();
    void switchSignal();
    void printStatus() const;

private:
    std::string signalId;
    SignalType currentSignal;
    int remaingDuration;
    int greenDuration;
    int yellowDuration;
    int redDuration;

    bool isWorking_;
};



#endif