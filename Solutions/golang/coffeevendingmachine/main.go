package coffeevendingmachine

import (
	"fmt"
	"sync"
)

// Demo 演示咖啡售货机的使用
func Demo() {
	fmt.Println("===== 咖啡售货机演示 =====")

	// 1. 创建并初始化咖啡机
	machine := NewMachine()
	machine.InitializeMenu()
	machine.InitializeInventory()

	// 2. 显示菜单
	fmt.Println(machine.DisplayMenu())

	// 3. 显示当前库存
	fmt.Println("当前库存:")
	inventory := machine.CheckInventoryStatus()
	for name, quantity := range inventory {
		fmt.Printf("  %s: %d\n", name, quantity)
	}
	fmt.Println()

	// 4. 用户1购买 Espresso
	fmt.Println("--- 用户1购买 Espresso ---")
	payment1 := NewPayment(5.0)
	change1, err := machine.SelectAndDispenseCoffee(ESPRESSO, payment1)
	if err != nil {
		fmt.Printf("购买失败: %v\n", err)
	} else {
		fmt.Printf("购买成功! 找零: $%.2f\n", change1)
	}
	fmt.Println()

	// 5. 用户2购买 Cappuccino
	fmt.Println("--- 用户2购买 Cappuccino ---")
	payment2 := NewPayment(10.0)
	change2, err := machine.SelectAndDispenseCoffee(CAPPUCCINO, payment2)
	if err != nil {
		fmt.Printf("购买失败: %v\n", err)
	} else {
		fmt.Printf("购买成功! 找零: $%.2f\n", change2)
	}
	fmt.Println()

	// 6. 用户3购买 Latte
	fmt.Println("--- 用户3购买 Latte ---")
	payment3 := NewPayment(6.0)
	change3, err := machine.SelectAndDispenseCoffee(LATTE, payment3)
	if err != nil {
		fmt.Printf("购买失败: %v\n", err)
	} else {
		fmt.Printf("购买成功! 找零: $%.2f\n", change3)
	}
	fmt.Println()

	// 7. 检查库存状态
	fmt.Println("购买后的库存:")
	inventory = machine.CheckInventoryStatus()
	for name, quantity := range inventory {
		fmt.Printf("  %s: %d\n", name, quantity)
	}
	fmt.Println()

	// 8. 检查低库存警告
	lowStock := machine.GetLowStockAlert()
	if len(lowStock) > 0 {
		fmt.Println("⚠️  低库存警告:")
		for _, name := range lowStock {
			fmt.Printf("  - %s\n", name)
		}
		fmt.Println()
	}

	// 9. 补充原料
	fmt.Println("--- 补充原料 ---")
	machine.RefillIngredient("Coffee Beans", 50)
	machine.RefillIngredient("Milk", 30)
	fmt.Println()

	// 10. 演示并发购买
	fmt.Println("--- 演示并发购买 (5个用户同时购买) ---")
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()

			var coffeeType CoffeeType
			switch userID % 3 {
			case 0:
				coffeeType = ESPRESSO
			case 1:
				coffeeType = CAPPUCCINO
			case 2:
				coffeeType = LATTE
			}

			payment := NewPayment(10.0)
			change, err := machine.SelectAndDispenseCoffee(coffeeType, payment)
			if err != nil {
				fmt.Printf("用户%d 购买失败: %v\n", userID, err)
			} else {
				fmt.Printf("用户%d 购买成功! 找零: $%.2f\n", userID, change)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println()

	// 11. 最终库存状态
	fmt.Println("最终库存:")
	inventory = machine.CheckInventoryStatus()
	for name, quantity := range inventory {
		fmt.Printf("  %s: %d\n", name, quantity)
	}
	fmt.Println()

	fmt.Println("===== 演示结束 =====")
}
