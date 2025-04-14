#include "user.h"
#include <iostream>



User::User(std::string userId, std::string username, std::string email)
    : userId(userId), username(username), email(email), reputation(1), active(true) {}


std::string User::getUserId() const {
    return userId;
}

std::string User::getUsername() const {
    return username;
}

std::string User::getEmail() const {
    return email;
}


int User::getReputation() const {
    return reputation;
}



std::vector<std::string> User::getBadges() const {
    return badges;
}

bool User::isActive() const {
    return active;
}

void User::updateReputation(int delta) {
    reputation += delta;
}

void User::addBadge(const std::string& badge) {
    badges.push_back(badge);
}

void User::setActive(bool isActive) {
    active = isActive;
}


void User::displayInfo() const {
    std::cout << "User ID: " << userId << std::endl;
    std::cout << "Username: " << username << std::endl;
    std::cout << "Email: " << email << std::endl;
    std::cout << "Reputation: " << reputation << std::endl;
    std::cout << "Active: " << (active ? "Yes" : "No") << std::endl;

    if (!badges.empty()) {
        std::cout << "Badges: ";
        for (const auto& badge : badges) {
            std::cout << badge << " ";
        }
        std::cout << std::endl;
    }
}


