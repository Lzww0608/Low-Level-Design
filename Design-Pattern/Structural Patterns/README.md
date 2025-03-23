# Structural Patterns

Focus: Combination of classes and objects.

Goal:
+ Combine classes or objects into larger structures to simplify system design and improve code maintainability.
+ Achieve new functionalities or change system behavior through different combination methods.
+ Extend system functionality without changing existing code.

Common Patterns:
Adapter: Convert the interface of a class into another interface that clients expect, allowing classes that are originally incompatible to work together.

Bridge: Decouple the abstraction from its implementation so that the two can vary independently.

Composite: Compose objects into tree structures to represent part-whole hierarchies, allowing clients to treat individual objects and composite objects uniformly.

Decorator: Dynamically add additional responsibilities to an object without modifying its original class.

Facade: Provide a unified interface to a set of interfaces in a subsystem, making the subsystem easier to use.

Flyweight: Use sharing to support large numbers of fine-grained objects efficiently.

Proxy: Provide a surrogate or placeholder for another object to control access to it.


## Adapter
The Adapter pattern lets you create a middle-layer class that serves as a translator between your code and a legacy class, a 3rd-party class or any other class with a weird interface.

You could extend each subclass and put the missing functionality into new child classes. However, you’ll need to duplicate the code across all of these new classes, which smells really bad.