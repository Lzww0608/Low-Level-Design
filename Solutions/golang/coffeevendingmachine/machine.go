package coffeevendingmachine

import (
	"fmt"
	"sync"
)

type Machine struct {
	coffeeMenu []Coffee
	inventory  *Inventory
	mu         sync.Mutex // 确保机器操作的线程安全
}

func NewMachine() *Machine {
	return &Machine{
		coffeeMenu: make([]Coffee, 0),
		inventory:  NewInventory(),
	}
}

// InitializeInventory 初始化库存
func (m *Machine) InitializeInventory() {
	m.inventory.AddIngredient("Coffee Beans", 100)
	m.inventory.AddIngredient("Milk", 100)
	m.inventory.AddIngredient("Water", 200)
	m.inventory.AddIngredient("Sugar", 50)
	m.inventory.AddIngredient("Chocolate", 50)
}

// InitializeMenu 初始化咖啡菜单
func (m *Machine) InitializeMenu() {
	// Espresso
	espresso := NewCoffee(ESPRESSO, "Espresso", 3.5, []Ingredient{
		{Name: "Coffee Beans", Quantity: 2},
		{Name: "Water", Quantity: 1},
	})
	m.AddCoffee(*espresso)

	// Cappuccino
	cappuccino := NewCoffee(CAPPUCCINO, "Cappuccino", 4.5, []Ingredient{
		{Name: "Coffee Beans", Quantity: 2},
		{Name: "Water", Quantity: 1},
		{Name: "Milk", Quantity: 2},
	})
	m.AddCoffee(*cappuccino)

	// Latte
	latte := NewCoffee(LATTE, "Latte", 5.0, []Ingredient{
		{Name: "Coffee Beans", Quantity: 2},
		{Name: "Water", Quantity: 1},
		{Name: "Milk", Quantity: 3},
	})
	m.AddCoffee(*latte)
}

// AddCoffee 添加咖啡到菜单
func (m *Machine) AddCoffee(coffee Coffee) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, c := range m.coffeeMenu {
		if c.GetType() == coffee.GetType() {
			return
		}
	}
	m.coffeeMenu = append(m.coffeeMenu, coffee)
}

// RemoveCoffee 从菜单移除咖啡
func (m *Machine) RemoveCoffee(coffee Coffee) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, c := range m.coffeeMenu {
		if c.GetType() == coffee.GetType() {
			m.coffeeMenu = append(m.coffeeMenu[:i], m.coffeeMenu[i+1:]...)
			break
		}
	}
}

// GetCoffeeMenu 获取咖啡菜单
func (m *Machine) GetCoffeeMenu() []Coffee {
	m.mu.Lock()
	defer m.mu.Unlock()

	result := make([]Coffee, len(m.coffeeMenu))
	copy(result, m.coffeeMenu)
	return result
}

// DisplayMenu 显示菜单
func (m *Machine) DisplayMenu() string {
	m.mu.Lock()
	defer m.mu.Unlock()

	menu := "========== Coffee Menu ==========\n"
	for i, coffee := range m.coffeeMenu {
		menu += fmt.Sprintf("%d. %s - $%.2f\n", i+1, coffee.GetName(), coffee.GetPrice())
	}
	menu += "================================\n"
	return menu
}

// GetInventory 获取库存
func (m *Machine) GetInventory() *Inventory {
	return m.inventory
}

// GetCoffee 根据类型获取咖啡
func (m *Machine) GetCoffee(coffeeType CoffeeType) *Coffee {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, c := range m.coffeeMenu {
		if c.GetType() == coffeeType {
			return &c
		}
	}
	return nil
}

// SelectAndDispenseCoffee 选择并分配咖啡（包含支付处理）
func (m *Machine) SelectAndDispenseCoffee(coffeeType CoffeeType, payment *Payment) (float64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 1. 检查咖啡是否存在
	coffee := m.getCoffeeUnsafe(coffeeType)
	if coffee == nil {
		return 0, fmt.Errorf("coffee type not found")
	}

	// 2. 处理支付
	change, err := payment.ProcessPayment(coffee.GetPrice())
	if err != nil {
		return 0, err
	}

	// 3. 检查原料是否足够
	if !m.inventory.HasEnoughIngredients(coffee.GetRecipe()) {
		return 0, fmt.Errorf("insufficient ingredients for %s", coffee.GetName())
	}

	// 4. 扣除原料
	if !m.inventory.DeductIngredients(coffee.GetRecipe()) {
		return 0, fmt.Errorf("failed to deduct ingredients")
	}

	// 5. 检查低库存
	lowStock := m.inventory.GetLowStockIngredients()
	if len(lowStock) > 0 {
		fmt.Printf("Warning: Low stock for ingredients: %v\n", lowStock)
	}

	// 6. 分配咖啡
	fmt.Printf("Dispensing %s...\n", coffee.GetName())

	return change, nil
}

// getCoffeeUnsafe 内部方法，不加锁获取咖啡（调用者需要持有锁）
func (m *Machine) getCoffeeUnsafe(coffeeType CoffeeType) *Coffee {
	for _, c := range m.coffeeMenu {
		if c.GetType() == coffeeType {
			return &c
		}
	}
	return nil
}

// CheckInventoryStatus 检查库存状态
func (m *Machine) CheckInventoryStatus() map[string]int {
	return m.inventory.GetAllIngredients()
}

// GetLowStockAlert 获取低库存警告
func (m *Machine) GetLowStockAlert() []string {
	return m.inventory.GetLowStockIngredients()
}

// RefillIngredient 补充原料
func (m *Machine) RefillIngredient(name string, quantity int) {
	m.inventory.AddIngredient(name, quantity)
	fmt.Printf("Refilled %s: +%d\n", name, quantity)
}
