#ifndef COFFEE_H
#define COFFEE_H

#include "CoffeType.h"
#include <string>

class Coffee {
public:
    Coffee(CoffeeType type, std::string name, double price);

    CoffeeType getType() const;
    std::string getName() const;
    double getPrice() const;

private:
    CoffeeType type;
    std::string name;
    double price;
};

#endif