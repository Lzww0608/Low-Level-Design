
// Meyers Singleton

/*
 * Meyers Singleton 是一种线程安全的单例模式实现方式。
 * 
 * 特点：
 * 1. 利用C++11标准中的静态局部变量初始化保证线程安全
 * 2. 懒加载 - 只有在第一次调用getInstance()时才会创建实例
 * 3. 自动处理实例的销毁（程序结束时自动调用析构函数）
 * 4. 不需要使用锁或其他同步机制
 * 
 * 优势：
 * - 简洁易用
 * - 线程安全
 * - 避免了双检锁(DCLP)的复杂性
 * - 符合C++标准
 */

class MeyersSingleton {
private:
    MeyersSingleton() {}

    MeyersSingleton(const MeyersSingleton&) = delete;
    MeyersSingleton& operator=(const MeyersSingleton&) = delete;

public:

    static MeyersSingleton& getInstance() {
        static MeyersSingleton instance;
        return instance;
    }
    
    
};