package main

import "fmt"

// VendingMachineState 定义售货机在不同状态下的行为接口
type VendingMachineState interface {
	// SelectProduct 选择产品
	SelectProduct(product *Product) error
	// InsertCoin 投入硬币
	InsertCoin(coin Coin) error
	// InsertNote 投入纸币
	InsertNote(note Note) error
	// DispenseProduct 发放产品
	DispenseProduct() error
	// ReturnChange 返回找零
	ReturnChange() float64
}

// IdleState 空闲状态 - 等待用户选择产品
type IdleState struct {
	vendingMachine *VendingMachine
}

func NewIdleState(vm *VendingMachine) *IdleState {
	return &IdleState{vendingMachine: vm}
}

func (s *IdleState) SelectProduct(product *Product) error {
	vm := s.vendingMachine

	// 检查产品是否有库存
	if !vm.inventory.IsAvailable(product) {
		return fmt.Errorf("产品 %s 已售罄", product.Name)
	}

	vm.selectedProduct = product
	vm.SetState(vm.readyState)
	fmt.Printf("已选择产品: %s, 价格: %.2f元\n", product.Name, product.Price)
	return nil
}

func (s *IdleState) InsertCoin(coin Coin) error {
	return fmt.Errorf("请先选择产品")
}

func (s *IdleState) InsertNote(note Note) error {
	return fmt.Errorf("请先选择产品")
}

func (s *IdleState) DispenseProduct() error {
	return fmt.Errorf("请先选择产品")
}

func (s *IdleState) ReturnChange() float64 {
	// 空闲状态没有找零
	return 0
}

// ReadyState 准备状态 - 等待用户投币
type ReadyState struct {
	vendingMachine *VendingMachine
}

func NewReadyState(vm *VendingMachine) *ReadyState {
	return &ReadyState{vendingMachine: vm}
}

func (s *ReadyState) SelectProduct(product *Product) error {
	// 允许重新选择产品
	vm := s.vendingMachine

	if !vm.inventory.IsAvailable(product) {
		return fmt.Errorf("产品 %s 已售罄", product.Name)
	}

	vm.selectedProduct = product
	fmt.Printf("已重新选择产品: %s, 价格: %.2f元\n", product.Name, product.Price)
	return nil
}

func (s *ReadyState) InsertCoin(coin Coin) error {
	vm := s.vendingMachine
	vm.totalPayment += coin.Value()
	fmt.Printf("投入硬币: %.2f元, 当前已投入: %.2f元\n", coin.Value(), vm.totalPayment)

	// 检查是否已支付足够金额
	s.checkPaymentSufficient()
	return nil
}

func (s *ReadyState) InsertNote(note Note) error {
	vm := s.vendingMachine
	vm.totalPayment += note.Value()
	fmt.Printf("投入纸币: %.2f元, 当前已投入: %.2f元\n", note.Value(), vm.totalPayment)

	// 检查是否已支付足够金额
	s.checkPaymentSufficient()
	return nil
}

func (s *ReadyState) checkPaymentSufficient() {
	vm := s.vendingMachine
	if vm.totalPayment >= vm.selectedProduct.Price {
		vm.SetState(vm.dispenseState)
	}
}

func (s *ReadyState) DispenseProduct() error {
	return fmt.Errorf("金额不足，还需支付: %.2f元",
		s.vendingMachine.selectedProduct.Price-s.vendingMachine.totalPayment)
}

func (s *ReadyState) ReturnChange() float64 {
	vm := s.vendingMachine
	change := vm.totalPayment
	vm.totalPayment = 0
	vm.selectedProduct = nil
	vm.SetState(vm.idleState)
	fmt.Printf("取消交易，退还金额: %.2f元\n", change)
	return change
}

// DispenseState 发放状态 - 准备发放产品和找零
type DispenseState struct {
	vendingMachine *VendingMachine
}

func NewDispenseState(vm *VendingMachine) *DispenseState {
	return &DispenseState{vendingMachine: vm}
}

func (s *DispenseState) SelectProduct(product *Product) error {
	return fmt.Errorf("正在处理当前交易，请稍候")
}

func (s *DispenseState) InsertCoin(coin Coin) error {
	return fmt.Errorf("正在处理当前交易，请稍候")
}

func (s *DispenseState) InsertNote(note Note) error {
	return fmt.Errorf("正在处理当前交易，请稍候")
}

func (s *DispenseState) DispenseProduct() error {
	vm := s.vendingMachine

	// 从库存中移除产品
	err := vm.inventory.RemoveProduct(vm.selectedProduct)
	if err != nil {
		vm.SetState(vm.idleState)
		return err
	}

	fmt.Printf("发放产品: %s\n", vm.selectedProduct.Name)
	return nil
}

func (s *DispenseState) ReturnChange() float64 {
	vm := s.vendingMachine
	change := vm.totalPayment - vm.selectedProduct.Price

	if change > 0 {
		fmt.Printf("找零: %.2f元\n", change)
	}

	// 更新收入
	vm.collectedMoney += vm.selectedProduct.Price

	// 重置状态
	vm.totalPayment = 0
	vm.selectedProduct = nil
	vm.SetState(vm.idleState)

	return change
}
