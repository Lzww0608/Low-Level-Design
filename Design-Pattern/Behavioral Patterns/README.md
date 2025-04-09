# Behavioral Patterns

Focus: Communication and responsibility assignment between objects.

Goal:
+ Define the interaction methods between objects to achieve a more flexible and maintainable system.
+ Break down complex behaviors into a series of smaller, reusable objects to improve code readability and maintainability.
+ Implement different system behaviors through different interaction methods.
Common Patterns:

Chain of Responsibility: Pass the request along a chain of handlers until one of the handlers handles it.

Command: Encapsulate a request as an object, thereby parameterizing clients with different requests.

Interpreter: Given a language, define a representation of its grammar, and define an interpreter that uses this representation to interpret sentences in the language.

Iterator: Provide a way to sequentially access elements in an aggregate object without exposing its internal representation.

Mediator: Use a mediator object to encapsulate a series of object interactions.

Memento: Without violating encapsulation, capture an object's internal state and save this state outside the object.

Observer: Define a one-to-many dependency relationship between objects, such that when the state of one object changes, all objects that depend on it are notified and automatically updated.

State: Allows an object to change its behavior when its internal state changes.

Strategy: Define a series of algorithms, encapsulate them individually, and make them interchangeable.
Template Method: Define the skeleton of an algorithm in an operation, while deferring some steps to subclasses.

Visitor: Represents an operation to be performed on the elements of an object structure


# Iterator
Iterator is a behavioral design pattern that lets you traverse elements of a collection without exposing its underlying representation (list, stack, tree, etc.).

# Observer
Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the object they’re observing.


# Strategy
Strategy is a behavioral design pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.

# Command 
Command is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request. This transformation lets you pass requests as a method arguments, delay or queue a request’s execution, and support undoable operations.

# State
State is a behavioral design pattern that lets an object alter its behavior when its internal state changes. It appears as if the object changed its class.


# Template Method
Template Method is a behavioral design pattern that defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of the algorithm without changing its structure.

# Visitor
Visitor is a behavioral design pattern that lets you separate algorithms from the objects on which they operate.

# Mediator
Mediator is a behavioral design pattern that lets you reduce chaotic dependencies between objects. The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object.