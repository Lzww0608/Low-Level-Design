
// Not thread safe

// Checks if an instance already exists (instance == null).

// If not, it creates a new instance.

// If an instance already exists, it skips the creation step.
class LazySingleton {
private:
    static LazySingleton* instance;

    LazySingleton() {}
    LazySingleton(const LazySingleton&) = delete;
    LazySingleton& operator=(const LazySingleton&) = delete;

public:
    static LazySingleton* getInstance() {
        if (instance == nullptr) {
            instance = new LazySingleton();
        }
        return instance;
    }
}