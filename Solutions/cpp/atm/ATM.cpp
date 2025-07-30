#include "ATM.h"

ATM::ATM() : currentAccount(nullptr), isAuthenticated(false) {}

ATM::~ATM() {
    for (auto account : accounts) {
        delete account;
    }
}

void ATM::addAccount(Account* account) {
    accounts.push_back(account);
}

bool ATM::authenticate(const std::string& accountNumber, const std::string& pin) {
    for (auto account : accounts) {
        if (account->getAccountNumber() == accountNumber && account->validatePin(pin)) {
            currentAccount = account;
            isAuthenticated = true;
            return true;
        }   
    }
    return false;
}

void ATM::logout() {
    currentAccount = nullptr;
    isAuthenticated = false;
}

bool ATM::withdraw(double amount) {
    if (isAuthenticated && currentAccount) {
        return currentAccount->withdraw(amount);
    }
    return false;
}

bool ATM::deposit(double amount) {
    if (isAuthenticated && currentAccount) {
        return currentAccount->deposit(amount);     
    }
    return false;
}

void ATM::displayBalance() const {
    if (isAuthenticated && currentAccount) {
        currentAccount->display();
    }   
}

void ATM::displayMenu() const {
    std::cout << "ATM Menu" << std::endl;
    std::cout << "0. Authenticate" << std::endl;
    std::cout << "1. Withdraw" << std::endl;
    std::cout << "2. Deposit" << std::endl;
    std::cout << "3. Display Balance" << std::endl;
    std::cout << "4. Logout" << std::endl;
}

void ATM::start() {
    while (true) {
        displayMenu();
        int choice;
        std::cin >> choice;
        switch (choice) {
            case 0:
                std::string accountNumber, pin;
                std::cout << "Enter account number: ";
                std::cin >> accountNumber;
                std::cout << "Enter PIN: ";
                std::cin >> pin;
                if (authenticate(accountNumber, pin)) {
                    std::cout << "Authentication successful" << std::endl;
                } else {
                    std::cout << "Authentication failed" << std::endl;
                }
                break;
            case 1:
                double amount;
                std::cout << "Enter amount to withdraw: ";
                std::cin >> amount; 
                if (withdraw(amount)) {
                    std::cout << "Withdrawal successful" << std::endl;
                } else {
                    std::cout << "Withdrawal failed" << std::endl;
                }
                break;
            case 2:
                std::cout << "Enter amount to deposit: ";
                std::cin >> amount; 
                if (deposit(amount)) {
                    std::cout << "Deposit successful" << std::endl;
                } else {
                    std::cout << "Deposit failed" << std::endl;
                }
                break;
            case 3:
                displayBalance();   
                break;
            case 4:
                logout();
                return;
            default:
                std::cout << "Invalid choice" << std::endl;
        }
    }
}