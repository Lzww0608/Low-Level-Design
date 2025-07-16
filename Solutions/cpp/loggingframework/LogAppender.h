#ifndef LOG_APPENDER_H
#define LOG_APPENDER_H

#include "LogMessage.h"

class LogAppender {
public:
    virtual ~LogAppender() = default;
    virtual void append(const LogMessage& message) = 0;
};

#endif 