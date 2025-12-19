package coffeevendingmachine

import (
	"sync"
	"testing"
)

// TestMachineInitialization 测试机器初始化
func TestMachineInitialization(t *testing.T) {
	machine := NewMachine()
	if machine == nil {
		t.Fatal("Machine initialization failed")
	}

	machine.InitializeMenu()
	menu := machine.GetCoffeeMenu()
	if len(menu) != 3 {
		t.Errorf("Expected 3 coffees in menu, got %d", len(menu))
	}

	machine.InitializeInventory()
	inventory := machine.CheckInventoryStatus()
	if len(inventory) == 0 {
		t.Error("Inventory should not be empty after initialization")
	}
}

// TestDisplayMenu 测试显示菜单
func TestDisplayMenu(t *testing.T) {
	machine := NewMachine()
	machine.InitializeMenu()

	menuStr := machine.DisplayMenu()
	if menuStr == "" {
		t.Error("Menu display should not be empty")
	}
	t.Logf("Menu:\n%s", menuStr)
}

// TestAddAndRemoveCoffee 测试添加和删除咖啡
func TestAddAndRemoveCoffee(t *testing.T) {
	machine := NewMachine()

	// 添加咖啡
	coffee := NewCoffee(ESPRESSO, "Espresso", 3.5, []Ingredient{
		{Name: "Coffee Beans", Quantity: 2},
		{Name: "Water", Quantity: 1},
	})
	machine.AddCoffee(*coffee)

	menu := machine.GetCoffeeMenu()
	if len(menu) != 1 {
		t.Errorf("Expected 1 coffee in menu, got %d", len(menu))
	}

	// 删除咖啡
	machine.RemoveCoffee(*coffee)
	menu = machine.GetCoffeeMenu()
	if len(menu) != 0 {
		t.Errorf("Expected 0 coffees in menu after removal, got %d", len(menu))
	}
}

// TestPaymentProcessing 测试支付处理
func TestPaymentProcessing(t *testing.T) {
	tests := []struct {
		name           string
		paymentAmount  float64
		coffeePrice    float64
		expectError    bool
		expectedChange float64
	}{
		{"Exact payment", 5.0, 5.0, false, 0.0},
		{"Overpayment", 10.0, 5.0, false, 5.0},
		{"Insufficient payment", 3.0, 5.0, true, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payment := NewPayment(tt.paymentAmount)
			change, err := payment.ProcessPayment(tt.coffeePrice)

			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !tt.expectError && change != tt.expectedChange {
				t.Errorf("Expected change %.2f, got %.2f", tt.expectedChange, change)
			}
		})
	}
}

// TestInventoryManagement 测试库存管理
func TestInventoryManagement(t *testing.T) {
	inventory := NewInventory()

	// 添加原料
	inventory.AddIngredient("Coffee Beans", 10)
	inventory.AddIngredient("Milk", 20)

	// 检查数量
	if quantity := inventory.GetIngredientQuantity("Coffee Beans"); quantity != 10 {
		t.Errorf("Expected 10 coffee beans, got %d", quantity)
	}

	// 检查足够的原料
	recipe := []Ingredient{
		{Name: "Coffee Beans", Quantity: 5},
		{Name: "Milk", Quantity: 10},
	}
	if !inventory.HasEnoughIngredients(recipe) {
		t.Error("Should have enough ingredients")
	}

	// 扣除原料
	if !inventory.DeductIngredients(recipe) {
		t.Error("Failed to deduct ingredients")
	}

	// 验证扣除后的数量
	if quantity := inventory.GetIngredientQuantity("Coffee Beans"); quantity != 5 {
		t.Errorf("Expected 5 coffee beans after deduction, got %d", quantity)
	}
}

// TestLowStockAlert 测试低库存警告
func TestLowStockAlert(t *testing.T) {
	inventory := NewInventory()
	inventory.SetLowStockThreshold(15)

	inventory.AddIngredient("Coffee Beans", 10) // 低于阈值
	inventory.AddIngredient("Milk", 20)         // 高于阈值

	lowStock := inventory.GetLowStockIngredients()
	if len(lowStock) != 1 {
		t.Errorf("Expected 1 low stock ingredient, got %d", len(lowStock))
	}

	found := false
	for _, name := range lowStock {
		if name == "Coffee Beans" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Coffee Beans should be in low stock list")
	}
}

// TestSelectAndDispenseCoffee 测试选择和分配咖啡
func TestSelectAndDispenseCoffee(t *testing.T) {
	machine := NewMachine()
	machine.InitializeMenu()
	machine.InitializeInventory()

	// 成功购买咖啡
	payment := NewPayment(10.0)
	change, err := machine.SelectAndDispenseCoffee(ESPRESSO, payment)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if change <= 0 {
		t.Error("Expected positive change")
	}
	t.Logf("Change: $%.2f", change)
}

// TestInsufficientPayment 测试支付不足
func TestInsufficientPayment(t *testing.T) {
	machine := NewMachine()
	machine.InitializeMenu()
	machine.InitializeInventory()

	payment := NewPayment(1.0) // 不足以购买任何咖啡
	_, err := machine.SelectAndDispenseCoffee(ESPRESSO, payment)
	if err == nil {
		t.Error("Expected error for insufficient payment")
	}
	t.Logf("Expected error: %v", err)
}

// TestInsufficientIngredients 测试原料不足
func TestInsufficientIngredients(t *testing.T) {
	machine := NewMachine()
	machine.InitializeMenu()

	// 添加刚好够一杯Espresso的原料（Espresso需要2个Coffee Beans和1个Water）
	machine.GetInventory().AddIngredient("Coffee Beans", 2)
	machine.GetInventory().AddIngredient("Water", 1)

	payment := NewPayment(10.0)

	// 第一次应该成功
	_, err := machine.SelectAndDispenseCoffee(ESPRESSO, payment)
	if err != nil {
		t.Errorf("First purchase should succeed: %v", err)
	}

	// 第二次应该失败（原料不足）
	payment2 := NewPayment(10.0)
	_, err = machine.SelectAndDispenseCoffee(ESPRESSO, payment2)
	if err == nil {
		t.Error("Expected error for insufficient ingredients")
	}
	t.Logf("Expected error: %v", err)
}

// TestConcurrentOrders 测试并发订单（线程安全）
func TestConcurrentOrders(t *testing.T) {
	machine := NewMachine()
	machine.InitializeMenu()
	machine.InitializeInventory()

	// 添加足够的库存
	machine.GetInventory().AddIngredient("Coffee Beans", 100)
	machine.GetInventory().AddIngredient("Water", 100)
	machine.GetInventory().AddIngredient("Milk", 100)

	numOrders := 50
	var wg sync.WaitGroup
	successCount := 0
	var mu sync.Mutex

	for i := 0; i < numOrders; i++ {
		wg.Add(1)
		go func(orderNum int) {
			defer wg.Done()

			payment := NewPayment(10.0)
			coffeeType := ESPRESSO
			if orderNum%3 == 1 {
				coffeeType = CAPPUCCINO
			} else if orderNum%3 == 2 {
				coffeeType = LATTE
			}

			_, err := machine.SelectAndDispenseCoffee(coffeeType, payment)
			if err == nil {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()

	t.Logf("Successfully processed %d out of %d concurrent orders", successCount, numOrders)

	if successCount == 0 {
		t.Error("Expected at least some successful orders")
	}
}

// TestRefillIngredient 测试补充原料
func TestRefillIngredient(t *testing.T) {
	machine := NewMachine()
	machine.InitializeInventory()

	initialQuantity := machine.GetInventory().GetIngredientQuantity("Coffee Beans")

	machine.RefillIngredient("Coffee Beans", 50)

	newQuantity := machine.GetInventory().GetIngredientQuantity("Coffee Beans")
	expectedQuantity := initialQuantity + 50

	if newQuantity != expectedQuantity {
		t.Errorf("Expected %d coffee beans after refill, got %d", expectedQuantity, newQuantity)
	}
}

// TestGetCoffee 测试获取特定咖啡
func TestGetCoffee(t *testing.T) {
	machine := NewMachine()
	machine.InitializeMenu()

	// 测试存在的咖啡
	coffee := machine.GetCoffee(ESPRESSO)
	if coffee == nil {
		t.Error("Expected to find Espresso")
	}
	if coffee.GetName() != "Espresso" {
		t.Errorf("Expected coffee name 'Espresso', got '%s'", coffee.GetName())
	}

	// 测试不存在的咖啡
	coffee = machine.GetCoffee(MOCHA) // 菜单中没有
	if coffee != nil {
		t.Error("Expected nil for non-existent coffee type")
	}
}

// TestInventoryThreadSafety 测试库存的线程安全性
func TestInventoryThreadSafety(t *testing.T) {
	inventory := NewInventory()
	inventory.AddIngredient("Test Ingredient", 1000)

	var wg sync.WaitGroup
	numGoroutines := 100

	// 多个goroutine同时读写库存
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 读取
			_ = inventory.GetIngredientQuantity("Test Ingredient")

			// 扣除
			recipe := []Ingredient{{Name: "Test Ingredient", Quantity: 1}}
			inventory.DeductIngredients(recipe)

			// 检查低库存
			_ = inventory.GetLowStockIngredients()
		}()
	}

	wg.Wait()

	// 验证最终数量
	finalQuantity := inventory.GetIngredientQuantity("Test Ingredient")
	expectedQuantity := 1000 - numGoroutines

	if finalQuantity != expectedQuantity {
		t.Errorf("Expected final quantity %d, got %d", expectedQuantity, finalQuantity)
	}
}

// BenchmarkSelectAndDispenseCoffee 性能测试
func BenchmarkSelectAndDispenseCoffee(b *testing.B) {
	machine := NewMachine()
	machine.InitializeMenu()
	machine.InitializeInventory()

	// 添加足够的库存
	machine.GetInventory().AddIngredient("Coffee Beans", 100000)
	machine.GetInventory().AddIngredient("Water", 100000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		payment := NewPayment(10.0)
		_, _ = machine.SelectAndDispenseCoffee(ESPRESSO, payment)
	}
}
