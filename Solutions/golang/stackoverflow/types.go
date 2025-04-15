package stackoverflow

import (
	"fmt"
	"time"
	"sync"
	"github.com/google/uuid"
)

// Interface 
type Votable interface {
	Vote(user *User, value int) error
	GetVoteCount() int
}

type Commentable interface {
	AddComment(comment *Comment) error
	GetComments() []*Comment
}


// Reputation constants
const (
	QuestionReputation = 5 
	AnswerReputation = 10
	CommentReputation = 2
)

// Vote type 
type Vote struct {
	User *User 
	Value int
}

// Comment type 
type Comment struct {
	ID         string
	Content    string
	Author     *User
	Creation   time.Time
	mu         sync.RWMutex
}

func NewComment(author *User, content string) *Comment {
	return &Comment {
		ID:      	generateID(),
		Content: 	content,
		Author:  	author,
		Creation:   time.Now(),
	}
}

// Tag type 
type Tag struct {
	ID   string 
	Name string 
}

func NewTag(name string) *Tag {
	return &Tag {
		ID:    generateID(),
		Name:  name,
	}
}


// User type 
type User struct {
	ID         string 
	Username   string
	Email      string 
	Reputation int
	Questions  []*Question
	Answers    []*Answer
	Comments   []*Comment
	mu         sync.RWMutex
}


func NewUser(id, username, email string) *User {
	return &User {
		ID:       id,
		Username: username,
		Email:    email,
		Reputation: 0,
		Questions:   []*Question{},
		Answers:     []*Answer{},
		Comments:    []*Comment{},
	}
}

func (u *User) UpdateReputation(value int) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.Reputation += value 
	if u.Reputation < 0 {
		u.Reputation = 0 
	}
}

func (u *User) GetReputation() int {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.Reputation
}

func (u *User) AddQuestion(q *Question) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.Questions = append(u.Questions, q)
	u.UpdateReputation(QuestionReputation)
}

func (u *User) AddAnswer(a *Answer) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.Answers = append(u.Answers, a)
	u.UpdateReputation(AnswerReputation)
}

func (u *User) AddComment(c *Comment) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.Comments = append(u.Comments, c)
	u.UpdateReputation(CommentReputation)
}


func (u *User) GetQuestions() []*Question {
	u.mu.RLock()
	defer u.mu.RUnlock()

	questions := make([]*Question, len(u.Questions))
	copy(questions, u.Questions)
	return questions
}


func (u *User) GetAnswers() []*Answer {
	u.mu.RLock()
	defer u.mu.RUnlock()

	answers := make([]*Answer, len(u.Answers))
	copy(answers, u.Answers)
	return answers
}

func (u *User) GetComments() []*Comment {
	u.mu.RLock()
	defer u.mu.RUnlock()

	comments := make([]*Comment, len(u.Comments))
	copy(comments, u.Comments)
	return comments
}

// Question type 
type Question struct {
	ID         string
	Title      string
	Content    string
	Author     *User
	CreationDate time.Time
	Answers    []*Answer
	Comments   []*Comment
	Tags       []*Tag
	Votes      []*Vote
	mu         sync.RWMutex
	
}


func NewQuestion(author *User, title, content string, tagNames []string) *Question {
	q := &Question {
		ID:         generateID(),
		Title:      title,
		Content:    content,
		Author:     author,
		CreationDate: time.Now(),
		Answers:    []*Answer{},
		Comments:   []*Comment{},
		Tags:       []*Tag{},
		Votes:      []*Vote{},
	}

	for _, tagName := range tagNames {
		tag := NewTag(tagName)
		q.Tags = append(q.Tags, tag)
	}
	
	return q 
}

func (q *Question) AddAnswer(a *Answer) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	// Check for duplicate answer 
	for _, answer := range q.Answers {
		if answer.ID == a.ID {
			return nil
		}
	}
	
	q.Answers = append(q.Answers, a)
	return nil
}

func (q *Question) Vote(user *User, value int) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if value != 1 && value != -1 {
		return fmt.Errorf("vote value must be 1 or -1")
	}

	// Remove existing vote if any 
	for i, vote := range q.Votes {
		if vote.User.ID == user.ID {
			q.Votes = append(q.Votes[:i], q.Votes[i+1:]...)
			break
		}
	}

	// Add new vote 
	q.Votes = append(q.Votes, &Vote{
		User: user,
		Value: value,
	})
	q.Author.UpdateReputation(value * 5)
	return nil
}

func (q *Question) GetVoteCount() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	count := 0
	for _, vote := range q.Votes {
		count += vote.Value
	}
	return count
}


func (q *Question) AddComment(c *Comment) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.Comments = append(q.Comments, c)
	return nil
}

func (q *Question) GetComments() []*Comment {
	q.mu.RLock()
	defer q.mu.RUnlock()
	
	comments := make([]*Comment, len(q.Comments))
	copy(comments, q.Comments)
	return comments
}

func (q *Question) GetAnswers() []*Answer {
	q.mu.RLock()
	defer q.mu.RUnlock()

	answers := make([]*Answer, len(q.Answers))
	copy(answers, q.Answers)
	return answers
}

func (q *Question) GetTags() []*Tag {
	q.mu.RLock()
	defer q.mu.RUnlock()
	
	tags := make([]*Tag, len(q.Tags))
	copy(tags, q.Tags)
	return tags
}


// Utility
func generateID() string {
	return uuid.New().String()
}

