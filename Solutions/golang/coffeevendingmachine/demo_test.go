package coffeevendingmachine

import (
	"fmt"
	"testing"
)

// TestDemo 演示咖啡售货机的完整功能
func TestDemo(t *testing.T) {
	Demo()
}

// ExampleMachine_SelectAndDispenseCoffee 示例：购买咖啡
func ExampleMachine_SelectAndDispenseCoffee() {
	machine := NewMachine()
	machine.InitializeMenu()
	machine.InitializeInventory()

	// 购买 Espresso
	payment := NewPayment(5.0)
	change, err := machine.SelectAndDispenseCoffee(ESPRESSO, payment)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Change: $%.2f\n", change)
}

// ExampleMachine_DisplayMenu 示例：显示菜单
func ExampleMachine_DisplayMenu() {
	machine := NewMachine()
	machine.InitializeMenu()

	menu := machine.DisplayMenu()
	fmt.Print(menu)
}

// ExampleInventory_GetLowStockIngredients 示例：检查低库存
func ExampleInventory_GetLowStockIngredients() {
	inventory := NewInventory()
	inventory.SetLowStockThreshold(15)
	inventory.AddIngredient("Coffee Beans", 5) // 低于阈值
	inventory.AddIngredient("Milk", 20)        // 高于阈值

	lowStock := inventory.GetLowStockIngredients()
	for _, name := range lowStock {
		fmt.Printf("Low stock: %s\n", name)
	}
}

// ExamplePayment_ProcessPayment 示例：处理支付
func ExamplePayment_ProcessPayment() {
	payment := NewPayment(10.0)
	change, err := payment.ProcessPayment(4.5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Payment successful. Change: $%.2f\n", change)
}
