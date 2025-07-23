#ifndef TASK_H
#define TASK_H

#include <string>
#include <ctime>
#include <vector>
#include <memory>

enum class TaskStatus {
    TODO,
    IN_PROGRESS,
    COMPLETED,
    BLOCKED
};

enum class TaskPriority {
    LOW,
    MEDIUM,
    HIGH,
    URGENT
};

class Task {
private:
    std::string taskId;
    std::string title;
    std::string description;
    TaskStatus status;
    TaskPriority priority;
    std::time_t dueDate;
    std::string assignedTo;
    std::vector<std::string> dependencies;
    std::vector<std::string> comments;

public:
    Task(const std::string& id, const std::string& title, const std::string& description, TaskStatus status, TaskPriority priority, const std::time_t& dueDate, const std::string& assignedTo);
    ~Task();

    const std::string& getTaskId() const;
    const std::string& getTitle() const;
    const std::string& getDescription() const;
    TaskStatus getStatus() const;
    TaskPriority getPriority() const;
    const std::time_t& getDueDate() const;
    const std::string& getAssignedTo() const;
    const std::vector<std::string>& getDependencies() const;
    const std::vector<std::string>& getComments() const;

    void setTitle(const std::string& title);
    void setDescription(const std::string& description);
    void setStatus(TaskStatus status);
    void setPriority(TaskPriority priority);
    void setDueDate(const std::time_t& dueDate);
    void setAssignedTo(const std::string& assignedTo);
    void addDependency(const std::string& dependency);
    void removeDependency(const std::string& dependency);
    void addComment(const std::string& comment);
    void removeComment(const std::string& comment);

    void displayTaskInfo() const;
};

#endif