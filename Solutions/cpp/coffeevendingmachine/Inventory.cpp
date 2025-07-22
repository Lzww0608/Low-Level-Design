#include "Inventory.h"

Inventory::Inventory() {
    ingredients[CoffeeType::ESPRESSO] = 100;
    ingredients[CoffeeType::CAPPUCCINO] = 100;
    ingredients[CoffeeType::LATTE] = 100;
    ingredients[CoffeeType::AMERICANO] = 100;
    ingredients[CoffeeType::MOCHA] = 100;
    ingredients[CoffeeType::LATTE_MACCHIATO] = 100;
    ingredients[CoffeeType::IRISH_COFFEE] = 100;
}


void Inventory::addIngredients(CoffeeType coffeeType, int quantity) {
    std::lock_guard<std::mutex> lock(ingredientsMutex);
    ingredients[coffeeType] += quantity;
}

void Inventory::deductIngredients(CoffeeType coffeeType) {
    std::lock_guard<std::mutex> lock(ingredientsMutex);
    ingredients[coffeeType]--;
}

bool Inventory::hasIngredients(CoffeeType coffeeType) {
    std::lock_guard<std::mutex> lock(ingredientsMutex);
    return ingredients[coffeeType] > 0;
}