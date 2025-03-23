#include <mutex>
#include <memory>

class DoubleCheckedSingleton {
private:
    static DoubleCheckedSingleton* instance;
    static std::mutex mtx;

    DoubleCheckedSingleton() {}

    DoubleCheckedSingleton(const DoubleCheckedSingleton&) = delete;
    DoubleCheckedSingleton& operator=(const DoubleCheckedSingleton&) = delete;

public:
    static DoubleCheckedSingleton* getInstance() {
        if (instance == nullptr) {
            std::lock_guard<std::mutex> lock(mtx);
            if (instance == nullptr) {
                instance = new DoubleCheckedSingleton();
            }
        }
        return instance;
    }
};