package main

import (
	"fmt"
	"sync"
)

// Inventory 管理售货机中的产品库存
// 使用 sync.RWMutex 确保线程安全
type Inventory struct {
	products map[*Product]int
	mu       sync.RWMutex
}

// NewInventory 创建一个新的库存管理器
func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[*Product]int),
	}
}

// AddProduct 添加产品到库存
func (inv *Inventory) AddProduct(product *Product, quantity int) {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	inv.products[product] += quantity
}

// RemoveProduct 从库存中移除一个产品
func (inv *Inventory) RemoveProduct(product *Product) error {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	if qty, exists := inv.products[product]; !exists || qty <= 0 {
		return fmt.Errorf("产品 %s 已售罄", product.Name)
	}

	inv.products[product]--
	return nil
}

// GetQuantity 获取产品的库存数量
func (inv *Inventory) GetQuantity(product *Product) int {
	inv.mu.RLock()
	defer inv.mu.RUnlock()
	return inv.products[product]
}

// IsAvailable 检查产品是否有库存
func (inv *Inventory) IsAvailable(product *Product) bool {
	inv.mu.RLock()
	defer inv.mu.RUnlock()
	return inv.products[product] > 0
}

// GetAllProducts 获取所有产品及其数量
func (inv *Inventory) GetAllProducts() map[*Product]int {
	inv.mu.RLock()
	defer inv.mu.RUnlock()

	// 返回副本以避免并发问题
	result := make(map[*Product]int)
	for p, q := range inv.products {
		result[p] = q
	}
	return result
}
