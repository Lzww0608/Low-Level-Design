package main

import (
	"fmt"
	"time"

	tms "taskmanagementsystem"
)

func main() {
	fmt.Println("=== Task Management System Demo ===\n")

	// Get the singleton instance of TaskManager
	taskManager := tms.GetInstance()

	// Create users
	fmt.Println("Creating users...")
	user1 := tms.NewUser("U001", "Alice Johnson", "alice@example.com")
	user2 := tms.NewUser("U002", "Bob Smith", "bob@example.com")
	user3 := tms.NewUser("U003", "Charlie Brown", "charlie@example.com")

	taskManager.AddUser(user1)
	taskManager.AddUser(user2)
	taskManager.AddUser(user3)
	fmt.Println("Users created successfully\n")

	// Create tasks
	fmt.Println("Creating tasks...")

	task1 := tms.NewTask(
		"T001",
		"Design System Architecture",
		"Create high-level design for the new microservices architecture",
		time.Now().Add(3*24*time.Hour),
		tms.TaskPriorityHigh,
		nil,
	)

	task2 := tms.NewTask(
		"T002",
		"Implement User Authentication",
		"Develop JWT-based authentication system",
		time.Now().Add(5*24*time.Hour),
		tms.TaskPriorityUrgent,
		nil,
	)

	task3 := tms.NewTask(
		"T003",
		"Write Unit Tests",
		"Create unit tests for the task management module",
		time.Now().Add(7*24*time.Hour),
		tms.TaskPriorityMedium,
		nil,
	)

	task4 := tms.NewTask(
		"T004",
		"Update Documentation",
		"Update API documentation with new endpoints",
		time.Now().Add(10*24*time.Hour),
		tms.TaskPriorityLow,
		nil,
	)

	taskManager.CreateTask(task1)
	taskManager.CreateTask(task2)
	taskManager.CreateTask(task3)
	taskManager.CreateTask(task4)
	fmt.Println("Tasks created successfully\n")

	// Assign tasks to users
	fmt.Println("Assigning tasks to users...")
	taskManager.AssignTask("T001", "U001")
	taskManager.AssignTask("T002", "U001")
	taskManager.AssignTask("T003", "U002")
	taskManager.AssignTask("T004", "U003")
	fmt.Println("Tasks assigned successfully\n")

	// Display user information
	fmt.Println("=== User Information ===")
	user1.DisplayUserInfo()
	fmt.Println()

	// Update task status
	fmt.Println("Updating task statuses...")
	task1.SetStatus(tms.TaskStatusInProgress)
	taskManager.UpdateTask(task1)
	task3.SetStatus(tms.TaskStatusInProgress)
	taskManager.UpdateTask(task3)
	fmt.Println("Task statuses updated\n")

	// Add comments to tasks
	fmt.Println("Adding comments to tasks...")
	task1.AddComment("Started working on the architecture diagram")
	task1.AddComment("Reviewed with senior architect")
	task2.AddComment("Setting up OAuth2 integration")
	fmt.Println("Comments added\n")

	// Search for tasks
	fmt.Println("=== Searching for tasks ===")
	searchResults := taskManager.SearchTasks("Authentication")
	fmt.Printf("Found %d task(s) containing 'Authentication':\n", len(searchResults))
	for _, task := range searchResults {
		fmt.Printf("  - %s: %s\n", task.GetId(), task.GetTitle())
	}
	fmt.Println()

	// Filter tasks by priority
	fmt.Println("=== Filtering by Priority ===")
	highPriorityTasks := taskManager.FilterTasksByPriority(tms.TaskPriorityHigh)
	fmt.Printf("High priority tasks (%d):\n", len(highPriorityTasks))
	for _, task := range highPriorityTasks {
		fmt.Printf("  - %s: %s (Priority: High)\n", task.GetId(), task.GetTitle())
	}
	fmt.Println()

	urgentTasks := taskManager.FilterTasksByPriority(tms.TaskPriorityUrgent)
	fmt.Printf("Urgent priority tasks (%d):\n", len(urgentTasks))
	for _, task := range urgentTasks {
		fmt.Printf("  - %s: %s (Priority: Urgent)\n", task.GetId(), task.GetTitle())
	}
	fmt.Println()

	// Filter tasks by status
	fmt.Println("=== Filtering by Status ===")
	inProgressTasks := taskManager.FilterTasksByStatus(tms.TaskStatusInProgress)
	fmt.Printf("In-progress tasks (%d):\n", len(inProgressTasks))
	for _, task := range inProgressTasks {
		fmt.Printf("  - %s: %s\n", task.GetId(), task.GetTitle())
	}
	fmt.Println()

	// Filter tasks by assigned user
	fmt.Println("=== Filtering by Assigned User ===")
	aliceTasks := taskManager.FilterTasksByAssignedUser("U001")
	fmt.Printf("Tasks assigned to Alice (%d):\n", len(aliceTasks))
	for _, task := range aliceTasks {
		fmt.Printf("  - %s: %s\n", task.GetId(), task.GetTitle())
	}
	fmt.Println()

	// Filter tasks by due date range
	fmt.Println("=== Filtering by Due Date Range ===")
	startDate := time.Now()
	endDate := time.Now().Add(6 * 24 * time.Hour)
	tasksInRange := taskManager.FilterTasksByDueDateRange(startDate, endDate)
	fmt.Printf("Tasks due within next 6 days (%d):\n", len(tasksInRange))
	for _, task := range tasksInRange {
		fmt.Printf("  - %s: %s (Due: %s)\n",
			task.GetId(),
			task.GetTitle(),
			task.GetDueDate().Format("2006-01-02"))
	}
	fmt.Println()

	// Mark task as completed
	fmt.Println("=== Completing Tasks ===")
	fmt.Println("Marking task T003 as completed...")
	taskManager.MarkTaskAsCompleted("T003")

	completedTask := taskManager.GetTask("T003")
	fmt.Printf("Task %s status: ", completedTask.GetId())
	switch completedTask.GetStatus() {
	case tms.TaskStatusCompleted:
		fmt.Println("Completed")
	case tms.TaskStatusInProgress:
		fmt.Println("In Progress")
	case tms.TaskStatusPending:
		fmt.Println("Pending")
	default:
		fmt.Println("Unknown")
	}
	fmt.Println()

	// View task history for a user
	fmt.Println("=== Task History ===")
	bobHistory := taskManager.GetTaskHistory("U002")
	fmt.Printf("Bob's completed tasks (%d):\n", len(bobHistory))
	for _, task := range bobHistory {
		fmt.Printf("  - %s: %s\n", task.GetId(), task.GetTitle())
	}
	fmt.Println()

	// Reassign a task
	fmt.Println("=== Task Reassignment ===")
	fmt.Println("Reassigning task T004 from Charlie to Bob...")
	taskManager.AssignTask("T004", "U002")

	updatedTask := taskManager.GetTask("T004")
	if updatedTask.GetAssignedTo() != nil {
		fmt.Printf("Task T004 is now assigned to: %s\n", updatedTask.GetAssignedTo().GetName())
	}
	fmt.Println()

	// Update task details
	fmt.Println("=== Updating Task Details ===")
	task4.SetTitle("Update API and User Documentation")
	task4.SetDescription("Update both API documentation and user guide with new features")
	task4.SetPriority(tms.TaskPriorityMedium)
	taskManager.UpdateTask(task4)
	fmt.Printf("Task %s updated: %s\n", task4.GetId(), task4.GetTitle())
	fmt.Println()

	// Display all tasks summary
	fmt.Println("=== All Tasks Summary ===")
	allTasks := taskManager.GetAllTasks()
	fmt.Printf("Total tasks in system: %d\n\n", len(allTasks))

	for _, task := range allTasks {
		fmt.Printf("Task ID: %s\n", task.GetId())
		fmt.Printf("  Title: %s\n", task.GetTitle())
		fmt.Printf("  Description: %s\n", task.GetDescription())
		fmt.Printf("  Priority: ")
		switch task.GetPriority() {
		case tms.TaskPriorityLow:
			fmt.Println("Low")
		case tms.TaskPriorityMedium:
			fmt.Println("Medium")
		case tms.TaskPriorityHigh:
			fmt.Println("High")
		case tms.TaskPriorityUrgent:
			fmt.Println("Urgent")
		}
		fmt.Printf("  Status: ")
		switch task.GetStatus() {
		case tms.TaskStatusPending:
			fmt.Println("Pending")
		case tms.TaskStatusInProgress:
			fmt.Println("In Progress")
		case tms.TaskStatusCompleted:
			fmt.Println("Completed")
		case tms.TaskStatusBlocked:
			fmt.Println("Blocked")
		}
		if task.GetAssignedTo() != nil {
			fmt.Printf("  Assigned to: %s\n", task.GetAssignedTo().GetName())
		} else {
			fmt.Println("  Assigned to: None")
		}
		fmt.Printf("  Due date: %s\n", task.GetDueDate().Format("2006-01-02"))
		if len(task.GetComments()) > 0 {
			fmt.Printf("  Comments (%d):\n", len(task.GetComments()))
			for i, comment := range task.GetComments() {
				fmt.Printf("    %d. %s\n", i+1, comment)
			}
		}
		fmt.Println()
	}

	fmt.Println("=== Demo Completed ===")
}
