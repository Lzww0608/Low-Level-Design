package stackoverflow

import (
	"fmt"
	"strings"
	"sync"
)

type StackOverflow struct {
	users      map[string]*User 
	questions  map[string]*Question
	answers    map[string]*Answer 
	tags 	   map[string]*Tag 
	mu         sync.RWMutex
}

func NewStackOverflow() *StackOverflow {
	return &StackOverflow {
		users:      make(map[string]*User),
		questions:  make(map[string]*Question),
		answers:    make(map[string]*Answer),
		tags:       make(map[string]*Tag),	
	}
}

func (s *StackOverflow) RegisterUser(username, email string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := generateID()
	user := NewUser(id, username, email)
	s.users[id] = user
	return nil
}

func (s *StackOverflow) AskQuestion(user *User, title, content string, tagNames []string) (*Question, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	question := NewQuestion(user, title, content, tagNames)
	
	// 解决循环依赖问题：使用goroutine异步添加问题到用户
	go func() {
		user.AddQuestion(question)
	}()

	// 保存问题到映射中
	s.questions[question.ID] = question

	// Register tags 
	for _, tag := range question.GetTags() {
		s.tags[tag.Name] = tag
	}
	
	return question, nil 
}

func (s *StackOverflow) AddComment(user *User, target Commentable, content string) (*Comment, error) {
	if user == nil || target == nil {
		return nil, fmt.Errorf("user or target is nil")
	}

	comment := NewComment(user, content)
	
	// 使用goroutine防止死锁
	commentErr := make(chan error, 1)
	go func() {
		commentErr <- target.AddComment(comment)
	}()
	
	// 使用另一个goroutine异步添加评论到用户
	go func() {
		user.AddComment(comment)
	}()
	
	// 等待评论添加操作完成
	if err := <-commentErr; err != nil {
		return nil, err
	}

	return comment, nil 
}


func (s *StackOverflow) VoteQuestion(user *User, question *Question, value int) error {
	if user == nil || question == nil {
		return fmt.Errorf("user or question is nil")
	}
	
	// 使用goroutine防止死锁
	voteErr := make(chan error, 1)
	go func() {
		voteErr <- question.Vote(user, value)
	}()
	
	return <-voteErr
}

func (s *StackOverflow) VoteAnswer(user *User, answer *Answer, value int) error {
	if user == nil || answer == nil {
		return fmt.Errorf("user or answer is nil")
	}

	// 使用goroutine防止死锁
	voteErr := make(chan error, 1)
	go func() {
		voteErr <- answer.Vote(user, value)
	}()
	
	return <-voteErr
}	

func (s *StackOverflow) AcceptAnswer(answer *Answer) error {
	if answer == nil {
		return fmt.Errorf("answer is nil")
	}

	// 使用goroutine防止死锁
	acceptErr := make(chan error, 1)
	go func() {
		acceptErr <- answer.MarkAsAccepted()
	}()
	
	return <-acceptErr
}


func (s *StackOverflow) SearchQuestions(query string) []*Question {
	s.mu.RLock()
	defer s.mu.RUnlock()

	query = strings.ToLower(query)
	var results []*Question 

	for _, q := range s.questions {
		if strings.Contains(strings.ToLower(q.Title), query) || 
			strings.Contains(strings.ToLower(q.Content), query) {
			results = append(results, q)
			continue
		}

		// Search in tags 
		for _, tag := range q.GetTags() {
			if strings.EqualFold(tag.Name, query) {
				results = append(results, q)
				break
			}
		}
	}

	return results 
}

func (s *StackOverflow) GetQuestionByUser(user *User) []*Question {
	if user == nil {
		return nil
	}
	
	return user.GetQuestions()
}

// 添加一个方法根据用户名获取用户
func (s *StackOverflow) GetUserByUsername(username string) *User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	for _, user := range s.users {
		if user.Username == username {
			return user
		}
	}
	return nil
}
