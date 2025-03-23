#include <iostream>
#include <string>
#include <algorithm>
#include <memory>
#include <gtest/gtest.h>

/**
 * The Target defines the domain-specific interface used by the client code.
 */
class Target {
 public:
  virtual ~Target() = default;

  virtual std::string Request() const {
    return "Target: The default target's behavior.";
  }
};

/**
 * The Adaptee contains some useful behavior, but its interface is incompatible
 * with the existing client code. The Adaptee needs some adaptation before the
 * client code can use it.
 */
class Adaptee {
 public:
  std::string SpecificRequest() const {
    return ".eetpadA eht fo roivaheb laicepS";
  }
};

/**
 * The Adapter makes the Adaptee's interface compatible with the Target's
 * interface.
 */
class Adapter : public Target {
 private:
  Adaptee *adaptee_;

 public:
  Adapter(Adaptee *adaptee) : adaptee_(adaptee) {}
  std::string Request() const override {
    std::string to_reverse = this->adaptee_->SpecificRequest();
    std::reverse(to_reverse.begin(), to_reverse.end());
    return "Adapter: (TRANSLATED) " + to_reverse;
  }
};

/**
 * The client code supports all classes that follow the Target interface.
 */
void ClientCode(const Target *target) {
  std::cout << target->Request();
}

TEST(AdapterTest, TestAdapterPattern) {
  std::cout << "Client: I can work just fine with the Target objects:\n";
  Target *target = new Target;
  std::string result1 = target->Request();
  EXPECT_EQ(result1, "Target: The default target's behavior.");
  delete target;

  std::cout << "\n\n";
  
  Adaptee *adaptee = new Adaptee;
  std::cout << "Client: The Adaptee class has a weird interface. "
            << "See, I don't understand it:\n";
  std::string result2 = adaptee->SpecificRequest();
  EXPECT_EQ(result2, ".eetpadA eht fo roivaheb laicepS");
  
  std::cout << "\n\n";
  
  std::cout << "Client: But I can work with it via the Adapter:\n";
  Adapter *adapter = new Adapter(adaptee);
  std::string result3 = adapter->Request();
  EXPECT_EQ(result3, "Adapter: (TRANSLATED) Special behavior of the Adaptee.");
  
  delete adapter;
  delete adaptee;
}
