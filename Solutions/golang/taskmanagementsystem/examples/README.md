# Task Management System - Demo Example

This directory contains a demonstration of the Task Management System.

## Running the Demo

```bash
go run main.go
```

## What the Demo Shows

The demo application demonstrates all major features of the task management system:

1. **User Creation**: Creates three users (Alice, Bob, Charlie)
2. **Task Creation**: Creates four tasks with different priorities and due dates
3. **Task Assignment**: Assigns tasks to different users
4. **Status Updates**: Updates task statuses to show progress
5. **Comments**: Adds comments to tasks
6. **Search**: Searches for tasks by keyword
7. **Filtering**: 
   - By priority (High, Urgent)
   - By status (In Progress)
   - By assigned user
   - By due date range
8. **Task Completion**: Marks a task as completed
9. **Task History**: Views completed tasks for a user
10. **Task Reassignment**: Reassigns a task from one user to another
11. **Task Updates**: Updates task details (title, description, priority)

## Expected Output

The demo produces detailed output showing:
- User information
- Task creation and assignment
- Search and filter results
- Task status changes
- Complete task summaries with all properties

## Customization

You can modify `main.go` to:
- Create different users and tasks
- Test different scenarios
- Experiment with various features
- Add your own functionality

## Module Setup

This example uses a local module replacement in `go.mod`:

```go
replace taskmanagementsystem => ../
```

This allows the example to use the parent package without publishing it.

