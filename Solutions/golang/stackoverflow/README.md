# Designing Stack Overflow

## Requirements
1. Users can post questions, answer questions, and comment on questions and answers.
2. Users can vote on questions and answers.
3. Questions should have tags associated with them.
4. Users can search for questions based on keywords, tags, or user profiles.
5. The system should assign reputation score to users based on their activity and the quality of their contributions.
6. The system should handle concurrent access and ensure data consistency.

## 系统架构与实现

本系统采用Go语言实现了一个简化版的Stack Overflow，支持问答、评论、投票等核心功能，并使用互斥锁保证并发安全。

### 核心接口

#### Votable 接口
可投票对象必须实现的接口：
```go
type Votable interface {
    Vote(user *User, value int) error
    GetVoteCount() int
}
```

#### Commentable 接口
可评论对象必须实现的接口：
```go
type Commentable interface {
    AddComment(comment *Comment) error
    GetComments() []*Comment
}
```

### 核心数据结构

#### Vote 结构体
表示投票：
```go
type Vote struct {
    User  *User 
    Value int    // 值为+1或-1
}
```

#### Comment 结构体
表示评论：
```go
type Comment struct {
    ID        string
    Content   string
    Author    *User
    Creation  time.Time
    mu        sync.RWMutex
}
```

#### Tag 结构体
表示问题标签：
```go
type Tag struct {
    ID   string 
    Name string 
}
```

#### User 结构体
表示用户：
```go
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
```

#### Question 结构体
表示提问：
```go
type Question struct {
    ID           string
    Title        string
    Content      string
    Author       *User
    CreationDate time.Time
    Answers      []*Answer
    Comments     []*Comment
    Tags         []*Tag
    Votes        []*Vote
    mu           sync.RWMutex
}
```

#### Answer 结构体
表示回答：
```go
type Answer struct {
    ID           string
    Content      string
    Author       *User 
    Question     *Question
    CreationDate time.Time
    isAccepted   bool 
    comments     []*Comment
    votes        []*Vote
    mu           sync.RWMutex
}
```

#### StackOverflow 结构体
系统主体结构，管理所有实体：
```go
type StackOverflow struct {
    users      map[string]*User 
    questions  map[string]*Question
    answers    map[string]*Answer 
    tags       map[string]*Tag 
    mu         sync.RWMutex
}
```

### 主要功能实现

#### 用户管理
- `RegisterUser(username, email string) error`: 注册新用户
- `GetUserByUsername(username string) *User`: 根据用户名查找用户

#### 问答系统
- `AskQuestion(user *User, title, content string, tagNames []string) (*Question, error)`: 提问
- `AddComment(user *User, target Commentable, content string) (*Comment, error)`: 添加评论
- `VoteQuestion(user *User, question *Question, value int) error`: 对问题投票
- `VoteAnswer(user *User, answer *Answer, value int) error`: 对回答投票
- `AcceptAnswer(answer *Answer) error`: 接受回答

#### 搜索功能
- `SearchQuestions(query string) []*Question`: 搜索问题
- `GetQuestionByUser(user *User) []*Question`: 获取用户的问题

### 并发控制
系统使用读写锁处理并发访问，主要解决方案包括：

1. 每个实体（User, Question, Answer, Comment）都有自己的互斥锁
2. 读取操作使用读锁（RLock），修改操作使用写锁（Lock）
3. 使用goroutine和channel解决死锁问题，特别是在循环依赖场景

例如，在添加问题时解决循环依赖：
```go
// 解决循环依赖问题：使用goroutine异步添加问题到用户
go func() {
    user.AddQuestion(question)
}()
```

### 数据一致性保证
1. 所有对象的修改都通过方法进行，而不是直接访问字段
2. 返回切片时使用深拷贝，避免外部修改影响内部数据
3. 使用原子操作和事务性更新确保数据一致性

### 系统注意事项
1. 投票只能是+1或-1
2. 回答只能被接受一次
3. 用户声望（Reputation）不会低于0
4. 问题、回答和评论一旦创建就不能被修改或删除