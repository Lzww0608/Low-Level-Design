#ifndef CONSOLE_APPENDER_H
#define CONSOLE_APPENDER_H

#include "LogAppender.h"

class ConsoleAppender : public LogAppender {
public:
    void append(const LogMessage& message) override;
};

#endif 