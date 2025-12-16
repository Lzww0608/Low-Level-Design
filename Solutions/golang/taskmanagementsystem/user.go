package taskmanagementsystem

import (
	"fmt"
	"slices"
)

type User struct {
	id            string
	name          string
	email         string
	assignedTasks []string
}

func NewUser(id, name, email string) *User {
	return &User{id: id, name: name, email: email}
}

func (u *User) GetId() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetAssignedTasks() []string {
	return u.assignedTasks
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) AddAssignedTask(task *Task) {
	u.assignedTasks = append(u.assignedTasks, task.GetId())
	task.SetAssignedTo(u)
}

func (u *User) RemoveAssignedTask(task *Task) {
	u.assignedTasks = slices.DeleteFunc(u.assignedTasks, func(id string) bool { return id == task.GetId() })
	task.SetAssignedTo(nil)
}

func (u *User) DisplayUserInfo() {
	fmt.Printf("User ID: %s\n", u.GetId())
	fmt.Printf("Name: %s\n", u.GetName())
	fmt.Printf("Email: %s\n", u.GetEmail())
	fmt.Printf("Assigned Tasks: %v\n", u.GetAssignedTasks())
}
