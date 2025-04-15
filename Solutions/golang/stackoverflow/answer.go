package stackoverflow

import (
	"fmt"
	"time"
	"sync"
)



type Answer struct {
	ID         string
	Content    string
	Author     *User 
	Question   *Question
	CreationDate time.Time
	isAccepted bool 
	comments   []*Comment
	votes      []*Vote
	mu         sync.RWMutex
}

func NewAnswer(author *User, question *Question, content string) *Answer {
	return &Answer {
		ID:         generateID(),
		Content:    content,
		Author:     author,
		Question:   question,
		CreationDate: time.Now(),
		isAccepted: false,
		comments:   []*Comment{},
		votes:      []*Vote{},
	}
}


func (a *Answer) Vote(user *User, value int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if value != 1 && value != -1 {
		return fmt.Errorf("vote value must be 1 or -1")
	}

	// Remove existing vote if any 
	for i, vote := range a.votes {
		if vote.User.ID == user.ID {
			a.votes = append(a.votes[:i], a.votes[i+1:]...)
			break
		}
	}

	// Add new vote 
	a.votes = append(a.votes, &Vote{
		User: user,
		Value: value,
	})
	a.Author.UpdateReputation(value * 10)
	return nil
}

func (a *Answer) GetVoteCount() int {
	a.mu.RLock()	
	defer a.mu.RUnlock()

	count := 0
	for _, vote := range a.votes {
		count += vote.Value
	}
	return count
}

func (a *Answer) AddComment(comment *Comment) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.comments = append(a.comments, comment)
	return nil
}

func (a *Answer) GetComments() []*Comment {
	a.mu.RLock()
	defer a.mu.RUnlock()

	comments := make([]*Comment, len(a.comments))
	copy(comments, a.comments)
	return comments
}

func (a *Answer) MarkAsAccepted() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.isAccepted {
		return fmt.Errorf("answer is already accepted")
	}

	a.isAccepted = true
	a.Author.UpdateReputation(15)
	return nil
}


func (a *Answer) IsAccepted() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.isAccepted
}
