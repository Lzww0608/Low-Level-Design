#include "Logger.h"

Logger::Logger(const std::string& name, LogLevel min_level)
    : name_(name), min_level_(min_level) {}

void Logger::addAppender(const std::shared_ptr<LogAppender>& appender) {
    appenders_.push_back(appender);
}

void Logger::log(LogLevel level, const std::string& message, const std::string& source) {
    if (isLevelEnabled(level)) {
        LogMessage log_message(level, message, source);
        for (auto& appender : appenders_) {
            appender->append(log_message);
        }
    }
}

void Logger::debug(const std::string& message, const std::string& source) {
    log(LogLevel::DEBUG, message, source);
}

void Logger::info(const std::string& message, const std::string& source) {
    log(LogLevel::INFO, message, source);
}

void Logger::warning(const std::string& message, const std::string& source) {
    log(LogLevel::WARNING, message, source);
}

void Logger::error(const std::string& message, const std::string& source) {
    log(LogLevel::ERROR, message, source);
}

void Logger::fatal(const std::string& message, const std::string& source) {
    log(LogLevel::FATAL, message, source);
}

bool Logger::isLevelEnabled(LogLevel level) const {
    return static_cast<int>(level) >= static_cast<int>(min_level_);
}

