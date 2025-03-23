
# Creational Patterns

Focus: Object creation mechanisms.

Goal:
+ Separate the creation of objects from their usage to improve code flexibility and reusability.
+ Hide the creation details of objects, so client code only needs to care about how to use objects, without needing to worry about how they are created.
+ Control the object creation process, such as limiting the number of objects, delaying object creation, etc.

Common Patterns:
Singleton (Singleton): Ensure a class has only one instance and provide a global access point to it.

Factory Method (Factory Method): Define an interface for creating objects, but let subclasses decide which class to instantiate.

Abstract Factory (Abstract Factory): Provide an interface for creating families of related or dependent objects without specifying their concrete classes.

Builder (Builder): Separate the construction of a complex object from its representation, so that the same construction process can create different representations.

Prototype (Prototype): Create new objects by copying existing objects, without relying on their classes.

## Singleton
To implement the singleton pattern, we must prevent external objects from creating instances of the singleton class. Only the singleton class should be permitted to create its own objects.

This can be achieved by making the **constructor private** and providing a **static method** for external objects to access it.


## Factory
The Factory Method separates product construction code from the code that actually uses the product. Therefore it’s easier to extend the product construction code independently from the rest of the code.

## Abstract Factory
Use the Abstract Factory when your code needs to work with various families of related products, but you don’t want it to depend on the concrete classes of those products—they might be unknown beforehand or you simply want to allow for future extensibility.

## Builder
Say you have a constructor with ten optional parameters. Calling such a beast is very inconvenient; therefore, you overload the constructor and create several shorter versions with fewer parameters. These constructors still refer to the main one, passing some default values into any omitted parameters.
```c++
class Pizza {
    Pizza(int size) { ... }
    Pizza(int size, boolean cheese) { ... }
    Pizza(int size, boolean cheese, boolean pepperoni) { ... }
    // ...
}
```
The Builder pattern lets you build objects step by step, using only those steps that you really need. After implementing the pattern, you don’t have to cram dozens of parameters into your constructors anymore.
The Builder pattern lets you construct products step-by-step. You could defer execution of some steps without breaking the final product. You can even call steps recursively, which comes in handy when you need to build an object tree.