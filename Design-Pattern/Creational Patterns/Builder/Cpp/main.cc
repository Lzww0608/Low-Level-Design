/**
 * It makes sense to use the Builder pattern only when your products are quite
 * complex and require extensive configuration.
 *
 * Unlike in other creational patterns, different concrete builders can produce
 * unrelated products. In other words, results of various builders may not
 * always follow the same interface.
 */

#include <gtest/gtest.h>
#include <vector>
#include <string>
#include <iostream>
#include <sstream>

class Product1{
    public:
    std::vector<std::string> parts_;
    void ListParts()const{
        std::cout << "Product parts: ";
        for (size_t i=0;i<parts_.size();i++){
            if(parts_[i]== parts_.back()){
                std::cout << parts_[i];
            }else{
                std::cout << parts_[i] << ", ";
            }
        }
        std::cout << "\n\n"; 
    }
};


/**
 * The Builder interface specifies methods for creating the different parts of
 * the Product objects.
 */
class Builder{
    public:
    virtual ~Builder(){}
    virtual void ProducePartA() const =0;
    virtual void ProducePartB() const =0;
    virtual void ProducePartC() const =0;
};
/**
 * The Concrete Builder classes follow the Builder interface and provide
 * specific implementations of the building steps. Your program may have several
 * variations of Builders, implemented differently.
 */
class ConcreteBuilder1 : public Builder{
    private:

    Product1* product;

    /**
     * A fresh builder instance should contain a blank product object, which is
     * used in further assembly.
     */
    public:

    ConcreteBuilder1(){
        this->Reset();
    }

    ~ConcreteBuilder1(){
        delete product;
    }

    void Reset(){
        this->product= new Product1();
    }
    /**
     * All production steps work with the same product instance.
     */

    void ProducePartA()const override{
        this->product->parts_.push_back("PartA1");
    }

    void ProducePartB()const override{
        this->product->parts_.push_back("PartB1");
    }

    void ProducePartC()const override{
        this->product->parts_.push_back("PartC1");
    }

    /**
     * Concrete Builders are supposed to provide their own methods for
     * retrieving results. That's because various types of builders may create
     * entirely different products that don't follow the same interface.
     * Therefore, such methods cannot be declared in the base Builder interface
     * (at least in a statically typed programming language). Note that PHP is a
     * dynamically typed language and this method CAN be in the base interface.
     * However, we won't declare it there for the sake of clarity.
     *
     * Usually, after returning the end result to the client, a builder instance
     * is expected to be ready to start producing another product. That's why
     * it's a usual practice to call the reset method at the end of the
     * `getProduct` method body. However, this behavior is not mandatory, and
     * you can make your builders wait for an explicit reset call from the
     * client code before disposing of the previous result.
     */

    /**
     * Please be careful here with the memory ownership. Once you call
     * GetProduct the user of this function is responsable to release this
     * memory. Here could be a better option to use smart pointers to avoid
     * memory leaks
     */

    Product1* GetProduct() {
        Product1* result= this->product;
        this->Reset();
        return result;
    }
};

/**
 * The Director is only responsible for executing the building steps in a
 * particular sequence. It is helpful when producing products according to a
 * specific order or configuration. Strictly speaking, the Director class is
 * optional, since the client can control builders directly.
 */
class Director{
    /**
     * @var Builder
     */
    private:
    Builder* builder;
    /**
     * The Director works with any builder instance that the client code passes
     * to it. This way, the client code may alter the final type of the newly
     * assembled product.
     */

    public:

    void set_builder(Builder* builder){
        this->builder=builder;
    }

    /**
     * The Director can construct several product variations using the same
     * building steps.
     */

    void BuildMinimalViableProduct(){
        this->builder->ProducePartA();
    }
    
    void BuildFullFeaturedProduct(){
        this->builder->ProducePartA();
        this->builder->ProducePartB();
        this->builder->ProducePartC();
    }
};
/**
 * The client code creates a builder object, passes it to the director and then
 * initiates the construction process. The end result is retrieved from the
 * builder object.
 */
/**
 * I used raw pointers for simplicity however you may prefer to use smart
 * pointers here
 */
void ClientCode(Director& director)
{
    ConcreteBuilder1* builder = new ConcreteBuilder1();
    director.set_builder(builder);
    std::cout << "Standard basic product:\n"; 
    director.BuildMinimalViableProduct();
    
    Product1* p= builder->GetProduct();
    p->ListParts();
    delete p;

    std::cout << "Standard full featured product:\n"; 
    director.BuildFullFeaturedProduct();

    p= builder->GetProduct();
    p->ListParts();
    delete p;

    // Remember, the Builder pattern can be used without a Director class.
    std::cout << "Custom product:\n";
    builder->ProducePartA();
    builder->ProducePartC();
    p=builder->GetProduct();
    p->ListParts();
    delete p;

    delete builder;
}

// 测试 ConcreteBuilder1 的基本功能
TEST(BuilderPatternTest, TestConcreteBuilder1) {
    ConcreteBuilder1 builder;
    
    // 测试 ProducePartA
    builder.ProducePartA();
    Product1* product = builder.GetProduct();
    ASSERT_EQ(product->parts_.size(), 1);
    EXPECT_EQ(product->parts_[0], "PartA1");
    delete product;
    
    // 测试 ProducePartB
    builder.ProducePartB();
    product = builder.GetProduct();
    ASSERT_EQ(product->parts_.size(), 1);
    EXPECT_EQ(product->parts_[0], "PartB1");
    delete product;
    
    // 测试 ProducePartC
    builder.ProducePartC();
    product = builder.GetProduct();
    ASSERT_EQ(product->parts_.size(), 1);
    EXPECT_EQ(product->parts_[0], "PartC1");
    delete product;
    
    // 测试组合多个部件
    builder.ProducePartA();
    builder.ProducePartB();
    builder.ProducePartC();
    product = builder.GetProduct();
    ASSERT_EQ(product->parts_.size(), 3);
    EXPECT_EQ(product->parts_[0], "PartA1");
    EXPECT_EQ(product->parts_[1], "PartB1");
    EXPECT_EQ(product->parts_[2], "PartC1");
    delete product;
}

// 测试 Director 类的功能
TEST(BuilderPatternTest, TestDirector) {
    Director director;
    ConcreteBuilder1* builder = new ConcreteBuilder1();
    director.set_builder(builder);
    
    // 测试 BuildMinimalViableProduct
    director.BuildMinimalViableProduct();
    Product1* product = builder->GetProduct();
    ASSERT_EQ(product->parts_.size(), 1);
    EXPECT_EQ(product->parts_[0], "PartA1");
    delete product;
    
    // 测试 BuildFullFeaturedProduct
    director.BuildFullFeaturedProduct();
    product = builder->GetProduct();
    ASSERT_EQ(product->parts_.size(), 3);
    EXPECT_EQ(product->parts_[0], "PartA1");
    EXPECT_EQ(product->parts_[1], "PartB1");
    EXPECT_EQ(product->parts_[2], "PartC1");
    delete product;
    
    delete builder;
}

// 测试客户端代码
TEST(BuilderPatternTest, TestClientCode) {
    // 重定向 cout 输出进行测试
    std::stringstream buffer;
    std::streambuf* old = std::cout.rdbuf(buffer.rdbuf());
    
    Director director;
    ClientCode(director);
    
    // 验证输出包含预期的字符串
    std::string output = buffer.str();
    EXPECT_TRUE(output.find("Standard basic product") != std::string::npos);
    EXPECT_TRUE(output.find("Standard full featured product") != std::string::npos);
    EXPECT_TRUE(output.find("Custom product") != std::string::npos);
    EXPECT_TRUE(output.find("Product parts: PartA1") != std::string::npos);
    EXPECT_TRUE(output.find("Product parts: PartA1, PartB1, PartC1") != std::string::npos);
    EXPECT_TRUE(output.find("Product parts: PartA1, PartC1") != std::string::npos);
    
    // 恢复 cout
    std::cout.rdbuf(old);
}

int main(int argc, char **argv){
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}