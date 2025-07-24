#include "TaskManager.h"
#include <iostream>
#include <algorithm>

// 构造函数
TaskManager::TaskManager() : taskIdCounter(1), userIdCounter(1) {}

// 析构函数
TaskManager::~TaskManager() {
    for (auto t : tasks) delete t;
    for (auto u : users) delete u;
}

// 用户注册
User* TaskManager::registerUser(const std::string& username, const std::string& email) {
    std::string id = generateUserId();
    User* user = new User(id, username, email);
    users.push_back(user);
    return user;
}

// 移除用户
void TaskManager::removeUser(const std::string& userId) {
    auto it = std::find_if(users.begin(), users.end(),
        [&](User* u) { return u->getUserId() == userId; });
    if (it != users.end()) {
        // 解除任务分配
        for (const std::string& tid : (*it)->getAssignedTasks()) {
            Task* t = findTask(tid);
            if (t) t->setAssignedTo("");
        }
        delete *it;
        users.erase(it);
    }
}

// 创建任务
Task* TaskManager::createTask(const std::string& title, const std::string& description, TaskPriority priority) {
    std::string id = generateTaskId();
    Task* task = new Task(id, title, description, TaskStatus::TODO, priority, 0, "");
    tasks.push_back(task);
    return task;
}

// 移除任务
void TaskManager::removeTask(const std::string& taskId) {
    auto it = std::find_if(tasks.begin(), tasks.end(),
        [&](Task* t) { return t->getTaskId() == taskId; });
    if (it != tasks.end()) {
        // 解除用户分配
        std::string assignee = (*it)->getAssignedTo();
        if (!assignee.empty()) {
            User* u = findUser(assignee);
            if (u) u->removeAssignedTask(taskId);
        }
        delete *it;
        tasks.erase(it);
    }
}

// 分配任务
bool TaskManager::assignTask(const std::string& taskId, const std::string& userId) {
    Task* task = findTask(taskId);
    User* user = findUser(userId);
    if (!task || !user || !user->isActive()) return false;
    // 解除原分配
    std::string prevAssignee = task->getAssignedTo();
    if (!prevAssignee.empty() && prevAssignee != userId) {
        User* prevUser = findUser(prevAssignee);
        if (prevUser) prevUser->removeAssignedTask(taskId);
    }
    task->setAssignedTo(userId);
    user->addAssignedTask(taskId);
    return true;
}

// 更新任务状态
bool TaskManager::updateTaskStatus(const std::string& taskId, TaskStatus status) {
    Task* task = findTask(taskId);
    if (!task) return false;
    // 检查依赖
    if (status == TaskStatus::IN_PROGRESS || status == TaskStatus::COMPLETED) {
        if (!checkDependenciesMet(task)) return false;
    }
    task->setStatus(status);
    return true;
}

// 添加任务依赖
bool TaskManager::addTaskDependency(const std::string& taskId, const std::string& dependencyId) {
    Task* task = findTask(taskId);
    Task* dep = findTask(dependencyId);
    if (!task || !dep || taskId == dependencyId) return false;
    task->addDependency(dependencyId);
    return true;
}

// 添加任务评论
bool TaskManager::addTaskComment(const std::string& taskId, const std::string& comment) {
    Task* task = findTask(taskId);
    if (!task) return false;
    task->addComment(comment);
    return true;
}

// 显示用户任务
void TaskManager::displayUserTasks(const std::string& userId) const {
    const User* user = findUser(userId);
    if (!user) {
        std::cout << "User not found.\n";
        return;
    }
    std::cout << "Tasks for user " << user->getName() << ":\n";
    for (const std::string& tid : user->getAssignedTasks()) {
        const Task* t = findTask(tid);
        if (t) t->displayTaskInfo();
    }
}

// 显示所有任务
void TaskManager::displayAllTasks() const {
    std::cout << "All Tasks:\n";
    for (const Task* t : tasks) {
        t->displayTaskInfo();
    }
}

// 显示所有用户
void TaskManager::displayAllUsers() const {
    std::cout << "All Users:\n";
    for (const User* u : users) {
        u->displayUserInfo();
    }
}

// 按状态获取任务
std::vector<Task*> TaskManager::getTasksByStatus(TaskStatus status) const {
    std::vector<Task*> result;
    for (Task* t : tasks) {
        if (t->getStatus() == status) result.push_back(t);
    }
    return result;
}

// 按优先级获取任务
std::vector<Task*> TaskManager::getTasksByPriority(TaskPriority priority) const {
    std::vector<Task*> result;
    for (Task* t : tasks) {
        if (t->getPriority() == priority) result.push_back(t);
    }
    return result;
}

// 私有：查找用户
User* TaskManager::findUser(const std::string& userId) const {
    auto it = std::find_if(users.begin(), users.end(),
        [&](User* u) { return u->getUserId() == userId; });
    return it != users.end() ? *it : nullptr;
}

// 私有：查找任务
Task* TaskManager::findTask(const std::string& taskId) const {
    auto it = std::find_if(tasks.begin(), tasks.end(),
        [&](Task* t) { return t->getTaskId() == taskId; });
    return it != tasks.end() ? *it : nullptr;
}

// 私有：检查依赖是否满足
bool TaskManager::checkDependenciesMet(const Task* task) const {
    for (const std::string& depId : task->getDependencies()) {
        const Task* dep = findTask(depId);
        if (!dep || dep->getStatus() != TaskStatus::COMPLETED) return false;
    }
    return true;
}

// 私有：生成任务ID
std::string TaskManager::generateTaskId() {
    return "T" + std::to_string(taskIdCounter++);
}

// 私有：生成用户ID
std::string TaskManager::generateUserId() {
    return "U" + std::to_string(userIdCounter++);
}
