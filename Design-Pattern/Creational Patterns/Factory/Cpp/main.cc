#include <string>
#include <gtest/gtest.h>
/** 
* The Product interface declares the operations that all concrete products must implement.
*/

class Product {
public:
    virtual ~Product() {}
    virtual std::string Operation() const = 0;
};

/**
* Concrete Products provide various implementations of the Product interface.
 */

class ConcreteProduct1: public Product {
public:
    std::string Operation() const override {
        return "{Result of the ConcreteProduct1}";
    }
};

class ConcreteProduct2 : public Product {
public:
    std::string Operation() const override {
        return "{Result of the ConcreteProduct2}";
    }
};

/**
* The Creator class declares the factory method that is supposed to return an object of a Product class.
* The Creator's subclasses usually provide the implementation of this method.
*/

class Creator {
public:
    virtual ~Creator() {}
    virtual Product* FactoryMethod() const = 0;

    std::string SomeOperation() const {
        // Call the factory method to create a Product object.
        Product* product = this->FactoryMethod();
        // Now, use the product.
        std::string result = "Creator: The same creator's code has just worked with " + product->Operation();
        delete product;
        return result;
    }
};


class ConcreteCreator1: public Creator {
public:
    Product* FactoryMethod() const override {
        return new ConcreteProduct1();
    }
};

class ConcreteCreator2: public Creator {
public:
    Product* FactoryMethod() const override {
        return new ConcreteProduct2();
    }
};


void ClientCode(const Creator& creator) {
  // ...
  std::cout << "Client: I'm not aware of the creator's class, but it still works.\n"
            << creator.SomeOperation() << std::endl;
  // ...
}



// 测试 ConcreteCreator1 是否正确创建 ConcreteProduct1
TEST(FactoryMethodTest, TestConcreteCreator1) {
    Creator* creator = new ConcreteCreator1();
    std::string result = creator->SomeOperation();
    
    // 验证结果中包含 ConcreteProduct1 的输出
    EXPECT_TRUE(result.find("Result of the ConcreteProduct1") != std::string::npos);
    
    delete creator;
}

// 测试 ConcreteCreator2 是否正确创建 ConcreteProduct2
TEST(FactoryMethodTest, TestConcreteCreator2) {
    Creator* creator = new ConcreteCreator2();
    std::string result = creator->SomeOperation();
    
    // 验证结果中包含 ConcreteProduct2 的输出
    EXPECT_TRUE(result.find("Result of the ConcreteProduct2") != std::string::npos);
    
    delete creator;
}

// 测试 ClientCode 函数是否正常工作
TEST(FactoryMethodTest, TestClientCode) {
    // 重定向 cout 到 stringstream 以捕获输出
    std::stringstream buffer;
    std::streambuf* old = std::cout.rdbuf(buffer.rdbuf());
    
    // 使用 ConcreteCreator1 测试 ClientCode
    Creator* creator1 = new ConcreteCreator1();
    ClientCode(*creator1);
    
    // 恢复 cout
    std::cout.rdbuf(old);
    
    // 验证输出包含预期内容
    std::string output = buffer.str();
    EXPECT_TRUE(output.find("Client: I'm not aware of the creator's class") != std::string::npos);
    EXPECT_TRUE(output.find("Result of the ConcreteProduct1") != std::string::npos);
    
    delete creator1;
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}

