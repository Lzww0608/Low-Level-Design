#include "CoffeeVendingMachine.h"
#include <iostream>
#include <algorithm>

CoffeeVendingMachine::CoffeeVendingMachine()
    : moneyCollected(0.0) {
    initializeMenu();
}

void CoffeeVendingMachine::initializeMenu() {
    coffeeMenu.clear();
    coffeeMenu.emplace_back(CoffeeType::ESPRESSO, "Espresso", 2.0);
    coffeeMenu.emplace_back(CoffeeType::CAPPUCCINO, "Cappuccino", 2.5);
    coffeeMenu.emplace_back(CoffeeType::LATTE, "Latte", 2.5);
    coffeeMenu.emplace_back(CoffeeType::AMERICANO, "Americano", 2.0);
    coffeeMenu.emplace_back(CoffeeType::MOCHA, "Mocha", 3.0);
    coffeeMenu.emplace_back(CoffeeType::LATTE_MACCHIATO, "Latte Macchiato", 3.0);
    coffeeMenu.emplace_back(CoffeeType::IRISH_COFFEE, "Irish Coffee", 3.5);
}

void CoffeeVendingMachine::displayMenu() const {
    std::cout << "Coffee Menu:\n";
    for (const auto& coffee : coffeeMenu) {
        std::cout << static_cast<int>(coffee.getType()) << ". "
                  << coffee.getName() << " - $" << coffee.getPrice() << "\n";
    }
}

bool CoffeeVendingMachine::selectCoffee(CoffeeType coffeeType, double payment) {
    // Find the coffee in the menu
    auto it = std::find_if(coffeeMenu.begin(), coffeeMenu.end(),
        [coffeeType](const Coffee& c) { return c.getType() == coffeeType; });

    if (it == coffeeMenu.end()) {
        std::cout << "Invalid coffee selection.\n";
        return false;
    }

    const Coffee& selectedCoffee = *it;

    if (payment < selectedCoffee.getPrice()) {
        std::cout << "Insufficient payment. Please insert at least $" << selectedCoffee.getPrice() << ".\n";
        return false;
    }

    if (!inventory.hasIngredients(coffeeType)) {
        std::cout << "Sorry, " << selectedCoffee.getName() << " is out of stock.\n";
        return false;
    }

    inventory.deductIngredients(coffeeType);
    moneyCollected += selectedCoffee.getPrice();

    std::cout << "Dispensing " << selectedCoffee.getName() << ". Enjoy your coffee!\n";
    if (payment > selectedCoffee.getPrice()) {
        std::cout << "Returning change: $" << (payment - selectedCoffee.getPrice()) << "\n";
    }
    return true;
}

void CoffeeVendingMachine::refillIngredients(CoffeeType coffeeType, int quantity) {
    inventory.addIngredients(coffeeType, quantity);
    std::cout << "Refilled " << quantity << " units for coffee type " << static_cast<int>(coffeeType) << ".\n";
}

double CoffeeVendingMachine::getMoneyCollected() const {
    return moneyCollected;
}

void CoffeeVendingMachine::displayInventory() {
    std::cout << "Inventory Status:\n";
    for (const auto& coffee : coffeeMenu) {
        std::cout << coffee.getName() << ": ";
        if (inventory.hasIngredients(coffee.getType())) {
            std::cout << "Available\n";
        } else {
            std::cout << "Out of stock\n";
        }
    }
}
