#ifndef USER_H
#define USER_H

#include <string>
#include <vector>

class User {
private:
    std::string userId;
    std::string name;
    std::string email;
    std::vector<std::string> assignedTasks; // task ids
    bool active;

public:
    User(const std::string& id, const std::string& name, const std::string& email);
    ~User();
    const std::string& getUserId() const;
    const std::string& getName() const;
    const std::string& getEmail() const;
    const std::vector<std::string>& getAssignedTasks() const;
    bool isActive() const;
    void addAssignedTask(const std::string& taskId);
    void removeAssignedTask(const std::string& taskId);
    void setActive(bool active);

    void displayUserInfo() const;
};

#endif