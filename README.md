# Low-Level Design (LLD) Learning Repository

A comprehensive collection of design patterns, system design problems, and their implementations in multiple programming languages (C++ and Go). This repository serves as both a learning resource and a reference guide for software engineers preparing for system design interviews or looking to improve their object-oriented design skills.

## 📚 Repository Structure

```
LLD/
├── Design-Pattern/          # Design pattern implementations and examples
│   ├── Creational Patterns/
│   ├── Behavioral Patterns/
│   └── Structural Patterns/
├── Problems/                # LLD problem statements and requirements
└── Solutions/               # Implementations of problems
    ├── cpp/                # C++ implementations
    └── golang/             # Go implementations
```

## 🎯 What's Inside

### 1. Design Patterns

Comprehensive implementations of **23 classic Gang of Four (GoF) design patterns**, organized into three categories:

#### **Creational Patterns** (5 patterns)
Focus on object creation mechanisms, trying to create objects in a manner suitable to the situation.

- **Abstract Factory** - Creates families of related objects without specifying their concrete classes
- **Builder** - Constructs complex objects step by step
- **Factory Method** - Defines an interface for creating objects, letting subclasses decide which class to instantiate
- **Prototype** - Creates new objects by copying existing objects
- **Singleton** - Ensures a class has only one instance with global access

#### **Behavioral Patterns** (11 patterns)
Focus on communication between objects, what goes on between objects and how they operate together.

- **Chain of Responsibility** - Passes requests along a chain of handlers
- **Command** - Encapsulates a request as an object
- **Iterator** - Provides a way to access elements sequentially without exposing underlying representation
- **Mediator** - Reduces coupling between components by making them communicate through a mediator
- **Memento** - Captures and restores an object's internal state
- **Observer** - Defines a subscription mechanism to notify multiple objects about events
- **State** - Allows an object to alter its behavior when its internal state changes
- **Strategy** - Defines a family of algorithms and makes them interchangeable
- **Template Method** - Defines the skeleton of an algorithm, deferring some steps to subclasses
- **Visitor** - Separates algorithms from the objects on which they operate

#### **Structural Patterns** (7 patterns)
Focus on composing classes and objects to form larger structures while keeping them flexible and efficient.

- **Adapter** - Allows incompatible interfaces to work together
- **Bridge** - Separates abstraction from implementation
- **Composite** - Composes objects into tree structures to represent part-whole hierarchies
- **Decorator** - Adds responsibilities to objects dynamically
- **Facade** - Provides a simplified interface to a complex subsystem
- **Flyweight** - Shares common state between multiple objects instead of keeping all data in each object
- **Proxy** - Provides a placeholder for another object to control access to it

### 2. Real-World System Design Problems

Practical low-level design problems commonly asked in technical interviews, complete with:
- ✅ Detailed requirements
- ✅ Class diagrams and relationships
- ✅ Design patterns applied
- ✅ Thread-safety considerations
- ✅ Multiple language implementations

#### Available Problems:

| Problem | Description | C++ | Go |
|---------|-------------|-----|-----|
| **ATM System** | Design an Automated Teller Machine with card authentication, transactions, and account management | ✅ | ⏳ |
| **Coffee Vending Machine** | Design a machine that dispenses coffee with ingredient management | ✅ | ⏳ |
| **Logging Framework** | Build a flexible, thread-safe logging system with multiple outputs and log levels | ✅ | ✅ |
| **Parking Lot System** | Design a multi-level parking lot with different vehicle types and spot allocation | ✅ | ✅ |
| **Stack Overflow Clone** | Design a Q&A platform with voting, comments, and reputation system | ✅ | ✅ |
| **Task Management System** | Build a system for creating, assigning, and tracking tasks | ✅ | ✅ |
| **Traffic Signal Control** | Design an intelligent traffic signal system for intersections | ✅ | ✅ |
| **Vending Machine** | Design a vending machine with inventory and payment handling | ✅ | ✅ |

## 🚀 Getting Started

### Prerequisites

**For C++ implementations:**
```bash
# Requires C++17 or later
g++ --version  # or clang++
```

**For Go implementations:**
```bash
# Requires Go 1.18 or later
go version
```

### Running Examples

#### C++ Examples
```bash
cd Solutions/cpp/<problem-name>
g++ -std=c++17 *.cc -o main
./main
```

#### Go Examples
```bash
cd Solutions/golang/<problem-name>
go run .
```

### Example: Logging Framework

```bash
# Go implementation
cd Solutions/golang/loggingframework/example
go run main.go
```

## 📖 How to Use This Repository

### For Learning
1. **Start with Design Patterns**: Begin with the `Design-Pattern/` directory to understand fundamental patterns
2. **Study Problems**: Read problem statements in `Problems/` to understand requirements
3. **Analyze Solutions**: Review implementations in `Solutions/` to see patterns in action
4. **Compare Languages**: Compare C++ and Go implementations to understand language-specific idioms

### For Interview Preparation
1. Read the problem statement without looking at the solution
2. Try to design the system yourself (draw class diagrams, list classes/interfaces)
3. Identify which design patterns are applicable
4. Implement your solution
5. Compare with the provided implementation
6. Understand trade-offs and alternative approaches

### For Reference
- Use as a quick reference for design pattern implementation
- Copy and adapt code for your projects (with attribution)
- Learn language-specific best practices from implementations

## 🛠️ Key Design Principles Applied

All implementations in this repository follow SOLID principles:

- **S**ingle Responsibility Principle - Each class has one reason to change
- **O**pen/Closed Principle - Open for extension, closed for modification
- **L**iskov Substitution Principle - Subtypes must be substitutable for their base types
- **I**nterface Segregation Principle - Many specific interfaces are better than one general interface
- **D**ependency Inversion Principle - Depend on abstractions, not concretions

## 📝 Problem Structure

Each problem in the `Problems/` directory follows this structure:

```markdown
# Problem Title

## Requirements
- Numbered list of functional and non-functional requirements

## Implementations
- Links to C++ and Go implementations

## Classes, Interfaces and Enumerations
- Description of key components

## Design Patterns Used
- List of patterns applied and why
```

## 🔍 Featured Implementation: Logging Framework

A production-ready logging framework demonstrating:
- ✅ Strategy Pattern for different log outputs
- ✅ Singleton Pattern for logger instance
- ✅ Thread-safe design with mutexes
- ✅ Multiple log levels (DEBUG, INFO, WARNING, ERROR, FATAL)
- ✅ Extensible appender system (Console, File, Database)
- ✅ Log level filtering

**Quick Example:**
```go
logger := loggingframework.NewLogger("MyApp", loggingframework.LogLevelInfo)
logger.AddAppender(loggingframework.NewConsoleAppender())

fileAppender, _ := loggingframework.NewFileAppender("app.log")
defer fileAppender.Close()
logger.AddAppender(fileAppender)

logger.Info("Application started", "main")
logger.Error("An error occurred", "database")
```

## 📚 Additional Resources

### Books
- "Design Patterns: Elements of Reusable Object-Oriented Software" - Gang of Four
- "Head First Design Patterns" - Freeman & Freeman
- "Clean Architecture" - Robert C. Martin

### Online Resources
- [Refactoring.Guru - Design Patterns](https://refactoring.guru/design-patterns)
- [SourceMaking - Design Patterns](https://sourcemaking.com/design_patterns)


