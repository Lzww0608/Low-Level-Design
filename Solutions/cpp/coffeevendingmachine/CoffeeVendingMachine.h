#ifndef COFFEEVENDINGMACHINE_H
#define COFFEEVENDINGMACHINE_H

#include "Coffee.h"
#include "Inventory.h"

#include <vector>

class CoffeeVendingMachine {
private:
    std::vector<Coffee> coffeeMenu;
    Inventory inventory;
    double moneyCollected;

public:
    CoffeeVendingMachine();

    void initializeMenu();
    void displayMenu() const;
    bool selectCoffee(CoffeeType coffeeType, double payment);
    void refillIngredients(CoffeeType coffeeType, int quantity);
    double getMoneyCollected() const;
    void displayInventory();
};

#endif