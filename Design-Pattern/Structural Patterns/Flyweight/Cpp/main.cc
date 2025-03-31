#include <iostream>
#include <unordered_map>
#include <string>
#include <gtest/gtest.h>

/**
 * Flyweight Design Pattern
 *
 * Intent: Lets you fit more objects into the available amount of RAM by sharing
 * common parts of state between multiple objects, instead of keeping all of the
 * data in each object.
 */

struct SharedState
{
    std::string brand_;
    std::string model_;
    std::string color_;

    SharedState(const std::string &brand, const std::string &model, const std::string &color)
        : brand_(brand), model_(model), color_(color)
    {
    }

    friend std::ostream &operator<<(std::ostream &os, const SharedState &ss)
    {
        return os << "[ " << ss.brand_ << " , " << ss.model_ << " , " << ss.color_ << " ]";
    }
};

struct UniqueState
{
    std::string owner_;
    std::string plates_;

    UniqueState(const std::string &owner, const std::string &plates)
        : owner_(owner), plates_(plates)
    {
    }

    friend std::ostream &operator<<(std::ostream &os, const UniqueState &us)
    {
        return os << "[ " << us.owner_ << " , " << us.plates_ << " ]";
    }
};

/**
 * The Flyweight stores a common portion of the state (also called intrinsic
 * state) that belongs to multiple real business entities. The Flyweight accepts
 * the rest of the state (extrinsic state, unique for each entity) via its
 * method parameters.
 */
class Flyweight
{
private:
    SharedState *shared_state_;

public:
    Flyweight(const SharedState *shared_state) : shared_state_(new SharedState(*shared_state))
    {
    }
    Flyweight(const Flyweight &other) : shared_state_(new SharedState(*other.shared_state_))
    {
    }
    ~Flyweight()
    {
        delete shared_state_;
    }
    SharedState *shared_state() const
    {
        return shared_state_;
    }
    void Operation(const UniqueState &unique_state) const
    {
        std::cout << "Flyweight: Displaying shared (" << *shared_state_ << ") and unique (" << unique_state << ") state.\n";
    }
};
/**
 * The Flyweight Factory creates and manages the Flyweight objects. It ensures
 * that flyweights are shared correctly. When the client requests a flyweight,
 * the factory either returns an existing instance or creates a new one, if it
 * doesn't exist yet.
 */
class FlyweightFactory
{
    /**
     * @var Flyweight[]
     */
private:
    std::unordered_map<std::string, Flyweight> flyweights_;
    /**
     * Returns a Flyweight's string hash for a given state.
     */
    std::string GetKey(const SharedState &ss) const
    {
        return ss.brand_ + "_" + ss.model_ + "_" + ss.color_;
    }

public:
    FlyweightFactory(std::initializer_list<SharedState> share_states)
    {
        for (const SharedState &ss : share_states)
        {
            this->flyweights_.insert(std::make_pair<std::string, Flyweight>(this->GetKey(ss), Flyweight(&ss)));
        }
    }

    /**
     * Returns an existing Flyweight with a given state or creates a new one.
     */
    Flyweight GetFlyweight(const SharedState &shared_state)
    {
        std::string key = this->GetKey(shared_state);
        if (this->flyweights_.find(key) == this->flyweights_.end())
        {
            std::cout << "FlyweightFactory: Can't find a flyweight, creating new one.\n";
            this->flyweights_.insert(std::make_pair(key, Flyweight(&shared_state)));
        }
        else
        {
            std::cout << "FlyweightFactory: Reusing existing flyweight.\n";
        }
        return this->flyweights_.at(key);
    }
    void ListFlyweights() const
    {
        size_t count = this->flyweights_.size();
        std::cout << "\nFlyweightFactory: I have " << count << " flyweights:\n";
        for (std::pair<std::string, Flyweight> pair : this->flyweights_)
        {
            std::cout << pair.first << "\n";
        }
    }
};

// ...
void AddCarToPoliceDatabase(
    FlyweightFactory &ff, const std::string &plates, const std::string &owner,
    const std::string &brand, const std::string &model, const std::string &color)
{
    std::cout << "\nClient: Adding a car to database.\n";
    const Flyweight &flyweight = ff.GetFlyweight({brand, model, color});
    // The client code either stores or calculates extrinsic state and passes it
    // to the flyweight's methods.
    flyweight.Operation({owner, plates});
}

/**
 * The client code usually creates a bunch of pre-populated flyweights in the
 * initialization stage of the application.
 */

int main()
{
    FlyweightFactory *factory = new FlyweightFactory({{"Chevrolet", "Camaro2018", "pink"}, {"Mercedes Benz", "C300", "black"}, {"Mercedes Benz", "C500", "red"}, {"BMW", "M5", "red"}, {"BMW", "X6", "white"}});
    factory->ListFlyweights();

    AddCarToPoliceDatabase(*factory,
                            "CL234IR",
                            "James Doe",
                            "BMW",
                            "M5",
                            "red");

    AddCarToPoliceDatabase(*factory,
                            "CL234IR",
                            "James Doe",
                            "BMW",
                            "X1",
                            "red");
    factory->ListFlyweights();
    delete factory;

    return 0;
}


// 测试Flyweight模式
TEST(FlyweightTest, BasicFunctionality) {
    // 创建工厂并预先填充一些享元对象
    FlyweightFactory factory({
        {"BMW", "M5", "red"}, 
        {"Mercedes Benz", "C300", "black"}
    });
    
    // 测试获取已存在的享元对象
    const Flyweight &flyweight1 = factory.GetFlyweight({"BMW", "M5", "red"});
    
    // 测试获取新的享元对象
    const Flyweight &flyweight2 = factory.GetFlyweight({"BMW", "X6", "white"});
    
    // 测试享元对象数量
    std::stringstream buffer;
    std::streambuf* old = std::cout.rdbuf(buffer.rdbuf());
    factory.ListFlyweights();
    std::cout.rdbuf(old);
    
    // 验证输出中包含3个享元对象
    std::string output = buffer.str();
    EXPECT_TRUE(output.find("I have 3 flyweights") != std::string::npos);
    EXPECT_TRUE(output.find("BMW_M5_red") != std::string::npos);
    EXPECT_TRUE(output.find("Mercedes Benz_C300_black") != std::string::npos);
    EXPECT_TRUE(output.find("BMW_X6_white") != std::string::npos);
}

// 测试共享状态和唯一状态
TEST(FlyweightTest, StateHandling) {
    // 创建共享状态
    SharedState sharedState("BMW", "M5", "red");
    
    // 创建享元对象
    Flyweight flyweight(&sharedState);
    
    // 创建两个不同的唯一状态
    UniqueState uniqueState1("James Doe", "CL234IR");
    UniqueState uniqueState2("Jane Smith", "CL567IR");
    
    // 捕获操作输出以验证
    std::stringstream buffer1;
    std::streambuf* old1 = std::cout.rdbuf(buffer1.rdbuf());
    flyweight.Operation(uniqueState1);
    std::cout.rdbuf(old1);
    
    std::stringstream buffer2;
    std::streambuf* old2 = std::cout.rdbuf(buffer2.rdbuf());
    flyweight.Operation(uniqueState2);
    std::cout.rdbuf(old2);
    
    // 验证输出包含正确的共享状态和唯一状态
    std::string output1 = buffer1.str();
    std::string output2 = buffer2.str();
    
    EXPECT_TRUE(output1.find("BMW") != std::string::npos);
    EXPECT_TRUE(output1.find("M5") != std::string::npos);
    EXPECT_TRUE(output1.find("red") != std::string::npos);
    EXPECT_TRUE(output1.find("James Doe") != std::string::npos);
    EXPECT_TRUE(output1.find("CL234IR") != std::string::npos);
    
    EXPECT_TRUE(output2.find("BMW") != std::string::npos);
    EXPECT_TRUE(output2.find("M5") != std::string::npos);
    EXPECT_TRUE(output2.find("red") != std::string::npos);
    EXPECT_TRUE(output2.find("Jane Smith") != std::string::npos);
    EXPECT_TRUE(output2.find("CL567IR") != std::string::npos);
}
