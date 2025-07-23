#include "User.h"
#include <iostream>
#include <algorithm>

User::User(const std::string& id, const std::string& name, const std::string& email)
    : userId(id), name(name), email(email), active(true) {}

User::~User() = default;

const std::string& User::getUserId() const {
    return userId;
}

const std::string& User::getName() const {
    return name;
}

const std::string& User::getEmail() const {
    return email;
}

const std::vector<std::string>& User::getAssignedTasks() const {
    return assignedTasks;
}

bool User::isActive() const {
    return active;
}

void User::addAssignedTask(const std::string& taskId) {
    if (std::find(assignedTasks.begin(), assignedTasks.end(), taskId) == assignedTasks.end()) {
        assignedTasks.push_back(taskId);
    }
}

void User::removeAssignedTask(const std::string& taskId) {
    auto it = std::remove(assignedTasks.begin(), assignedTasks.end(), taskId);
    if (it != assignedTasks.end()) {
        assignedTasks.erase(it, assignedTasks.end());
    }
}

void User::setActive(bool active) {
    this->active = active;
}

void User::displayUserInfo() const {
    std::cout << "User ID: " << userId << std::endl;
    std::cout << "Name: " << name << std::endl;
    std::cout << "Email: " << email << std::endl;
    std::cout << "Status: " << (active ? "Active" : "Inactive") << std::endl;
    std::cout << "Assigned Tasks: ";
    if (assignedTasks.empty()) {
        std::cout << "None";
    } else {
        for (size_t i = 0; i < assignedTasks.size(); ++i) {
            std::cout << assignedTasks[i];
            if (i != assignedTasks.size() - 1) std::cout << ", ";
        }
    }
    std::cout << std::endl;
}
