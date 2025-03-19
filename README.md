# Low-Level-Design

## Singleton
To implement the singleton pattern, we must prevent external objects from creating instances of the singleton class. Only the singleton class should be permitted to create its own objects.

This can be achieved by making the **constructor private** and providing a **static method** for external objects to access it.


## Factory
The Factory Method separates product construction code from the code that actually uses the product. Therefore it’s easier to extend the product construction code independently from the rest of the code.