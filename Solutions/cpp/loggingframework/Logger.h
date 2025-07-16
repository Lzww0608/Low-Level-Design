#ifndef LOGGER_H
#define LOGGER_H

#include "LogAppender.h"
#include "LogMessage.h"
#include <mutex>
#include <memory>
#include <vector>

class Logger {
private:
    std::string name_;
    LogLevel min_level_;
    std::vector<std::shared_ptr<LogAppender>> appenders_;

public:
    Logger(const std::string& name, LogLevel min_level = LogLevel::INFO);
    void addAppender(const std::shared_ptr<LogAppender>& appender);
    void log(LogLevel level, const std::string& message, const std::string& source);
    void debug(const std::string& message, const std::string& source);
    void info(const std::string& message, const std::string& source);
    void warning(const std::string& message, const std::string& source);
    void error(const std::string& message, const std::string& source);
    void fatal(const std::string& message, const std::string& source);

private:
    bool isLevelEnabled(LogLevel level) const;
};

#endif 