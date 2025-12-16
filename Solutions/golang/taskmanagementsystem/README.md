# Task Management System

A comprehensive task management system implemented in Go, demonstrating design patterns and concurrent programming.

## Features

- **Singleton Pattern**: Ensures a single instance of TaskManager
- **Thread-Safe Operations**: Uses mutex locks for concurrent access
- **User Management**: Create and manage users
- **Task Management**: Create, update, delete, and track tasks
- **Task Assignment**: Assign tasks to users
- **Task Filtering**: Filter tasks by status, priority, assigned user, and due date
- **Task Search**: Search tasks by keyword
- **Task History**: Track completed tasks per user
- **Comments**: Add comments to tasks

## Architecture

### Core Components

1. **User** (`user.go`): Represents a user with personal information and assigned tasks
2. **Task** (`task.go`): Represents a task with various properties including status, priority, and due date
3. **TaskManager** (`task_manager.go`): Singleton manager that coordinates all operations
4. **TaskStatus**: Enum for task states (Pending, InProgress, Completed, Blocked)
5. **TaskPriority**: Enum for task priorities (Low, Medium, High, Urgent)

### Design Patterns

- **Singleton Pattern**: TaskManager uses `sync.Once` to ensure single instance
- **Thread-Safe Design**: All public methods use `sync.RWMutex` for concurrent access
- **Encapsulation**: Private fields with public getters/setters

## Project Structure

```
taskmanagementsystem/
├── task.go                  # Task entity and enums
├── user.go                  # User entity
├── task_manager.go          # Core manager with singleton pattern
├── task_manager_test.go     # Comprehensive test suite
├── go.mod                   # Go module definition
├── README.md               # This file
└── examples/
    ├── main.go             # Demo application
    └── go.mod              # Example module definition
```

## Installation

```bash
cd /home/lab2439/Work/lzww/LLD/Solutions/golang/taskmanagementsystem
go mod tidy
```

## Running Tests

Run all tests:
```bash
go test -v
```

Run tests with coverage:
```bash
go test -v -cover
```

Run specific test:
```bash
go test -v -run TestSingletonPattern
```

## Running the Demo

```bash
cd examples
go run main.go
```

## Usage Examples

### Basic Usage

```go
package main

import (
    "time"
    tms "taskmanagementsystem"
)

func main() {
    // Get singleton instance
    taskManager := tms.GetInstance()
    
    // Create a user
    user := tms.NewUser("U001", "John Doe", "john@example.com")
    taskManager.AddUser(user)
    
    // Create a task
    dueDate := time.Now().Add(24 * time.Hour)
    task := tms.NewTask(
        "T001",
        "Complete Project",
        "Finish the task management system",
        dueDate,
        tms.TaskPriorityHigh,
        nil,
    )
    taskManager.CreateTask(task)
    
    // Assign task to user
    taskManager.AssignTask("T001", "U001")
    
    // Update task status
    task.SetStatus(tms.TaskStatusInProgress)
    taskManager.UpdateTask(task)
    
    // Mark as completed
    taskManager.MarkTaskAsCompleted("T001")
    
    // View user's task history
    history := taskManager.GetTaskHistory("U001")
}
```

### Searching and Filtering

```go
// Search by keyword
results := taskManager.SearchTasks("authentication")

// Filter by status
pendingTasks := taskManager.FilterTasksByStatus(tms.TaskStatusPending)

// Filter by priority
urgentTasks := taskManager.FilterTasksByPriority(tms.TaskPriorityUrgent)

// Filter by assigned user
userTasks := taskManager.FilterTasksByAssignedUser("U001")

// Filter by due date range
startDate := time.Now()
endDate := startDate.Add(7 * 24 * time.Hour)
tasksThisWeek := taskManager.FilterTasksByDueDateRange(startDate, endDate)
```

### Task Comments

```go
task := taskManager.GetTask("T001")
task.AddComment("Started working on this task")
task.AddComment("Encountered issue with API")
comments := task.GetComments()
```

## Thread Safety

The TaskManager is designed to handle concurrent operations safely:

- All public methods use appropriate locking mechanisms
- Read operations use `RLock()` for better performance
- Write operations use `Lock()` for exclusive access
- The singleton pattern is implemented with `sync.Once`

## Test Coverage

The test suite includes:

- Singleton pattern verification
- CRUD operations for tasks and users
- Task assignment and reassignment
- Search and filter functionality
- Task completion and history tracking
- Concurrent operations testing
- Edge cases and error handling

All tests pass successfully with comprehensive coverage.

## Requirements Met

✅ Create, update, and delete tasks  
✅ Task properties (title, description, due date, priority, status)  
✅ Assign tasks to users and set reminders  
✅ Search and filter tasks by various criteria  
✅ Mark tasks as completed and view history  
✅ Handle concurrent access with data consistency  
✅ Extensible design for future enhancements  

## Future Enhancements

Potential improvements:
- Database persistence
- REST API endpoints
- Task dependencies and subtasks
- Notifications and reminders
- Task templates
- Bulk operations
- Export/import functionality
- User roles and permissions

## License

This is a learning project for Low-Level Design (LLD) practice.

