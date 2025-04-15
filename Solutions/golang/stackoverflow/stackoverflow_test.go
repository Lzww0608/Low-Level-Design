package stackoverflow

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// 简单的顺序测试，不使用并行执行
func TestStackOverflowSimple(t *testing.T) {
	// 创建一个新的StackOverflow实例
	so := NewStackOverflow()
	
	// 测试用户注册
	err := so.RegisterUser("user1", "user1@example.com")
	assert.NoError(t, err)
	
	err = so.RegisterUser("user2", "user2@example.com")
	assert.NoError(t, err)
	
	// 获取已注册用户
	user1 := so.GetUserByUsername("user1")
	user2 := so.GetUserByUsername("user2")
	
	assert.NotNil(t, user1, "user1 应该已注册")
	assert.NotNil(t, user2, "user2 应该已注册")
	
	// 测试提问
	question, err := so.AskQuestion(user1, "测试问题", "这是一个测试问题的内容", []string{"测试", "golang"})
	assert.NoError(t, err)
	assert.NotNil(t, question)
	
	// 测试回答
	answer := NewAnswer(user2, question, "这是对测试问题的回答")
	err = question.AddAnswer(answer)
	assert.NoError(t, err)
	
	// 测试投票
	err = so.VoteQuestion(user2, question, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, question.GetVoteCount())
	
	// 测试评论
	comment, err := so.AddComment(user1, question, "这是对问题的评论")
	assert.NoError(t, err)
	assert.NotNil(t, comment)
	
	// 测试搜索
	results := so.SearchQuestions("测试")
	assert.NotEmpty(t, results)
	
	// 测试接受回答
	err = so.AcceptAnswer(answer)
	assert.NoError(t, err)
	assert.True(t, answer.IsAccepted())
}
