#include <iostream>
#include <string>   
#include <gtest/gtest.h>

/**
 * The base Component interface defines operations that can be altered by
 * decorators.
 */
class Component {
 public:
  virtual ~Component() {}
  virtual std::string Operation() const = 0;
};
/**
 * Concrete Components provide default implementations of the operations. There
 * might be several variations of these classes.
 */
class ConcreteComponent : public Component {
 public:
  std::string Operation() const override {
    return "ConcreteComponent";
  }
};
/**
 * The base Decorator class follows the same interface as the other components.
 * The primary purpose of this class is to define the wrapping interface for all
 * concrete decorators. The default implementation of the wrapping code might
 * include a field for storing a wrapped component and the means to initialize
 * it.
 */
class Decorator : public Component {
  /**
   * @var Component
   */
 protected:
  Component* component_;

 public:
  Decorator(Component* component) : component_(component) {
  }
  /**
   * The Decorator delegates all work to the wrapped component.
   */
  std::string Operation() const override {
    return this->component_->Operation();
  }
};
/**
 * Concrete Decorators call the wrapped object and alter its result in some way.
 */
class ConcreteDecoratorA : public Decorator {
  /**
   * Decorators may call parent implementation of the operation, instead of
   * calling the wrapped object directly. This approach simplifies extension of
   * decorator classes.
   */
 public:
  ConcreteDecoratorA(Component* component) : Decorator(component) {
  }
  std::string Operation() const override {
    return "ConcreteDecoratorA(" + Decorator::Operation() + ")";
  }
};
/**
 * Decorators can execute their behavior either before or after the call to a
 * wrapped object.
 */
class ConcreteDecoratorB : public Decorator {
 public:
  ConcreteDecoratorB(Component* component) : Decorator(component) {
  }

  std::string Operation() const override {
    return "ConcreteDecoratorB(" + Decorator::Operation() + ")";
  }
};
/**
 * The client code works with all objects using the Component interface. This
 * way it can stay independent of the concrete classes of components it works
 * with.
 */
void ClientCode(Component* component) {
  // ...
  std::cout << "RESULT: " << component->Operation();
  // ...
}


// 装饰器模式测试
TEST(DecoratorTest, BasicFunctionality) {
  Component* simple = new ConcreteComponent();
  EXPECT_EQ("ConcreteComponent", simple->Operation());
  
  Component* decorator1 = new ConcreteDecoratorA(simple);
  EXPECT_EQ("ConcreteDecoratorA(ConcreteComponent)", decorator1->Operation());
  
  Component* decorator2 = new ConcreteDecoratorB(decorator1);
  EXPECT_EQ("ConcreteDecoratorB(ConcreteDecoratorA(ConcreteComponent))", decorator2->Operation());
  
  // 清理资源
  delete decorator2;
  delete decorator1;
  delete simple;
}

// 测试嵌套装饰器
TEST(DecoratorTest, NestedDecorators) {
  Component* simple = new ConcreteComponent();
  
  // 创建多层装饰
  Component* decorator1 = new ConcreteDecoratorA(simple);
  Component* decorator2 = new ConcreteDecoratorB(decorator1);
  Component* decorator3 = new ConcreteDecoratorA(decorator2);
  
  // 验证嵌套装饰器的输出
  EXPECT_EQ("ConcreteDecoratorA(ConcreteDecoratorB(ConcreteDecoratorA(ConcreteComponent)))", decorator3->Operation());
  
  // 清理资源
  delete decorator3;
  delete decorator2;
  delete decorator1;
  delete simple;
}

// 主函数运行测试
int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
