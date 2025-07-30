#include "Account.h"

Account::Account(const std::string& accountNumer, const std::string& pin, double balance = 0.0): 
    accountNumber(accountNumer), pin(pin), balance(balance) {}

std::string Account::getAccountNumber() const {
    return accountNumber;
}

double Account::getBalance() const {
    return balance;
}

bool Account::validatePin(const std::string& inputPin) const {
    return pin == inputPin;
}

bool Account::deposit(double amount) {
    if (amount > 0) {
        balance += amount;
        return true;
    }
    return false;
}

bool Account::withdraw(double amount) {
    if (amount > 0 && amount <= balance) { 
        balance -= amount;
        return true;
    }
    return false;
}

void Account::display() const {
    std::cout << "Account Number: " << accountNumber << std::endl;
    std::cout << "Balance: " << balance << std::endl;
}