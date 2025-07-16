#include <gtest/gtest.h>
#include "Logger.h"
#include "ConsoleAppender.h"
#include "FileAppender.h"
#include "LogMessage.h"
#include <fstream>
#include <sstream>
#include <memory>

class LoggingFrameworkTest : public ::testing::Test {
protected:
    void SetUp() override {
        // Clean up any existing test files
        std::remove("test_log.txt");
    }

    void TearDown() override {
        // Clean up test files
        std::remove("test_log.txt");
    }
};

// Test LogMessage creation and formatting
TEST_F(LoggingFrameworkTest, LogMessageCreationAndFormatting) {
    LogMessage msg(LogLevel::INFO, "Test message", "TestSource");
    
    std::string formatted = msg.getFormattedMessage();
    EXPECT_TRUE(formatted.find("INFO") != std::string::npos);
    EXPECT_TRUE(formatted.find("Test message") != std::string::npos);
    EXPECT_TRUE(formatted.find("TestSource") != std::string::npos);
}

// Test Logger creation and basic functionality
TEST_F(LoggingFrameworkTest, LoggerCreation) {
    Logger logger("TestLogger", LogLevel::DEBUG);
    
    // Test that logger can be created without throwing
    EXPECT_NO_THROW(logger.debug("Debug message", "TestSource"));
    EXPECT_NO_THROW(logger.info("Info message", "TestSource"));
    EXPECT_NO_THROW(logger.warning("Warning message", "TestSource"));
    EXPECT_NO_THROW(logger.error("Error message", "TestSource"));
    EXPECT_NO_THROW(logger.fatal("Fatal message", "TestSource"));
}

// Test FileAppender functionality
TEST_F(LoggingFrameworkTest, FileAppenderWritesToFile) {
    auto fileAppender = std::make_shared<FileAppender>("test_log.txt");
    Logger logger("FileLogger", LogLevel::INFO);
    logger.addAppender(fileAppender);
    
    logger.info("Test file message", "FileTest");
    
    // Read the file and verify content
    std::ifstream file("test_log.txt");
    ASSERT_TRUE(file.is_open());
    
    std::string line;
    std::getline(file, line);
    EXPECT_TRUE(line.find("INFO") != std::string::npos);
    EXPECT_TRUE(line.find("Test file message") != std::string::npos);
    EXPECT_TRUE(line.find("FileTest") != std::string::npos);
    
    file.close();
}

// Test ConsoleAppender functionality (redirect cout for testing)
TEST_F(LoggingFrameworkTest, ConsoleAppenderWritesToConsole) {
    // Redirect cout to stringstream for testing
    std::stringstream buffer;
    std::streambuf* old = std::cout.rdbuf(buffer.rdbuf());
    
    auto consoleAppender = std::make_shared<ConsoleAppender>();
    Logger logger("ConsoleLogger", LogLevel::INFO);
    logger.addAppender(consoleAppender);
    
    logger.info("Test console message", "ConsoleTest");
    
    // Restore cout
    std::cout.rdbuf(old);
    
    std::string output = buffer.str();
    EXPECT_TRUE(output.find("INFO") != std::string::npos);
    EXPECT_TRUE(output.find("Test console message") != std::string::npos);
    EXPECT_TRUE(output.find("ConsoleTest") != std::string::npos);
}

// Test log level filtering
TEST_F(LoggingFrameworkTest, LogLevelFiltering) {
    std::stringstream buffer;
    std::streambuf* old = std::cout.rdbuf(buffer.rdbuf());
    
    auto consoleAppender = std::make_shared<ConsoleAppender>();
    Logger logger("FilterLogger", LogLevel::WARNING);  // Only WARNING and above
    logger.addAppender(consoleAppender);
    
    logger.debug("Debug message", "FilterTest");    // Should be filtered out
    logger.info("Info message", "FilterTest");      // Should be filtered out
    logger.warning("Warning message", "FilterTest"); // Should appear
    logger.error("Error message", "FilterTest");     // Should appear
    
    std::cout.rdbuf(old);
    
    std::string output = buffer.str();
    EXPECT_TRUE(output.find("Debug message") == std::string::npos);
    EXPECT_TRUE(output.find("Info message") == std::string::npos);
    EXPECT_TRUE(output.find("Warning message") != std::string::npos);
    EXPECT_TRUE(output.find("Error message") != std::string::npos);
}

// Test multiple appenders
TEST_F(LoggingFrameworkTest, MultipleAppenders) {
    std::stringstream buffer;
    std::streambuf* old = std::cout.rdbuf(buffer.rdbuf());
    
    auto consoleAppender = std::make_shared<ConsoleAppender>();
    auto fileAppender = std::make_shared<FileAppender>("test_log.txt");
    
    Logger logger("MultiLogger", LogLevel::INFO);
    logger.addAppender(consoleAppender);
    logger.addAppender(fileAppender);
    
    logger.info("Multi appender test", "MultiTest");
    
    std::cout.rdbuf(old);
    
    // Check console output
    std::string consoleOutput = buffer.str();
    EXPECT_TRUE(consoleOutput.find("Multi appender test") != std::string::npos);
    
    // Check file output
    std::ifstream file("test_log.txt");
    ASSERT_TRUE(file.is_open());
    
    std::string fileLine;
    std::getline(file, fileLine);
    EXPECT_TRUE(fileLine.find("Multi appender test") != std::string::npos);
    
    file.close();
}

// Test all log levels
TEST_F(LoggingFrameworkTest, AllLogLevels) {
    std::stringstream buffer;
    std::streambuf* old = std::cout.rdbuf(buffer.rdbuf());
    
    auto consoleAppender = std::make_shared<ConsoleAppender>();
    Logger logger("LevelLogger", LogLevel::DEBUG);  // Allow all levels
    logger.addAppender(consoleAppender);
    
    logger.debug("Debug test", "LevelTest");
    logger.info("Info test", "LevelTest");
    logger.warning("Warning test", "LevelTest");
    logger.error("Error test", "LevelTest");
    logger.fatal("Fatal test", "LevelTest");
    
    std::cout.rdbuf(old);
    
    std::string output = buffer.str();
    EXPECT_TRUE(output.find("DEBUG") != std::string::npos);
    EXPECT_TRUE(output.find("INFO") != std::string::npos);
    EXPECT_TRUE(output.find("WARNING") != std::string::npos);
    EXPECT_TRUE(output.find("ERROR") != std::string::npos);
    EXPECT_TRUE(output.find("FATAL") != std::string::npos);
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
