#ifndef COMMENT_H
#define COMMENT_H

#include <string>
#include <ctime>

class Comment {
private:
    std::string userId;
    std::string content;
    std::string commentId;
    std::time_t timestamp;

public:
    Comment(const std::string userId, const std::string content, const std::string commentId);
    std::string getUserId() const;
    std::string getContent() const;
    std::string getCommentId() const;
    std::time_t getTimestamp() const;
    void displayInfo() const;
};

#endif