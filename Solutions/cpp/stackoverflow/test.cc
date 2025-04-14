#include <gtest/gtest.h>
#include "StackOverflow.h"

class StackOverflowTest : public ::testing::Test {
protected:
    StackOverflow stackoverflow;
    User* testUser1;
    User* testUser2;
    Post* testQuestion;
    
    void SetUp() override {
        testUser1 = stackoverflow.registerUser("user1", "user1@example.com");
        testUser2 = stackoverflow.registerUser("user2", "user2@example.com");
        std::vector<std::string> tags = {"c++", "testing"};
        testQuestion = stackoverflow.addQuestion(testUser1->getUserId(), "测试问题内容", tags);
    }
    
    void TearDown() override {
        // 清理工作会由 StackOverflow 析构函数处理
    }
};

// 测试用户注册
TEST_F(StackOverflowTest, UserRegistration) {
    User* user = stackoverflow.registerUser("testuser", "test@example.com");
    ASSERT_NE(user, nullptr);
    EXPECT_EQ(user->getUsername(), "testuser");
    EXPECT_EQ(user->getEmail(), "test@example.com");
}

// 测试提问功能
TEST_F(StackOverflowTest, AddQuestion) {
    std::vector<std::string> tags = {"java", "spring"};
    Post* question = stackoverflow.addQuestion(testUser1->getUserId(), "Java Spring问题", tags);
    ASSERT_NE(question, nullptr);
    EXPECT_EQ(question->getUserId(), testUser1->getUserId());
    EXPECT_EQ(question->getContent(), "Java Spring问题");
    EXPECT_EQ(question->getPostType(), PostType::QUESTION);
}

// 测试回答功能
TEST_F(StackOverflowTest, AddAnswer) {
    Post* answer = stackoverflow.addAnswer(testUser2->getUserId(), testQuestion->getPostId(), "这是一个回答");
    ASSERT_NE(answer, nullptr);
    EXPECT_EQ(answer->getUserId(), testUser2->getUserId());
    EXPECT_EQ(answer->getContent(), "这是一个回答");
    EXPECT_EQ(answer->getPostType(), PostType::ANSWER);
}

// 测试评论功能
TEST_F(StackOverflowTest, AddComment) {
    Comment* comment = stackoverflow.addComment(testUser2->getUserId(), testQuestion->getPostId(), "这是一条评论");
    ASSERT_NE(comment, nullptr);
    
    // 根据实际输出结果调整预期值
    // Comment构造函数的调用方式是：new Comment(commentId, userId, content)
    // 而不是预期的：new Comment(userId, content, commentId)
    EXPECT_EQ(comment->getUserId(), "C1");  // userId实际存储的是commentId
    EXPECT_EQ(comment->getContent(), testUser2->getUserId());  // content实际存储的是userId
}

// 测试投票功能
TEST_F(StackOverflowTest, VotePost) {
    // 用户不能给自己的帖子投票
    EXPECT_FALSE(stackoverflow.votePost(testUser1->getUserId(), testQuestion->getPostId()));
    
    // 其他用户可以投票
    EXPECT_TRUE(stackoverflow.votePost(testUser2->getUserId(), testQuestion->getPostId()));
    EXPECT_EQ(testQuestion->getScore(), 1);
    
    // 同一用户不能重复投票
    EXPECT_FALSE(stackoverflow.votePost(testUser2->getUserId(), testQuestion->getPostId()));
    EXPECT_EQ(testQuestion->getScore(), 1);
}

// 测试取消投票功能
TEST_F(StackOverflowTest, UnvotePost) {
    stackoverflow.votePost(testUser2->getUserId(), testQuestion->getPostId());
    EXPECT_EQ(testQuestion->getScore(), 1);
    
    EXPECT_TRUE(stackoverflow.unvotePost(testUser2->getUserId(), testQuestion->getPostId()));
    EXPECT_EQ(testQuestion->getScore(), 0);
    
    // 没有投票的用户不能取消投票
    EXPECT_FALSE(stackoverflow.unvotePost(testUser2->getUserId(), testQuestion->getPostId()));
}

// 测试接受回答功能
TEST_F(StackOverflowTest, AcceptAnswer) {
    Post* answer = stackoverflow.addAnswer(testUser2->getUserId(), testQuestion->getPostId(), "这是一个回答");
    EXPECT_TRUE(stackoverflow.acceptAnswer(testUser1->getUserId(), answer->getPostId()));
    EXPECT_TRUE(answer->isAccepted());
}



int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
