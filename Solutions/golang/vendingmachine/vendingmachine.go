package main

import (
	"fmt"
	"sync"
)

// VendingMachine 自动售货机 - 单例模式
type VendingMachine struct {
	inventory       *Inventory
	currentState    VendingMachineState
	idleState       VendingMachineState
	readyState      VendingMachineState
	dispenseState   VendingMachineState
	selectedProduct *Product
	totalPayment    float64
	collectedMoney  float64 // 收集的金钱总额
	mu              sync.Mutex
}

var (
	instance *VendingMachine
	once     sync.Once
)

// GetVendingMachine 获取售货机单例
func GetVendingMachine() *VendingMachine {
	once.Do(func() {
		instance = &VendingMachine{
			inventory: NewInventory(),
		}
		// 初始化各种状态
		instance.idleState = NewIdleState(instance)
		instance.readyState = NewReadyState(instance)
		instance.dispenseState = NewDispenseState(instance)
		instance.currentState = instance.idleState
	})
	return instance
}

// ResetInstance 重置单例（仅用于测试）
func ResetInstance() {
	once = sync.Once{}
	instance = nil
}

// SetState 设置当前状态
func (vm *VendingMachine) SetState(state VendingMachineState) {
	vm.currentState = state
}

// SelectProduct 选择产品
func (vm *VendingMachine) SelectProduct(product *Product) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	return vm.currentState.SelectProduct(product)
}

// InsertCoin 投入硬币
func (vm *VendingMachine) InsertCoin(coin Coin) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	return vm.currentState.InsertCoin(coin)
}

// InsertNote 投入纸币
func (vm *VendingMachine) InsertNote(note Note) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	return vm.currentState.InsertNote(note)
}

// DispenseProduct 发放产品
func (vm *VendingMachine) DispenseProduct() error {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	return vm.currentState.DispenseProduct()
}

// ReturnChange 返回找零
func (vm *VendingMachine) ReturnChange() float64 {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	return vm.currentState.ReturnChange()
}

// AddProduct 补充产品库存
func (vm *VendingMachine) AddProduct(product *Product, quantity int) {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	vm.inventory.AddProduct(product, quantity)
	fmt.Printf("补充产品: %s, 数量: %d\n", product.Name, quantity)
}

// GetInventory 获取库存
func (vm *VendingMachine) GetInventory() *Inventory {
	return vm.inventory
}

// GetCollectedMoney 获取收集的金钱总额
func (vm *VendingMachine) GetCollectedMoney() float64 {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	return vm.collectedMoney
}

// CollectMoney 取出收集的金钱
func (vm *VendingMachine) CollectMoney() float64 {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	collected := vm.collectedMoney
	vm.collectedMoney = 0
	fmt.Printf("取出金钱: %.2f元\n", collected)
	return collected
}

// DisplayProducts 显示所有产品信息
func (vm *VendingMachine) DisplayProducts() {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	fmt.Println("\n========== 产品列表 ==========")
	products := vm.inventory.GetAllProducts()
	for product, quantity := range products {
		fmt.Printf("产品: %s, 价格: %.2f元, 库存: %d\n",
			product.Name, product.Price, quantity)
	}
	fmt.Println("==============================")
}
