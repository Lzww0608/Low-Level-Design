#include <iostream>
#include <string>
#include <gtest/gtest.h>

/**
 * The Subject interface declares common operations for both RealSubject and the
 * Proxy. As long as the client works with RealSubject using this interface,
 * you'll be able to pass it a proxy instead of a real subject.
 */
class Subject {
 public:
  virtual void Request() const = 0;
};
/**
 * The RealSubject contains some core business logic. Usually, RealSubjects are
 * capable of doing some useful work which may also be very slow or sensitive -
 * e.g. correcting input data. A Proxy can solve these issues without any
 * changes to the RealSubject's code.
 */
class RealSubject : public Subject {
 public:
  void Request() const override {
    std::cout << "RealSubject: Handling request.\n";
  }
};
/**
 * The Proxy has an interface identical to the RealSubject.
 */
class Proxy : public Subject {
  /**
   * @var RealSubject
   */
 private:
  RealSubject *real_subject_;

  bool CheckAccess() const {
    // Some real checks should go here.
    std::cout << "Proxy: Checking access prior to firing a real request.\n";
    return true;
  }
  void LogAccess() const {
    std::cout << "Proxy: Logging the time of request.\n";
  }

  /**
   * The Proxy maintains a reference to an object of the RealSubject class. It
   * can be either lazy-loaded or passed to the Proxy by the client.
   */
 public:
  Proxy(RealSubject *real_subject) : real_subject_(new RealSubject(*real_subject)) {
  }

  ~Proxy() {
    delete real_subject_;
  }
  /**
   * The most common applications of the Proxy pattern are lazy loading,
   * caching, controlling the access, logging, etc. A Proxy can perform one of
   * these things and then, depending on the result, pass the execution to the
   * same method in a linked RealSubject object.
   */
  void Request() const override {
    if (this->CheckAccess()) {
      this->real_subject_->Request();
      this->LogAccess();
    }
  }
};
/**
 * The client code is supposed to work with all objects (both subjects and
 * proxies) via the Subject interface in order to support both real subjects and
 * proxies. In real life, however, clients mostly work with their real subjects
 * directly. In this case, to implement the pattern more easily, you can extend
 * your proxy from the real subject's class.
 */
void ClientCode(const Subject &subject) {
  // ...
  subject.Request();
  // ...
}


int main() {
  std::cout << "客户端: 执行真实主题的代码:\n";
  RealSubject *real_subject = new RealSubject;
  ClientCode(*real_subject);
  std::cout << "\n";
  
  std::cout << "客户端: 执行代理的相同代码:\n";
  Proxy *proxy = new Proxy(real_subject);
  ClientCode(*proxy);
  
  delete real_subject;
  delete proxy;
  
  return 0;
}


// 为了测试，创建一个可以捕获输出的测试类
class ProxyPatternTest : public ::testing::Test {
protected:
  void SetUp() override {
    // 重定向cout到缓冲区
    old_buf = std::cout.rdbuf(buffer.rdbuf());
  }
  
  void TearDown() override {
    // 恢复cout
    std::cout.rdbuf(old_buf);
  }
  
  std::stringstream buffer;
  std::streambuf* old_buf;
};

// 测试真实主题的行为
TEST_F(ProxyPatternTest, RealSubjectTest) {
  RealSubject real_subject;
  ClientCode(real_subject);
  
  std::string output = buffer.str();
  EXPECT_TRUE(output.find("RealSubject: Handling request.") != std::string::npos);
}

// 测试代理的行为
TEST_F(ProxyPatternTest, ProxyTest) {
  RealSubject* real_subject = new RealSubject();
  Proxy proxy(real_subject);
  ClientCode(proxy);
  
  std::string output = buffer.str();
  EXPECT_TRUE(output.find("Proxy: Checking access prior to firing a real request.") != std::string::npos);
  EXPECT_TRUE(output.find("RealSubject: Handling request.") != std::string::npos);
  EXPECT_TRUE(output.find("Proxy: Logging the time of request.") != std::string::npos);
  
  delete real_subject;
}

// 测试代理的访问控制功能
TEST_F(ProxyPatternTest, ProxyAccessControl) {
  // 这里我们可以创建一个模拟的代理类来测试访问控制
  class MockProxy : public Proxy {
  public:
    MockProxy(RealSubject* real_subject) : Proxy(real_subject) {}
    
    // 覆盖CheckAccess方法以返回false
    bool CheckAccess() const {
      std::cout << "MockProxy: 拒绝访问请求。\n";
      return false;
    }
  };
  
  RealSubject* real_subject = new RealSubject();
  MockProxy mock_proxy(real_subject);
  ClientCode(mock_proxy);
  
  std::string output = buffer.str();
  EXPECT_TRUE(output.find("MockProxy: 拒绝访问请求.") != std::string::npos);
  EXPECT_FALSE(output.find("RealSubject: Handling request.") != std::string::npos);
  
  delete real_subject;
}
