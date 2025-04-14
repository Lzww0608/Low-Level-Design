#include "comment.h"
#include <iostream>



Comment::Comment(const std::string userId, const std::string content, const std::string commentId)
    : userId(userId), content(content), commentId(commentId) {
        timestamp = std::time(nullptr);
    }

std::string Comment::getUserId() const {
    return userId;
}

std::string Comment::getContent() const {
    return content;
}

std::string Comment::getCommentId() const {
    return commentId;
}

std::time_t Comment::getTimestamp() const {
    return timestamp;
}

void Comment::displayInfo() const {
    std::cout << "Comment ID: " << commentId << std::endl;
    std::cout << "User ID: " << userId << std::endl;
    std::cout << "Content: " << content << std::endl;
    std::cout << "Timestamp: " << std::ctime(&timestamp) << std::endl;
}
