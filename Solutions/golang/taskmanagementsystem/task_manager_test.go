package taskmanagementsystem

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSingletonPattern(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	if instance1 != instance2 {
		t.Error("TaskManager should be a singleton")
	}
}

func TestCreateAndGetTask(t *testing.T) {
	tm := GetInstance()

	dueDate := time.Now().Add(24 * time.Hour)
	task := NewTask("task1", "Test Task", "This is a test task", dueDate, TaskPriorityHigh, nil)

	tm.CreateTask(task)

	retrievedTask := tm.GetTask("task1")
	if retrievedTask == nil {
		t.Fatal("Task should be created and retrievable")
	}

	if retrievedTask.GetId() != "task1" {
		t.Errorf("Expected task ID 'task1', got '%s'", retrievedTask.GetId())
	}

	if retrievedTask.GetTitle() != "Test Task" {
		t.Errorf("Expected title 'Test Task', got '%s'", retrievedTask.GetTitle())
	}
}

func TestUpdateTask(t *testing.T) {
	tm := GetInstance()

	dueDate := time.Now().Add(24 * time.Hour)
	task := NewTask("task2", "Original Title", "Original Description", dueDate, TaskPriorityMedium, nil)

	tm.CreateTask(task)

	// Update task
	task.SetTitle("Updated Title")
	task.SetDescription("Updated Description")
	task.SetPriority(TaskPriorityUrgent)

	success := tm.UpdateTask(task)
	if !success {
		t.Error("Task update should succeed")
	}

	updatedTask := tm.GetTask("task2")
	if updatedTask.GetTitle() != "Updated Title" {
		t.Errorf("Expected updated title, got '%s'", updatedTask.GetTitle())
	}

	if updatedTask.GetPriority() != TaskPriorityUrgent {
		t.Error("Priority should be updated")
	}
}

func TestDeleteTask(t *testing.T) {
	tm := GetInstance()

	dueDate := time.Now().Add(24 * time.Hour)
	task := NewTask("task3", "Task to Delete", "This task will be deleted", dueDate, TaskPriorityLow, nil)

	tm.CreateTask(task)

	success := tm.DeleteTask("task3")
	if !success {
		t.Error("Task deletion should succeed")
	}

	deletedTask := tm.GetTask("task3")
	if deletedTask != nil {
		t.Error("Deleted task should not be retrievable")
	}
}

func TestUserManagement(t *testing.T) {
	tm := GetInstance()

	user := NewUser("user1", "John Doe", "john@example.com")
	tm.AddUser(user)

	retrievedUser := tm.GetUser("user1")
	if retrievedUser == nil {
		t.Fatal("User should be created and retrievable")
	}

	if retrievedUser.GetName() != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", retrievedUser.GetName())
	}

	if retrievedUser.GetEmail() != "john@example.com" {
		t.Errorf("Expected email 'john@example.com', got '%s'", retrievedUser.GetEmail())
	}
}

func TestAssignTask(t *testing.T) {
	tm := GetInstance()

	user := NewUser("user2", "Jane Smith", "jane@example.com")
	tm.AddUser(user)

	dueDate := time.Now().Add(48 * time.Hour)
	task := NewTask("task4", "Assigned Task", "This task will be assigned", dueDate, TaskPriorityHigh, nil)
	tm.CreateTask(task)

	success := tm.AssignTask("task4", "user2")
	if !success {
		t.Fatal("Task assignment should succeed")
	}

	retrievedTask := tm.GetTask("task4")
	if retrievedTask.GetAssignedTo() == nil {
		t.Fatal("Task should have an assigned user")
	}

	if retrievedTask.GetAssignedTo().GetId() != "user2" {
		t.Error("Task should be assigned to user2")
	}

	retrievedUser := tm.GetUser("user2")
	assignedTasks := retrievedUser.GetAssignedTasks()
	if len(assignedTasks) != 1 || assignedTasks[0] != "task4" {
		t.Error("User should have the assigned task in their list")
	}
}

func TestSearchTasks(t *testing.T) {
	tm := GetInstance()

	dueDate := time.Now().Add(24 * time.Hour)
	task1 := NewTask("search1", "Important Meeting", "Discuss project roadmap", dueDate, TaskPriorityHigh, nil)
	task2 := NewTask("search2", "Code Review", "Review pull requests", dueDate, TaskPriorityMedium, nil)
	task3 := NewTask("search3", "Team Meeting", "Weekly sync", dueDate, TaskPriorityLow, nil)

	tm.CreateTask(task1)
	tm.CreateTask(task2)
	tm.CreateTask(task3)

	results := tm.SearchTasks("Meeting")
	if len(results) != 2 {
		t.Errorf("Expected 2 tasks with 'Meeting', got %d", len(results))
	}

	results = tm.SearchTasks("Review")
	if len(results) != 1 {
		t.Errorf("Expected 1 task with 'Review', got %d", len(results))
	}
}

func TestFilterTasksByStatus(t *testing.T) {
	tm := GetInstance()

	dueDate := time.Now().Add(24 * time.Hour)
	task1 := NewTask("status1", "Pending Task", "Task 1", dueDate, TaskPriorityHigh, nil)
	task2 := NewTask("status2", "In Progress Task", "Task 2", dueDate, TaskPriorityMedium, nil)
	task3 := NewTask("status3", "Another Pending Task", "Task 3", dueDate, TaskPriorityLow, nil)

	task1.SetStatus(TaskStatusPending)
	task2.SetStatus(TaskStatusInProgress)
	task3.SetStatus(TaskStatusPending)

	tm.CreateTask(task1)
	tm.CreateTask(task2)
	tm.CreateTask(task3)

	pendingTasks := tm.FilterTasksByStatus(TaskStatusPending)
	if len(pendingTasks) < 2 {
		t.Errorf("Expected at least 2 pending tasks, got %d", len(pendingTasks))
	}

	inProgressTasks := tm.FilterTasksByStatus(TaskStatusInProgress)
	found := false
	for _, task := range inProgressTasks {
		if task.GetId() == "status2" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Should find in-progress task")
	}
}

func TestFilterTasksByPriority(t *testing.T) {
	tm := GetInstance()

	dueDate := time.Now().Add(24 * time.Hour)
	task1 := NewTask("priority1", "High Priority Task", "Task 1", dueDate, TaskPriorityHigh, nil)
	task2 := NewTask("priority2", "Low Priority Task", "Task 2", dueDate, TaskPriorityLow, nil)

	tm.CreateTask(task1)
	tm.CreateTask(task2)

	highPriorityTasks := tm.FilterTasksByPriority(TaskPriorityHigh)
	found := false
	for _, task := range highPriorityTasks {
		if task.GetId() == "priority1" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Should find high priority task")
	}
}

func TestFilterTasksByAssignedUser(t *testing.T) {
	tm := GetInstance()

	user := NewUser("user3", "Bob Brown", "bob@example.com")
	tm.AddUser(user)

	dueDate := time.Now().Add(24 * time.Hour)
	task1 := NewTask("assigned1", "User Task 1", "Task for user3", dueDate, TaskPriorityHigh, nil)
	task2 := NewTask("assigned2", "User Task 2", "Another task for user3", dueDate, TaskPriorityMedium, nil)

	tm.CreateTask(task1)
	tm.CreateTask(task2)
	tm.AssignTask("assigned1", "user3")
	tm.AssignTask("assigned2", "user3")

	userTasks := tm.FilterTasksByAssignedUser("user3")
	if len(userTasks) != 2 {
		t.Errorf("Expected 2 tasks for user3, got %d", len(userTasks))
	}
}

func TestFilterTasksByDueDateRange(t *testing.T) {
	tm := GetInstance()

	startDate := time.Now()
	endDate := startDate.Add(48 * time.Hour)

	task1 := NewTask("date1", "Task 1", "Due tomorrow", startDate.Add(24*time.Hour), TaskPriorityHigh, nil)
	task2 := NewTask("date2", "Task 2", "Due in 3 days", startDate.Add(72*time.Hour), TaskPriorityMedium, nil)

	tm.CreateTask(task1)
	tm.CreateTask(task2)

	tasksInRange := tm.FilterTasksByDueDateRange(startDate, endDate)
	found := false
	for _, task := range tasksInRange {
		if task.GetId() == "date1" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Should find task1 in date range")
	}
}

func TestMarkTaskAsCompleted(t *testing.T) {
	tm := GetInstance()

	user := NewUser("user4", "Alice Wilson", "alice@example.com")
	tm.AddUser(user)

	dueDate := time.Now().Add(24 * time.Hour)
	task := NewTask("complete1", "Task to Complete", "This task will be completed", dueDate, TaskPriorityHigh, nil)
	tm.CreateTask(task)
	tm.AssignTask("complete1", "user4")

	success := tm.MarkTaskAsCompleted("complete1")
	if !success {
		t.Fatal("Task should be marked as completed")
	}

	completedTask := tm.GetTask("complete1")
	if completedTask.GetStatus() != TaskStatusCompleted {
		t.Error("Task status should be completed")
	}

	history := tm.GetTaskHistory("user4")
	if len(history) != 1 {
		t.Errorf("Expected 1 task in history, got %d", len(history))
	}

	if len(history) > 0 && history[0].GetId() != "complete1" {
		t.Error("History should contain the completed task")
	}
}

func TestTaskComments(t *testing.T) {
	dueDate := time.Now().Add(24 * time.Hour)
	task := NewTask("comment1", "Task with Comments", "Testing comments", dueDate, TaskPriorityMedium, nil)

	task.AddComment("First comment")
	task.AddComment("Second comment")

	comments := task.GetComments()
	if len(comments) != 2 {
		t.Errorf("Expected 2 comments, got %d", len(comments))
	}

	if comments[0] != "First comment" {
		t.Errorf("Expected 'First comment', got '%s'", comments[0])
	}
}

func TestConcurrentTaskCreation(t *testing.T) {
	tm := GetInstance()

	var wg sync.WaitGroup
	numTasks := 100

	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			dueDate := time.Now().Add(24 * time.Hour)
			task := NewTask(
				fmt.Sprintf("concurrent%d", index),
				fmt.Sprintf("Task %d", index),
				fmt.Sprintf("Description %d", index),
				dueDate,
				TaskPriorityMedium,
				nil,
			)
			tm.CreateTask(task)
		}(i)
	}

	wg.Wait()

	// Verify all tasks were created
	allTasks := tm.GetAllTasks()
	concurrentCount := 0
	for _, task := range allTasks {
		if len(task.GetId()) >= 10 && task.GetId()[:10] == "concurrent" {
			concurrentCount++
		}
	}

	if concurrentCount != numTasks {
		t.Errorf("Expected %d concurrent tasks, got %d", numTasks, concurrentCount)
	}
}

func TestConcurrentTaskOperations(t *testing.T) {
	tm := GetInstance()

	// Create initial tasks
	dueDate := time.Now().Add(24 * time.Hour)
	for i := 0; i < 10; i++ {
		task := NewTask(
			fmt.Sprintf("op%d", i),
			fmt.Sprintf("Operation Task %d", i),
			"Description",
			dueDate,
			TaskPriorityMedium,
			nil,
		)
		tm.CreateTask(task)
	}

	var wg sync.WaitGroup

	// Concurrent reads
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			taskId := fmt.Sprintf("op%d", index%10)
			tm.GetTask(taskId)
		}(i)
	}

	// Concurrent updates
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			taskId := fmt.Sprintf("op%d", index%10)
			task := tm.GetTask(taskId)
			if task != nil {
				task.SetTitle(fmt.Sprintf("Updated Title %d", index))
				tm.UpdateTask(task)
			}
		}(i)
	}

	// Concurrent searches
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tm.SearchTasks("Operation")
		}()
	}

	wg.Wait()

	// Verify system integrity
	allTasks := tm.GetAllTasks()
	if len(allTasks) == 0 {
		t.Error("Tasks should still exist after concurrent operations")
	}
}

func TestUserTaskReassignment(t *testing.T) {
	tm := GetInstance()

	user1 := NewUser("reassign1", "User One", "user1@example.com")
	user2 := NewUser("reassign2", "User Two", "user2@example.com")
	tm.AddUser(user1)
	tm.AddUser(user2)

	dueDate := time.Now().Add(24 * time.Hour)
	task := NewTask("reassignTask", "Reassign Task", "Task to be reassigned", dueDate, TaskPriorityHigh, nil)
	tm.CreateTask(task)

	// Assign to user1
	tm.AssignTask("reassignTask", "reassign1")

	retrievedUser1 := tm.GetUser("reassign1")
	if len(retrievedUser1.GetAssignedTasks()) != 1 {
		t.Error("User1 should have 1 assigned task")
	}

	// Reassign to user2
	tm.AssignTask("reassignTask", "reassign2")

	retrievedUser1 = tm.GetUser("reassign1")
	retrievedUser2 := tm.GetUser("reassign2")

	if len(retrievedUser1.GetAssignedTasks()) != 0 {
		t.Error("User1 should have 0 assigned tasks after reassignment")
	}

	if len(retrievedUser2.GetAssignedTasks()) != 1 {
		t.Error("User2 should have 1 assigned task after reassignment")
	}
}
