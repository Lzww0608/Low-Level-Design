#include "LogMessage.h"
#include <chrono>
#include <sstream>
#include <iomanip>
#include <ctime>

LogMessage::LogMessage(LogLevel level, const std::string& message,
    const std::string& source)
    : level_(level), message_(message), source_(source) {
    // 获取当前时间并格式化
    auto now = std::chrono::system_clock::now();
    auto time_t = std::chrono::system_clock::to_time_t(now);
    auto ms = std::chrono::duration_cast<std::chrono::milliseconds>(
        now.time_since_epoch()) % 1000;
    
    std::stringstream ss;
    ss << std::put_time(std::localtime(&time_t), "%Y-%m-%d %H:%M:%S");
    ss << '.' << std::setfill('0') << std::setw(3) << ms.count();
    timestamp_ = ss.str();
}

LogLevel LogMessage::getLevel() const {
    return level_;
}

std::string LogMessage::getMessage() const {
    return message_;
}

std::string LogMessage::getSource() const {
    return source_;
}

std::string LogMessage::getTimestamp() const {
    return timestamp_;
}

std::string LogMessage::getFormattedMessage() const {
    std::ostringstream oss;
    oss << timestamp_ << " [" << logLevelToString(level_) << "] " << source_ << ": " << message_;
    return oss.str();
}