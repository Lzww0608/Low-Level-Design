#include "post.h"
#include <iostream>
#include <algorithm>

Post::Post(const std::string userId, const std::string content, const std::string postId, PostType postType)
    : userId(userId), content(content), postId(postId), postType(postType) {
        timestamp = std::time(nullptr);
        accepted = false;
        score = 0;
    }

Post::~Post() {
    for (Comment* comment : comments) {
        delete comment;
    }
}

std::string Post::getPostId() const {
    return postId;
}

std::string Post::getUserId() const {
    return userId;
}

std::string Post::getContent() const {
    return content;
}

std::vector<std::string> Post::getTags() const {
    return tags;
}

std::vector<Comment*> Post::getComments() const {
    return comments;
}

PostType Post::getPostType() const {
    return postType;
}

std::time_t Post::getTimestamp() const {
    return timestamp;
}

bool Post::isAccepted() const {
    return accepted;
}

int Post::getScore() const {
    return score;
}


void Post::addComment(Comment* comment) {
    comments.push_back(comment);
}

bool Post::addVote(const std::string& userId) {
    if (std::find(votes.begin(), votes.end(), userId) == votes.end()) {
        votes.push_back(userId);
        score++;
        return true;
    }
    return false;
}

bool Post::removeVote(const std::string& userId) {
    auto it = std::find(votes.begin(), votes.end(), userId);
    if (it != votes.end()) {
        votes.erase(it);
        score--;
        return true;
    }
    return false;
}   

void Post::setAccepted(bool isAccepted) {
    this->accepted = isAccepted;
}

void Post::displayInfo() const {
    std::cout << "Post ID: " << postId << std::endl;    
    std::cout << "User ID: " << userId << std::endl;
    std::cout << "Content: " << content << std::endl;
    std::cout << "Timestamp: " << std::ctime(&timestamp) << std::endl;
    std::cout << "Accepted: " << (accepted ? "Yes" : "No") << std::endl;
    std::cout << "Score: " << score << std::endl;


    if (!tags.empty()) {
        std::cout << "Tags: ";
        for (const auto& tag : tags) {
            std::cout << tag << " ";
        }
        std::cout << std::endl;
    }

    if (!comments.empty()) {
        std::cout << "Comments: " << std::endl;
        for (Comment* comment : comments) {
            comment->displayInfo();
        }
    }
}
