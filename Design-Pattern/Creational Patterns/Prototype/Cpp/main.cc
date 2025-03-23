#include <iostream>
#include <string>
#include <unordered_map>
#include <gtest/gtest.h>

using std::string;

// Prototype Design Pattern
//
// Intent: Lets you copy existing objects without making your code dependent on
// their classes.

enum Type {
  PROTOTYPE_1 = 0,
  PROTOTYPE_2
};

/**
 * The example class that has cloning ability. We'll see how the values of field
 * with different types will be cloned.
 */

class Prototype {
 protected:
  string prototype_name_;
  float prototype_field_;

 public:
  Prototype() {}
  Prototype(string prototype_name)
      : prototype_name_(prototype_name) {
  }
  virtual ~Prototype() {}
  virtual Prototype *Clone() const = 0;
  virtual void Method(float prototype_field) {
    this->prototype_field_ = prototype_field;
    std::cout << "Call Method from " << prototype_name_ << " with field : " << prototype_field << std::endl;
  }
};

/**
 * ConcretePrototype1 is a Sub-Class of Prototype and implement the Clone Method
 * In this example all data members of Prototype Class are in the Stack. If you
 * have pointers in your properties for ex: String* name_ ,you will need to
 * implement the Copy-Constructor to make sure you have a deep copy from the
 * clone method
 */

class ConcretePrototype1 : public Prototype {
 private:
  float concrete_prototype_field1_;

 public:
  ConcretePrototype1(string prototype_name, float concrete_prototype_field)
      : Prototype(prototype_name), concrete_prototype_field1_(concrete_prototype_field) {
  }

  /**
   * Notice that Clone method return a Pointer to a new ConcretePrototype1
   * replica. so, the client (who call the clone method) has the responsability
   * to free that memory. If you have smart pointer knowledge you may prefer to
   * use unique_pointer here.
   */
  Prototype *Clone() const override {
    return new ConcretePrototype1(*this);
  }
};

class ConcretePrototype2 : public Prototype {
 private:
  float concrete_prototype_field2_;

 public:
  ConcretePrototype2(string prototype_name, float concrete_prototype_field)
      : Prototype(prototype_name), concrete_prototype_field2_(concrete_prototype_field) {
  }
  Prototype *Clone() const override {
    return new ConcretePrototype2(*this);
  }
};

/**
 * In PrototypeFactory you have two concrete prototypes, one for each concrete
 * prototype class, so each time you want to create a bullet , you can use the
 * existing ones and clone those.
 */

class PrototypeFactory {
 private:
  std::unordered_map<Type, Prototype *, std::hash<int>> prototypes_;

 public:
  PrototypeFactory() {
    prototypes_[Type::PROTOTYPE_1] = new ConcretePrototype1("PROTOTYPE_1 ", 50.f);
    prototypes_[Type::PROTOTYPE_2] = new ConcretePrototype2("PROTOTYPE_2 ", 60.f);
  }

  /**
   * Be carefull of free all memory allocated. Again, if you have smart pointers
   * knowelege will be better to use it here.
   */

  ~PrototypeFactory() {
    delete prototypes_[Type::PROTOTYPE_1];
    delete prototypes_[Type::PROTOTYPE_2];
  }

  /**
   * Notice here that you just need to specify the type of the prototype you
   * want and the method will create from the object with this type.
   */
  Prototype *CreatePrototype(Type type) {
    return prototypes_[type]->Clone();
  }
};

void Client(PrototypeFactory &prototype_factory) {
  std::cout << "Let's create a Prototype 1\n";

  Prototype *prototype = prototype_factory.CreatePrototype(Type::PROTOTYPE_1);
  prototype->Method(90);
  delete prototype;

  std::cout << "\n";

  std::cout << "Let's create a Prototype 2 \n";

  prototype = prototype_factory.CreatePrototype(Type::PROTOTYPE_2);
  prototype->Method(10);

  delete prototype;
}


// 测试 ConcretePrototype1 的克隆功能
TEST(PrototypeTest, ConcretePrototype1Clone) {
  ConcretePrototype1 prototype("Test1", 100.f);
  Prototype* clone = prototype.Clone();
  
  // 验证克隆对象是否正确
  testing::internal::CaptureStdout();
  clone->Method(50);
  std::string output = testing::internal::GetCapturedStdout();
  
  EXPECT_TRUE(output.find("Test1") != std::string::npos);
  EXPECT_TRUE(output.find("50") != std::string::npos);
  
  delete clone;
}

// 测试 ConcretePrototype2 的克隆功能
TEST(PrototypeTest, ConcretePrototype2Clone) {
  ConcretePrototype2 prototype("Test2", 200.f);
  Prototype* clone = prototype.Clone();
  
  // 验证克隆对象是否正确
  testing::internal::CaptureStdout();
  clone->Method(75);
  std::string output = testing::internal::GetCapturedStdout();
  
  EXPECT_TRUE(output.find("Test2") != std::string::npos);
  EXPECT_TRUE(output.find("75") != std::string::npos);
  
  delete clone;
}

// 测试 PrototypeFactory 的功能
TEST(PrototypeTest, PrototypeFactory) {
  PrototypeFactory factory;
  
  // 测试创建 PROTOTYPE_1
  testing::internal::CaptureStdout();
  Prototype* prototype1 = factory.CreatePrototype(Type::PROTOTYPE_1);
  prototype1->Method(25);
  std::string output1 = testing::internal::GetCapturedStdout();
  
  EXPECT_TRUE(output1.find("PROTOTYPE_1") != std::string::npos);
  EXPECT_TRUE(output1.find("25") != std::string::npos);
  
  // 测试创建 PROTOTYPE_2
  testing::internal::CaptureStdout();
  Prototype* prototype2 = factory.CreatePrototype(Type::PROTOTYPE_2);
  prototype2->Method(35);
  std::string output2 = testing::internal::GetCapturedStdout();
  
  EXPECT_TRUE(output2.find("PROTOTYPE_2") != std::string::npos);
  EXPECT_TRUE(output2.find("35") != std::string::npos);
  
  delete prototype1;
  delete prototype2;
}

// 主测试函数
int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
