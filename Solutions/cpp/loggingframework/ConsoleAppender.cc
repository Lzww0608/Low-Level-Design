#include "ConsoleAppender.h"
#include <iostream>

void ConsoleAppender::append(const LogMessage& message) {
    std::cout << message.getFormattedMessage() << std::endl;
}