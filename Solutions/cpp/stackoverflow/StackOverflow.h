#ifndef STACKOVERFLOW_H
#define STACKOVERFLOW_H

#include <string>
#include <vector>

#include "user.h"
#include "post.h"
#include "comment.h"

class StackOverflow {
private:
    std::vector<User*> users;
    std::vector<Post*> posts;
    int userIdCounter;
    int postIdCounter;
    int commentIdCounter;

public:
    StackOverflow();
    ~StackOverflow();
    
    // User Management
    User* registerUser(const std::string&, const std::string&);
    void removeUser(const std::string&);

    // Post Management
    Post* addQuestion(const std::string&, const std::string&, const std::vector<std::string>&);
    Post* addAnswer(const std::string&, const std::string&, const std::string&);
    Comment* addComment(const std::string&, const std::string&, const std::string&);

    // Voting and Accepting Answers
    bool votePost(const std::string&, const std::string&);
    bool unvotePost(const std::string&, const std::string&);
    bool acceptAnswer(const std::string&, const std::string&);

    // Search and Display
    std::vector<Post*> searchQuestions(const std::string&) const;
    void displayUserProfile(const std::string&) const;
    void displayQuestion(const std::string&) const;
    void displayAllQuestions() const;
    
private:
    User* findUser(const std::string&) const;
    Post* findPost(const std::string&) const;
    void updateUserReputation(const std::string&, int);
    std::string generateUserId();
    std::string generatePostId();
    std::string generateCommentId();
};

#endif