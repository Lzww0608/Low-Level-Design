package taskmanagementsystem

import "time"

type TaskStatus int

const (
	TaskStatusPending TaskStatus = iota
	TaskStatusInProgress
	TaskStatusCompleted
	TaskStatusBlocked
)

type TaskPriority int

const (
	TaskPriorityLow TaskPriority = iota
	TaskPriorityMedium
	TaskPriorityHigh
	TaskPriorityUrgent
)

type Task struct {
	id          string
	title       string
	description string
	dueDate     time.Time
	priority    TaskPriority
	status      TaskStatus
	assignedTo  *User
	comments    []string
}

func NewTask(id, title, description string, dueDate time.Time, priority TaskPriority, assignedTo *User) *Task {
	return &Task{id: id, title: title, description: description, dueDate: dueDate, priority: priority, assignedTo: assignedTo}
}

func (t *Task) GetId() string {
	return t.id
}

func (t *Task) GetTitle() string {
	return t.title
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) GetDueDate() time.Time {
	return t.dueDate
}

func (t *Task) GetPriority() TaskPriority {
	return t.priority
}

func (t *Task) GetStatus() TaskStatus {
	return t.status
}

func (t *Task) GetAssignedTo() *User {
	return t.assignedTo
}

func (t *Task) GetComments() []string {
	return t.comments
}

func (t *Task) SetAssignedTo(assignedTo *User) {
	t.assignedTo = assignedTo
}

func (t *Task) SetStatus(status TaskStatus) {
	t.status = status
}

func (t *Task) SetPriority(priority TaskPriority) {
	t.priority = priority
}

func (t *Task) SetTitle(title string) {
	t.title = title
}

func (t *Task) SetDescription(description string) {
	t.description = description
}

func (t *Task) SetDueDate(dueDate time.Time) {
	t.dueDate = dueDate
}

func (t *Task) AddComment(comment string) {
	t.comments = append(t.comments, comment)
}
