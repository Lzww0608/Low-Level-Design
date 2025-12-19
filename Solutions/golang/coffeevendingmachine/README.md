# 咖啡售货机 (Coffee Vending Machine)

这是一个完整的咖啡售货机实现，使用 Go 语言编写，实现了所有需求功能并确保线程安全。

## 功能特性

✅ **多种咖啡类型支持**
- Espresso（浓缩咖啡）
- Cappuccino（卡布奇诺）
- Latte（拿铁）

✅ **价格和配方管理**
- 每种咖啡都有独立的价格
- 每种咖啡都有特定的原料配方

✅ **菜单显示**
- 显示所有可用的咖啡选项
- 显示每种咖啡的价格

✅ **支付系统**
- 支持用户支付
- 自动计算找零
- 支付不足时返回错误

✅ **库存管理**
- 跟踪所有原料的库存
- 自动检查原料是否充足
- 购买后自动扣除原料
- 低库存预警功能
- 支持原料补充

✅ **并发安全**
- 使用互斥锁确保线程安全
- 支持多个用户同时购买
- 库存操作完全线程安全

## 项目结构

```
coffeevendingmachine/
├── coffee.go           # Coffee 和 Ingredient 结构定义
├── coffee_type.go      # CoffeeType 枚举类型
├── inventory.go        # 库存管理（线程安全）
├── machine.go          # 咖啡机主类
├── payment.go          # 支付处理
├── main.go            # 演示程序
├── machine_test.go    # 完整的测试用例
├── cmd/
│   └── main.go        # 可执行程序入口
└── README.md          # 本文档
```

## 核心类设计

### 1. Coffee (咖啡)
```go
type Coffee struct {
    Type   CoffeeType    // 咖啡类型
    Name   string        // 咖啡名称
    Price  float64       // 价格
    Recipe []Ingredient  // 配方（原料列表）
}
```

### 2. Ingredient (原料)
```go
type Ingredient struct {
    Name     string  // 原料名称
    Quantity int     // 需要的数量
}
```

### 3. Inventory (库存)
```go
type Inventory struct {
    ingredients       map[string]int  // 原料名称 -> 数量
    mu                sync.RWMutex    // 读写锁
    lowStockThreshold int             // 低库存阈值
}
```

**主要功能：**
- `AddIngredient()` - 添加或更新原料
- `HasEnoughIngredients()` - 检查原料是否充足
- `DeductIngredients()` - 扣除原料
- `GetLowStockIngredients()` - 获取低库存原料列表
- 所有操作都是线程安全的

### 4. Payment (支付)
```go
type Payment struct {
    Amount float64  // 支付金额
}
```

**主要功能：**
- `ProcessPayment()` - 处理支付并返回找零

### 5. Machine (咖啡机)
```go
type Machine struct {
    coffeeMenu []Coffee     // 咖啡菜单
    inventory  *Inventory   // 库存管理
    mu         sync.Mutex   // 机器操作锁
}
```

**主要功能：**
- `InitializeMenu()` - 初始化咖啡菜单
- `InitializeInventory()` - 初始化库存
- `DisplayMenu()` - 显示菜单
- `SelectAndDispenseCoffee()` - 选择并分配咖啡（核心功能）
- `CheckInventoryStatus()` - 检查库存状态
- `GetLowStockAlert()` - 获取低库存警告
- `RefillIngredient()` - 补充原料

## 使用示例

```go
// 1. 创建并初始化咖啡机
machine := NewMachine()
machine.InitializeMenu()
machine.InitializeInventory()

// 2. 显示菜单
fmt.Println(machine.DisplayMenu())

// 3. 用户购买咖啡
payment := NewPayment(10.0)  // 用户支付 $10
change, err := machine.SelectAndDispenseCoffee(ESPRESSO, payment)
if err != nil {
    fmt.Printf("购买失败: %v\n", err)
} else {
    fmt.Printf("购买成功! 找零: $%.2f\n", change)
}

// 4. 检查库存状态
inventory := machine.CheckInventoryStatus()
for name, quantity := range inventory {
    fmt.Printf("%s: %d\n", name, quantity)
}

// 5. 检查低库存警告
lowStock := machine.GetLowStockAlert()
if len(lowStock) > 0 {
    fmt.Println("低库存警告:", lowStock)
}

// 6. 补充原料
machine.RefillIngredient("Coffee Beans", 50)
```

## 运行测试

### 运行所有测试
```bash
go test -v
```

### 运行性能测试
```bash
go test -bench=. -benchmem
```

### 运行特定测试
```bash
go test -run TestConcurrentOrders -v
```

# 推荐的完整测试命令（本地或 CI）
```bash
go test -v -race -cover -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 测试覆盖

测试用例涵盖了以下场景：

1. ✅ **机器初始化测试** - 验证机器和菜单正确初始化
2. ✅ **菜单显示测试** - 验证菜单正确显示
3. ✅ **添加/删除咖啡测试** - 验证咖啡管理功能
4. ✅ **支付处理测试** - 验证不同支付场景（精确支付、超额支付、支付不足）
5. ✅ **库存管理测试** - 验证库存添加、检查、扣除功能
6. ✅ **低库存警告测试** - 验证低库存预警功能
7. ✅ **咖啡购买测试** - 验证完整的购买流程
8. ✅ **支付不足测试** - 验证支付不足时的错误处理
9. ✅ **原料不足测试** - 验证原料不足时的错误处理
10. ✅ **并发订单测试** - 验证50个并发用户同时购买
11. ✅ **补充原料测试** - 验证原料补充功能
12. ✅ **获取咖啡测试** - 验证咖啡查询功能
13. ✅ **库存线程安全测试** - 验证100个并发goroutine的库存操作
14. ✅ **性能基准测试** - 测试购买操作的性能

## 运行演示程序

```bash
# 方式1: 运行 Demo 函数
go test -run TestDemo

# 方式2: 编译并运行
go build -o coffee-machine ./cmd
./coffee-machine
```

## 线程安全保证

### 机器级别
- 使用 `sync.Mutex` 保护咖啡机的所有操作
- 确保同一时刻只有一个用户操作机器

### 库存级别
- 使用 `sync.RWMutex` 保护库存操作
- 读操作使用读锁（支持并发读）
- 写操作使用写锁（独占访问）
- 扣除原料时先检查再扣除，保证原子性

## 设计模式

1. **单例模式的变体** - Machine 类管理整个售货机
2. **策略模式** - 不同的咖啡类型有不同的配方
3. **观察者模式的简化版** - 低库存通知

## 扩展性

系统设计具有良好的扩展性：

1. **添加新咖啡类型**：
   - 在 `coffee_type.go` 中添加新类型
   - 在 `InitializeMenu()` 中添加新咖啡

2. **添加新原料**：
   - 直接通过 `AddIngredient()` 添加

3. **自定义低库存阈值**：
   - 使用 `SetLowStockThreshold()` 设置

4. **支持更复杂的支付方式**：
   - 扩展 `Payment` 类

## 依赖

无外部依赖，仅使用 Go 标准库：
- `sync` - 并发控制
- `fmt` - 格式化输出
- `testing` - 测试框架

