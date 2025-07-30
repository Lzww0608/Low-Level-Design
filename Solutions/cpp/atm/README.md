# Designing an ATM System
## Requirements
+ The ATM system should support basic operations such as balance inquiry, cash withdrawal, and cash deposit.
+ Users should be able to authenticate themselves using a card and a PIN (Personal Identification Number).
+ The system should interact with a bank's backend system to validate user accounts and perform transactions.
+ The ATM should have a cash dispenser to dispense cash to users.
+ The system should handle concurrent access and ensure data consistency.
+ The ATM should have a user-friendly interface for users to interact with.

## Sulotions

### Class —— Account

+------------------------------------------------------------+  
|      Account                                               | // 类名  
+------------------------------------------------------------+  
| - accountNumber: std::string                               | // private成员变量  
| - pin: std::string                                         |  
| - balance: double                                          |  
+------------------------------------------------------------+  
| + Account(accountNumber: std::string, pin: std::string,    | // public构造函数 
|  balance: double)                                          |  
| + getAccountNumber() const : std::string                   |  
| + validatePin(const std::string&): bool                    |  
| + withdraw(double) bool                                    |  
| + deposit(double) bool                                     |
| + getBalance() const : double                              |
| + displayBalance() const                                   |
+------------------------------------------------------------+  


### Class —— ATM

