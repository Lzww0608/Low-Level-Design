#include <mutex>
#include <memory>
#include <gtest/gtest.h>
#include <thread>

// Call Once Singleton

/*
 * Call Once Singleton 是一种使用 std::call_once 和 std::once_flag 实现的线程安全单例模式。
 * 
 * 特点：
 * 1. 使用 std::call_once 确保初始化函数只被调用一次，即使在多线程环境下
 * 2. 懒加载 - 只有在第一次调用 getInstance() 时才会创建实例
 * 3. 使用 std::unique_ptr 自动管理内存，避免内存泄漏
 * 4. 比双检锁模式更简洁且更安全
 * 
 * 优势：
 * - 线程安全保证
 * - 代码简洁明了
 * - 符合 C++11 标准
 * - 避免了双检锁可能存在的问题
 */


class CallOnceSingleton {
private:
    static std::unique_ptr<CallOnceSingleton> instance;
    static std::once_flag initInstanceFlag;

    CallOnceSingleton() {}

    CallOnceSingleton(const CallOnceSingleton&) = delete;
    CallOnceSingleton& operator=(const CallOnceSingleton&) = delete;
    
public:
    static CallOnceSingleton& getInstance() {
        std::call_once(initInstanceFlag, []() {
            instance.reset(new CallOnceSingleton());
        });

        return *instance;
    }

};


std::unique_ptr<CallOnceSingleton> CallOnceSingleton::instance;
std::once_flag CallOnceSingleton::initInstanceFlag;

// 测试 CallOnceSingleton 的单例特性
TEST(CallOnceSingletonTest, SingletonInstance) {
    // 获取两个实例引用
    CallOnceSingleton& instance1 = CallOnceSingleton::getInstance();
    CallOnceSingleton& instance2 = CallOnceSingleton::getInstance();
    
    // 验证两个引用指向同一个对象
    EXPECT_EQ(&instance1, &instance2);
}

// 测试多线程环境下的单例特性
TEST(CallOnceSingletonTest, ThreadSafety) {
    CallOnceSingleton* instances[10];
    std::thread threads[10];
    
    // 在多个线程中同时获取单例实例
    for (int i = 0; i < 10; ++i) {
        threads[i] = std::thread([&instances, i]() {
            instances[i] = &CallOnceSingleton::getInstance();
        });
    }
    
    // 等待所有线程完成
    for (auto& t : threads) {
        t.join();
    }
    
    // 验证所有线程获取的都是同一个实例
    for (int i = 1; i < 10; ++i) {
        EXPECT_EQ(instances[0], instances[i]);
    }
}
