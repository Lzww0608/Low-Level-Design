package main

import (
	"fmt"
	"sync"
)

// VendingMachineDemo 演示自动售货机的使用
func main() {
	fmt.Println("===== 自动售货机演示 =====\n")

	// 获取售货机单例
	vm := GetVendingMachine()

	// 创建产品
	cola := NewProduct("可乐", 3.50)
	chips := NewProduct("薯片", 5.00)
	water := NewProduct("矿泉水", 2.00)
	chocolate := NewProduct("巧克力", 8.00)

	// 补充产品库存
	fmt.Println("--- 补充产品库存 ---")
	vm.AddProduct(cola, 5)
	vm.AddProduct(chips, 3)
	vm.AddProduct(water, 10)
	vm.AddProduct(chocolate, 2)

	// 显示产品列表
	vm.DisplayProducts()

	// 演示购买流程
	fmt.Println("\n--- 演示1: 购买可乐 ---")
	err := vm.SelectProduct(cola)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}

	// 投入硬币
	vm.InsertCoin(Dollar)  // 1元
	vm.InsertCoin(Dollar)  // 1元
	vm.InsertCoin(Dollar)  // 1元
	vm.InsertCoin(Quarter) // 0.25元
	vm.InsertCoin(Quarter) // 0.25元

	// 发放产品
	err = vm.DispenseProduct()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}

	// 返回找零
	change := vm.ReturnChange()
	fmt.Printf("交易完成，找零: %.2f元\n", change)

	// 演示购买流程2 - 使用纸币
	fmt.Println("\n--- 演示2: 购买巧克力（使用纸币）---")
	err = vm.SelectProduct(chocolate)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}

	vm.InsertNote(Ten) // 10元

	err = vm.DispenseProduct()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}

	change = vm.ReturnChange()
	fmt.Printf("交易完成，找零: %.2f元\n", change)

	// 演示购买流程3 - 金额不足
	fmt.Println("\n--- 演示3: 金额不足 ---")
	err = vm.SelectProduct(chips)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}

	vm.InsertCoin(Dollar) // 只投入1元

	err = vm.DispenseProduct()
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}

	// 取消交易，退还金额
	change = vm.ReturnChange()
	fmt.Printf("取消交易，退还: %.2f元\n", change)

	// 演示购买流程4 - 产品售罄
	fmt.Println("\n--- 演示4: 产品售罄 ---")
	// 先买光巧克力
	for i := 0; i < 3; i++ {
		err = vm.SelectProduct(chocolate)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			break
		}
		vm.InsertNote(Ten)
		vm.DispenseProduct()
		vm.ReturnChange()
	}

	// 显示更新后的库存
	vm.DisplayProducts()

	// 取出收集的金钱
	fmt.Println("\n--- 取出收集的金钱 ---")
	collected := vm.CollectMoney()
	fmt.Printf("总收入: %.2f元\n", collected)

	// 演示并发购买
	fmt.Println("\n--- 演示5: 并发购买 ---")
	vm.AddProduct(water, 20) // 补充库存

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 每个goroutine执行完整的购买流程
			err := vm.SelectProduct(water)
			if err != nil {
				fmt.Printf("[协程%d] 选择产品错误: %v\n", id, err)
				return
			}

			err = vm.InsertCoin(Dollar)
			if err != nil {
				fmt.Printf("[协程%d] 投币错误: %v\n", id, err)
				vm.ReturnChange()
				return
			}

			err = vm.InsertCoin(Dollar)
			if err != nil {
				fmt.Printf("[协程%d] 投币错误: %v\n", id, err)
				vm.ReturnChange()
				return
			}

			err = vm.DispenseProduct()
			if err != nil {
				fmt.Printf("[协程%d] 发放错误: %v\n", id, err)
				vm.ReturnChange()
				return
			}

			change := vm.ReturnChange()
			fmt.Printf("[协程%d] 购买成功，找零: %.2f元\n", id, change)
		}(i)
	}

	wg.Wait()

	// 显示最终库存
	fmt.Println("\n--- 最终库存 ---")
	vm.DisplayProducts()

	fmt.Println("\n===== 演示结束 =====")
}
