#ifndef POST_H
#define POST_H 

#include <string>
#include <vector>
#include <ctime>
#include "comment.h"


enum class PostType {
    QUESTION,
    ANSWER
};

class Post {
private:
    std::string postId;
    std::string userId;
    std::string content;
    std::vector<std::string> tags;
    std::vector<Comment*> comments;
    std::vector<std::string> votes;
    PostType postType;
    std::time_t timestamp;
    bool accepted;
    int score;

public:
    Post(const std::string userId, const std::string content, const std::string postId, PostType postType);
    ~Post();
    std::string getPostId() const;
    std::string getUserId() const;
    std::string getContent() const;
    std::vector<std::string> getTags() const;
    std::vector<Comment*> getComments() const; 
    PostType getPostType() const;
    std::time_t getTimestamp() const;
    bool isAccepted() const;
    int getScore() const;



    void addComment(Comment* comment);
    bool addVote(const std::string& userId);
    bool removeVote(const std::string& userId);
    void setAccepted(bool isAccepted);
    void displayInfo() const;
};


#endif
