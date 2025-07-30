#ifndef ACCOUNT_H
#define ACCOUNT_H

#include <string>

class Account {
public:
    Account(const std::string& accountNumer, const std::string& pin, double balance = 0.0);
    ~Account() = default;
    std::string getAccountNumber() const;
    std::string getBalance() const;
    bool validatePin(const std::string& inputPin) const;

    bool deposit(double amount);
    bool withdraw(double amount);
    void display() const;
private:
    std::string accountNumber;
    std::string pin;
    double balance;
};

#endif
