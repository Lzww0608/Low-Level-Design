#include "Coffee.h"

Coffee::Coffee(CoffeeType type, std::string name, double price)
    : type(type), name(std::move(name)), price(price) {}

CoffeeType Coffee::getType() const {
    return type;
}

std::string Coffee::getName() const {
    return name;
}

double Coffee::getPrice() const {
    return price;
}
