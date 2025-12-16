package taskmanagementsystem

import (
	"sync"
	"time"
)

// TaskManager is the core of the task management system
// It follows the Singleton pattern to ensure a single instance
type TaskManager struct {
	tasks       map[string]*Task
	users       map[string]*User
	taskHistory map[string][]*Task // userId -> completed tasks
	mu          sync.RWMutex
}

var (
	instance *TaskManager
	once     sync.Once
)

// GetInstance returns the singleton instance of TaskManager
func GetInstance() *TaskManager {
	once.Do(func() {
		instance = &TaskManager{
			tasks:       make(map[string]*Task),
			users:       make(map[string]*User),
			taskHistory: make(map[string][]*Task),
		}
	})
	return instance
}

// CreateTask creates a new task
func (tm *TaskManager) CreateTask(task *Task) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.tasks[task.GetId()] = task
}

// UpdateTask updates an existing task
func (tm *TaskManager) UpdateTask(task *Task) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if _, exists := tm.tasks[task.GetId()]; exists {
		tm.tasks[task.GetId()] = task
		return true
	}
	return false
}

// DeleteTask deletes a task by ID
func (tm *TaskManager) DeleteTask(taskId string) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if task, exists := tm.tasks[taskId]; exists {
		// Remove task from assigned user
		if task.GetAssignedTo() != nil {
			task.GetAssignedTo().RemoveAssignedTask(task)
		}
		delete(tm.tasks, taskId)
		return true
	}
	return false
}

// GetTask retrieves a task by ID
func (tm *TaskManager) GetTask(taskId string) *Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return tm.tasks[taskId]
}

// SearchTasks searches tasks by keyword in title or description
func (tm *TaskManager) SearchTasks(keyword string) []*Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	var results []*Task
	for _, task := range tm.tasks {
		if contains(task.GetTitle(), keyword) || contains(task.GetDescription(), keyword) {
			results = append(results, task)
		}
	}
	return results
}

// FilterTasksByStatus filters tasks by status
func (tm *TaskManager) FilterTasksByStatus(status TaskStatus) []*Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	var results []*Task
	for _, task := range tm.tasks {
		if task.GetStatus() == status {
			results = append(results, task)
		}
	}
	return results
}

// FilterTasksByPriority filters tasks by priority
func (tm *TaskManager) FilterTasksByPriority(priority TaskPriority) []*Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	var results []*Task
	for _, task := range tm.tasks {
		if task.GetPriority() == priority {
			results = append(results, task)
		}
	}
	return results
}

// FilterTasksByAssignedUser filters tasks by assigned user
func (tm *TaskManager) FilterTasksByAssignedUser(userId string) []*Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	var results []*Task
	for _, task := range tm.tasks {
		if task.GetAssignedTo() != nil && task.GetAssignedTo().GetId() == userId {
			results = append(results, task)
		}
	}
	return results
}

// FilterTasksByDueDateRange filters tasks by due date range
func (tm *TaskManager) FilterTasksByDueDateRange(startDate, endDate time.Time) []*Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	var results []*Task
	for _, task := range tm.tasks {
		dueDate := task.GetDueDate()
		if (dueDate.Equal(startDate) || dueDate.After(startDate)) &&
			(dueDate.Equal(endDate) || dueDate.Before(endDate)) {
			results = append(results, task)
		}
	}
	return results
}

// MarkTaskAsCompleted marks a task as completed and moves it to history
func (tm *TaskManager) MarkTaskAsCompleted(taskId string) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[taskId]
	if !exists {
		return false
	}

	task.SetStatus(TaskStatusCompleted)

	// Add to user's task history
	if task.GetAssignedTo() != nil {
		userId := task.GetAssignedTo().GetId()
		tm.taskHistory[userId] = append(tm.taskHistory[userId], task)
	}

	return true
}

// GetTaskHistory retrieves task history for a user
func (tm *TaskManager) GetTaskHistory(userId string) []*Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	history := tm.taskHistory[userId]
	if history == nil {
		return []*Task{}
	}

	// Return a copy to avoid external modification
	result := make([]*Task, len(history))
	copy(result, history)
	return result
}

// AddUser adds a user to the system
func (tm *TaskManager) AddUser(user *User) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.users[user.GetId()] = user
}

// GetUser retrieves a user by ID
func (tm *TaskManager) GetUser(userId string) *User {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return tm.users[userId]
}

// AssignTask assigns a task to a user
func (tm *TaskManager) AssignTask(taskId, userId string) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, taskExists := tm.tasks[taskId]
	user, userExists := tm.users[userId]

	if !taskExists || !userExists {
		return false
	}

	// Remove from previous user if assigned
	if task.GetAssignedTo() != nil {
		task.GetAssignedTo().RemoveAssignedTask(task)
	}

	// Assign to new user
	user.AddAssignedTask(task)
	return true
}

// GetAllTasks returns all tasks
func (tm *TaskManager) GetAllTasks() []*Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	tasks := make([]*Task, 0, len(tm.tasks))
	for _, task := range tm.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// Helper function to check if a string contains a substring (case-insensitive)
func contains(str, substr string) bool {
	return len(str) >= len(substr) && stringContains(str, substr)
}

func stringContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
