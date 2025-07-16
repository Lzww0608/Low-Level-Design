#ifndef LOG_MESSAGE_H
#define LOG_MESSAGE_H

#include "LogLevel.h"
#include <string>

class LogMessage {
public:
    LogMessage(LogLevel level, const std::string& message,
        const std::string& source);
    LogLevel getLevel() const;
    std::string getMessage() const;
    std::string getSource() const;
    std::string getTimestamp() const;
    std::string getFormattedMessage() const;
private:
    LogLevel level_;
    std::string message_;
    std::string timestamp_;
    std::string source_;
};

#endif 