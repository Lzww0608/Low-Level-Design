#ifndef ATM_H
#define ATM_H

#include "Account.h"
#include <vector>

class ATM {
public:
    ATM();
    ~ATM();

    void addAccount(Account* account);
    bool authenticate(const std::string& accountNumber, const std::string& pin);
    void logout();

    bool withdraw(double amount);
    bool deposit(double amount);
    void displayBalance() const;

    void displayMenu() const;
    void start();
private:
    std::vector<Account*> accounts;
    Account* currentAccount;
    bool isAuthenticated;
};

#endif
