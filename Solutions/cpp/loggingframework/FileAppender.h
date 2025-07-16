#ifndef FILE_APPENDER_H
#define FILE_APPENDER_H

#include "LogAppender.h"
#include <fstream>

class FileAppender : public LogAppender {
public:
    FileAppender(const std::string& filename);
    void append(const LogMessage& message) override;
private:
    std::string filename_;
    std::ofstream file_;
};

#endif 