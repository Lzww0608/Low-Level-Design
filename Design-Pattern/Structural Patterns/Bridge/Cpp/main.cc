#include <iostream>
#include <string>
#include <gtest/gtest.h>

/**
 * The Implementation defines the interface for all implementation classes. It
 * doesn't have to match the Abstraction's interface. In fact, the two
 * interfaces can be entirely different. Typically the Implementation interface
 * provides only primitive operations, while the Abstraction defines higher-
 * level operations based on those primitives.
 */

class Implementation {
 public:
  virtual ~Implementation() {}
  virtual std::string OperationImplementation() const = 0;
};

/**
 * Each Concrete Implementation corresponds to a specific platform and
 * implements the Implementation interface using that platform's API.
 */
class ConcreteImplementationA : public Implementation {
 public:
  std::string OperationImplementation() const override {
    return "ConcreteImplementationA: Here's the result on the platform A.\n";
  }
};
class ConcreteImplementationB : public Implementation {
 public:
  std::string OperationImplementation() const override {
    return "ConcreteImplementationB: Here's the result on the platform B.\n";
  }
};

/**
 * The Abstraction defines the interface for the "control" part of the two class
 * hierarchies. It maintains a reference to an object of the Implementation
 * hierarchy and delegates all of the real work to this object.
 */

class Abstraction {
  /**
   * @var Implementation
   */
 protected:
  Implementation* implementation_;

 public:
  Abstraction(Implementation* implementation) : implementation_(implementation) {
  }

  virtual ~Abstraction() {
  }

  virtual std::string Operation() const {
    return "Abstraction: Base operation with:\n" +
           this->implementation_->OperationImplementation();
  }
};
/**
 * You can extend the Abstraction without changing the Implementation classes.
 */
class ExtendedAbstraction : public Abstraction {
 public:
  ExtendedAbstraction(Implementation* implementation) : Abstraction(implementation) {
  }
  std::string Operation() const override {
    return "ExtendedAbstraction: Extended operation with:\n" +
           this->implementation_->OperationImplementation();
  }
};

/**
 * Except for the initialization phase, where an Abstraction object gets linked
 * with a specific Implementation object, the client code should only depend on
 * the Abstraction class. This way the client code can support any abstraction-
 * implementation combination.
 */
void ClientCode(const Abstraction& abstraction) {
  // ...
  std::cout << abstraction.Operation();
  // ...
}
/**
 * The client code should be able to work with any pre-configured abstraction-
 * implementation combination.
 */


// 桥接模式测试
TEST(BridgePatternTest, BasicFunctionality) {
  // 创建具体实现类的实例
  ConcreteImplementationA* implementationA = new ConcreteImplementationA();
  ConcreteImplementationB* implementationB = new ConcreteImplementationB();

  // 创建抽象类的实例，并与具体实现关联
  Abstraction* abstraction = new Abstraction(implementationA);
  ExtendedAbstraction* extendedAbstraction = new ExtendedAbstraction(implementationB);

  // 测试基本抽象类与实现A的组合
  std::string result1 = abstraction->Operation();
  EXPECT_TRUE(result1.find("Abstraction: Base operation with:") != std::string::npos);
  EXPECT_TRUE(result1.find("ConcreteImplementationA") != std::string::npos);

  // 测试扩展抽象类与实现B的组合
  std::string result2 = extendedAbstraction->Operation();
  EXPECT_TRUE(result2.find("ExtendedAbstraction: Extended operation with:") != std::string::npos);
  EXPECT_TRUE(result2.find("ConcreteImplementationB") != std::string::npos);

  // 测试客户端代码
  testing::internal::CaptureStdout();
  ClientCode(*abstraction);
  std::string output1 = testing::internal::GetCapturedStdout();
  EXPECT_TRUE(output1.find("Abstraction: Base operation with:") != std::string::npos);

  testing::internal::CaptureStdout();
  ClientCode(*extendedAbstraction);
  std::string output2 = testing::internal::GetCapturedStdout();
  EXPECT_TRUE(output2.find("ExtendedAbstraction: Extended operation with:") != std::string::npos);

  // 清理资源
  delete abstraction;
  delete extendedAbstraction;
  delete implementationA;
  delete implementationB;
}

// 测试不同抽象与实现的组合
TEST(BridgePatternTest, DifferentCombinations) {
  // 创建所有可能的组合并测试
  ConcreteImplementationA* implementationA = new ConcreteImplementationA();
  ConcreteImplementationB* implementationB = new ConcreteImplementationB();
  
  Abstraction* abstraction1 = new Abstraction(implementationA);
  Abstraction* abstraction2 = new Abstraction(implementationB);
  ExtendedAbstraction* extendedAbstraction1 = new ExtendedAbstraction(implementationA);
  ExtendedAbstraction* extendedAbstraction2 = new ExtendedAbstraction(implementationB);

  // 验证所有组合都能正常工作
  EXPECT_NO_THROW(abstraction1->Operation());
  EXPECT_NO_THROW(abstraction2->Operation());
  EXPECT_NO_THROW(extendedAbstraction1->Operation());
  EXPECT_NO_THROW(extendedAbstraction2->Operation());

  // 清理资源
  delete abstraction1;
  delete abstraction2;
  delete extendedAbstraction1;
  delete extendedAbstraction2;
  delete implementationA;
  delete implementationB;
}