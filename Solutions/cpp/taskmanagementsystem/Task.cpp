#include "Task.h"
#include <algorithm>
#include <iostream>

Task::Task(const std::string& id, const std::string& title, const std::string& description, TaskStatus status, TaskPriority priority, const std::time_t& dueDate, const std::string& assignedTo)
    : taskId(id),
      title(title),
      description(description),
      status(status),
      priority(priority),
      dueDate(dueDate),
      assignedTo(assignedTo)
{
}

Task::~Task() = default;

const std::string& Task::getTaskId() const {
    return taskId;
}

const std::string& Task::getTitle() const {
    return title;
}

const std::string& Task::getDescription() const {
    return description;
}

TaskStatus Task::getStatus() const {
    return status;
}

TaskPriority Task::getPriority() const {
    return priority;
}

const std::time_t& Task::getDueDate() const {
    return dueDate;
}

const std::string& Task::getAssignedTo() const {
    return assignedTo;
}

const std::vector<std::string>& Task::getDependencies() const {
    return dependencies;
}

const std::vector<std::string>& Task::getComments() const {
    return comments;
}

void Task::setTitle(const std::string& title) {
    this->title = title;
}

void Task::setDescription(const std::string& description) {
    this->description = description;
}

void Task::setStatus(TaskStatus status) {
    this->status = status;
}

void Task::setPriority(TaskPriority priority) {
    this->priority = priority;
}

void Task::setDueDate(const std::time_t& dueDate) {
    this->dueDate = dueDate;
}

void Task::setAssignedTo(const std::string& assignedTo) {
    this->assignedTo = assignedTo;
}

void Task::addDependency(const std::string& dependency) {
    if (std::find(dependencies.begin(), dependencies.end(), dependency) == dependencies.end()) {
        dependencies.push_back(dependency);
    }
}

void Task::removeDependency(const std::string& dependency) {
    auto it = std::remove(dependencies.begin(), dependencies.end(), dependency);
    if (it != dependencies.end()) {
        dependencies.erase(it, dependencies.end());
    }
}

void Task::addComment(const std::string& comment) {
    comments.push_back(comment);
}

void Task::removeComment(const std::string& comment) {
    auto it = std::remove(comments.begin(), comments.end(), comment);
    if (it != comments.end()) {
        comments.erase(it, comments.end());
    }
}


void Task::displayTaskInfo() const {
    std::cout << "Task ID: " << taskId << std::endl;
    std::cout << "Title: " << title << std::endl;
    std::cout << "Description: " << description << std::endl;

    std::cout << "Status: ";
    switch (status) {
        case TaskStatus::TODO: std::cout << "TODO"; break;
        case TaskStatus::IN_PROGRESS: std::cout << "IN_PROGRESS"; break;
        case TaskStatus::COMPLETED: std::cout << "COMPLETED"; break;
        case TaskStatus::BLOCKED: std::cout << "BLOCKED"; break;
        default: std::cout << "UNKNOWN"; break;
    }
    std::cout << std::endl;

    std::cout << "Priority: ";
    switch (priority) {
        case TaskPriority::LOW: std::cout << "LOW"; break;
        case TaskPriority::MEDIUM: std::cout << "MEDIUM"; break;
        case TaskPriority::HIGH: std::cout << "HIGH"; break;
        case TaskPriority::URGENT: std::cout << "URGENT"; break;
        default: std::cout << "UNKNOWN"; break;
    }
    std::cout << std::endl;

    char buf[64];
    std::tm* tm_info = std::localtime(&dueDate);
    if (tm_info && std::strftime(buf, sizeof(buf), "%Y-%m-%d %H:%M:%S", tm_info)) {
        std::cout << "Due Date: " << buf << std::endl;
    } else {
        std::cout << "Due Date: (invalid)" << std::endl;
    }

    std::cout << "Assigned To: " << assignedTo << std::endl;

    std::cout << "Dependencies: ";
    if (dependencies.empty()) {
        std::cout << "None";
    } else {
        for (size_t i = 0; i < dependencies.size(); ++i) {
            std::cout << dependencies[i];
            if (i != dependencies.size() - 1) std::cout << ", ";
        }
    }
    std::cout << std::endl;

    std::cout << "Comments: ";
    if (comments.empty()) {
        std::cout << "None";
    } else {
        for (size_t i = 0; i < comments.size(); ++i) {
            std::cout << "[" << (i+1) << "] " << comments[i];
            if (i != comments.size() - 1) std::cout << "; ";
        }
    }
    std::cout << std::endl;
}
