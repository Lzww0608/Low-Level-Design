#include "FileAppender.h"

FileAppender::FileAppender(const std::string& filename)
    : filename_(filename) {
    file_.open(filename_, std::ios::app);
}

void FileAppender::append(const LogMessage& message) {
    file_ << message.getFormattedMessage() << std::endl;
}