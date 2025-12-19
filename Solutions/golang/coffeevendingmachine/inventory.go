package coffeevendingmachine

import "sync"

type Inventory struct {
	ingredients       map[string]int // 原料名称 -> 数量
	mu                sync.RWMutex   // 读写锁，确保线程安全
	lowStockThreshold int            // 低库存阈值
}

func NewInventory() *Inventory {
	return &Inventory{
		ingredients:       make(map[string]int),
		lowStockThreshold: 10,
	}
}

// AddIngredient 添加或更新原料数量
func (i *Inventory) AddIngredient(name string, quantity int) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if current, exists := i.ingredients[name]; exists {
		i.ingredients[name] = current + quantity
	} else {
		i.ingredients[name] = quantity
	}
}

// GetIngredientQuantity 获取原料数量
func (i *Inventory) GetIngredientQuantity(name string) int {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.ingredients[name]
}

// HasEnoughIngredients 检查是否有足够的原料
func (i *Inventory) HasEnoughIngredients(recipe []Ingredient) bool {
	i.mu.RLock()
	defer i.mu.RUnlock()
	for _, ingredient := range recipe {
		if i.ingredients[ingredient.Name] < ingredient.Quantity {
			return false
		}
	}
	return true
}

// DeductIngredients 扣除原料
func (i *Inventory) DeductIngredients(recipe []Ingredient) bool {
	i.mu.Lock()
	defer i.mu.Unlock()

	// 先检查是否有足够的原料
	for _, ingredient := range recipe {
		if i.ingredients[ingredient.Name] < ingredient.Quantity {
			return false
		}
	}

	// 扣除原料
	for _, ingredient := range recipe {
		i.ingredients[ingredient.Name] -= ingredient.Quantity
	}
	return true
}

// GetLowStockIngredients 获取低库存原料列表
func (i *Inventory) GetLowStockIngredients() []string {
	i.mu.RLock()
	defer i.mu.RUnlock()

	lowStock := make([]string, 0)
	for name, quantity := range i.ingredients {
		if quantity < i.lowStockThreshold {
			lowStock = append(lowStock, name)
		}
	}
	return lowStock
}

// SetLowStockThreshold 设置低库存阈值
func (i *Inventory) SetLowStockThreshold(threshold int) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.lowStockThreshold = threshold
}

// GetAllIngredients 获取所有原料（用于显示）
func (i *Inventory) GetAllIngredients() map[string]int {
	i.mu.RLock()
	defer i.mu.RUnlock()

	result := make(map[string]int)
	for k, v := range i.ingredients {
		result[k] = v
	}
	return result
}
