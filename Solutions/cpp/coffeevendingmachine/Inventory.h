#ifndef INVENTORY_H
#define INVENTORY_H

#include "CoffeType.h"
#include <map>
#include <mutex>

class Inventory {
public:
    Inventory();

    void addIngredients(CoffeeType coffeeType, int quantity);
    void deductIngredients(CoffeeType coffeeType);
    bool hasIngredients(CoffeeType coffeeType);

private:
    std::map<CoffeeType, int> ingredients;

    std::mutex ingredientsMutex;
};


#endif