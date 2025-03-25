#include <algorithm>
#include <iostream>
#include <list>
#include <string>
#include <gtest/gtest.h>
/**
 * The base Component class declares common operations for both simple and
 * complex objects of a composition.
 */
class Component {
  /**
   * @var Component
   */
 protected:
  Component *parent_;
  /**
   * Optionally, the base Component can declare an interface for setting and
   * accessing a parent of the component in a tree structure. It can also
   * provide some default implementation for these methods.
   */
 public:
  virtual ~Component() {}
  void SetParent(Component *parent) {
    this->parent_ = parent;
  }
  Component *GetParent() const {
    return this->parent_;
  }
  /**
   * In some cases, it would be beneficial to define the child-management
   * operations right in the base Component class. This way, you won't need to
   * expose any concrete component classes to the client code, even during the
   * object tree assembly. The downside is that these methods will be empty for
   * the leaf-level components.
   */
  virtual void Add(Component *component) {}
  virtual void Remove(Component *component) {}
  /**
   * You can provide a method that lets the client code figure out whether a
   * component can bear children.
   */
  virtual bool IsComposite() const {
    return false;
  }
  /**
   * The base Component may implement some default behavior or leave it to
   * concrete classes (by declaring the method containing the behavior as
   * "abstract").
   */
  virtual std::string Operation() const = 0;
};
/**
 * The Leaf class represents the end objects of a composition. A leaf can't have
 * any children.
 *
 * Usually, it's the Leaf objects that do the actual work, whereas Composite
 * objects only delegate to their sub-components.
 */
class Leaf : public Component {
 public:
  std::string Operation() const override {
    return "Leaf";
  }
};
/**
 * The Composite class represents the complex components that may have children.
 * Usually, the Composite objects delegate the actual work to their children and
 * then "sum-up" the result.
 */
class Composite : public Component {
  /**
   * @var \SplObjectStorage
   */
 protected:
  std::list<Component *> children_;

 public:
  /**
   * A composite object can add or remove other components (both simple or
   * complex) to or from its child list.
   */
  void Add(Component *component) override {
    this->children_.push_back(component);
    component->SetParent(this);
  }
  /**
   * Have in mind that this method removes the pointer to the list but doesn't
   * frees the
   *     memory, you should do it manually or better use smart pointers.
   */
  void Remove(Component *component) override {
    children_.remove(component);
    component->SetParent(nullptr);
  }
  bool IsComposite() const override {
    return true;
  }
  /**
   * The Composite executes its primary logic in a particular way. It traverses
   * recursively through all its children, collecting and summing their results.
   * Since the composite's children pass these calls to their children and so
   * forth, the whole object tree is traversed as a result.
   */
  std::string Operation() const override {
    std::string result;
    for (const Component *c : children_) {
      if (c == children_.back()) {
        result += c->Operation();
      } else {
        result += c->Operation() + "+";
      }
    }
    return "Branch(" + result + ")";
  }
};
/**
 * The client code works with all of the components via the base interface.
 */
void ClientCode(Component *component) {
  // ...
  std::cout << "RESULT: " << component->Operation();
  // ...
}

/**
 * Thanks to the fact that the child-management operations are declared in the
 * base Component class, the client code can work with any component, simple or
 * complex, without depending on their concrete classes.
 */
void ClientCode2(Component *component1, Component *component2) {
  // ...
  if (component1->IsComposite()) {
    component1->Add(component2);
  }
  std::cout << "RESULT: " << component1->Operation();
  // ...
}

/**
 * This way the client code can support the simple leaf components...
 */


TEST(CompositeTest, LeafOperation) {
  Leaf *leaf = new Leaf();
  EXPECT_EQ("Leaf", leaf->Operation());
  EXPECT_FALSE(leaf->IsComposite());
  delete leaf;
}

TEST(CompositeTest, CompositeOperation) {
  Composite *tree = new Composite();
  Leaf *leaf1 = new Leaf();
  Leaf *leaf2 = new Leaf();
  
  tree->Add(leaf1);
  tree->Add(leaf2);
  
  EXPECT_EQ("Branch(Leaf+Leaf)", tree->Operation());
  EXPECT_TRUE(tree->IsComposite());
  
  delete tree;
}

TEST(CompositeTest, NestedComposite) {
  Composite *tree = new Composite();
  Composite *branch = new Composite();
  Leaf *leaf1 = new Leaf();
  Leaf *leaf2 = new Leaf();
  Leaf *leaf3 = new Leaf();
  
  branch->Add(leaf1);
  branch->Add(leaf2);
  tree->Add(branch);
  tree->Add(leaf3);
  
  EXPECT_EQ("Branch(Branch(Leaf+Leaf)+Leaf)", tree->Operation());
  
  delete tree;
}

TEST(CompositeTest, ClientCode) {
  Leaf *simple = new Leaf();
  Composite *tree = new Composite();
  Leaf *leaf1 = new Leaf();
  Leaf *leaf2 = new Leaf();
  
  tree->Add(leaf1);
  tree->Add(leaf2);
  
  testing::internal::CaptureStdout();
  ClientCode(simple);
  std::string output1 = testing::internal::GetCapturedStdout();
  EXPECT_TRUE(output1.find("RESULT: Leaf") != std::string::npos);
  
  testing::internal::CaptureStdout();
  ClientCode(tree);
  std::string output2 = testing::internal::GetCapturedStdout();
  EXPECT_TRUE(output2.find("RESULT: Branch(Leaf+Leaf)") != std::string::npos);
  
  testing::internal::CaptureStdout();
  ClientCode2(tree, simple);
  std::string output3 = testing::internal::GetCapturedStdout();
  EXPECT_TRUE(output3.find("RESULT: Branch(Leaf+Leaf+Leaf)") != std::string::npos);
  
  delete simple;
  delete tree;
}