# ATM系统 - Go实现

这是一个完整的ATM（自动取款机）系统的Go语言实现，满足所有设计要求。

## 功能特性

1. **基本操作**：余额查询、取款、存款
2. **用户认证**：使用卡号和PIN码进行身份验证
3. **银行服务**：与银行后端系统交互，验证账户和处理交易
4. **现金分发器**：管理ATM机的现金容量
5. **并发安全**：使用互斥锁和sync.Map确保数据一致性
6. **用户友好**：清晰的错误处理和反馈

## 项目结构

```
atm/
├── account.go              # 账户类
├── atm.go                  # ATM主类
├── banking_service.go      # 银行服务
├── card.go                 # 银行卡类
├── cash_dispenser.go       # 现金分发器
├── deposit_transaction.go  # 存款交易
├── errors.go               # 错误定义
├── transaction.go          # 交易接口
├── withdrawal_transaction.go # 取款交易
├── atm_driver.go           # 演示程序
├── atm_test.go             # 测试用例
└── README.md               # 说明文档
```

## 核心组件

### 1. Card（银行卡）
- 存储卡号、PIN码和关联的账户号
- 提供PIN验证功能

### 2. Account（账户）
- 管理账户余额
- 支持线程安全的存款（Credit）和取款（Debit）操作
- 使用互斥锁确保并发安全

### 3. Transaction（交易）
- 交易接口定义
- WithdrawalTransaction：取款交易
- DepositTransaction：存款交易

### 4. BankingService（银行服务）
- 管理账户和银行卡
- 处理用户认证
- 执行交易
- 使用sync.Map确保线程安全

### 5. CashDispenser（现金分发器）
- 管理ATM机的现金容量
- 分发和接收现金
- 线程安全操作

### 6. ATM（ATM机）
- 用户认证
- 余额查询
- 取款操作
- 存款操作
- 生成唯一交易ID

## 运行演示

```bash
# 运行演示程序
go run atm_driver.go
```

## 运行测试

```bash
# 运行所有测试
go test -v

# 运行特定测试
go test -v -run TestGetBalance

# 查看测试覆盖率
go test -cover

# 生成详细的覆盖率报告
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 使用示例

```go
package main

import (
    "fmt"
    "atm"
)

func main() {
    // 初始化银行服务
    bankingService := atm.NewBankingService()

    // 创建账户
    account := atm.NewAccount("ACC001", 1000.0)
    bankingService.AddAccount(account)

    // 创建银行卡
    card := atm.NewCard("CARD001", "1234", "ACC001")
    bankingService.AddCard(card)

    // 初始化ATM
    cashDispenser := atm.NewCashDispenser(10000)
    atmMachine := atm.NewATM(bankingService, cashDispenser)

    // 查询余额
    balance, err := atmMachine.GetBalance("CARD001", "1234")
    if err != nil {
        fmt.Printf("错误: %v\n", err)
    } else {
        fmt.Printf("当前余额: %.2f\n", balance)
    }

    // 取款
    err = atmMachine.WithdrawCash("CARD001", "1234", 200.0)
    if err != nil {
        fmt.Printf("取款失败: %v\n", err)
    } else {
        fmt.Println("取款成功")
    }

    // 存款
    err = atmMachine.DepositCash("CARD001", "1234", 300.0)
    if err != nil {
        fmt.Printf("存款失败: %v\n", err)
    } else {
        fmt.Println("存款成功")
    }
}
```

## 错误处理

系统定义了以下错误类型：

- `ErrInsufficientFunds`: 余额不足
- `ErrInsufficientCashInATM`: ATM现金不足
- `ErrInvalidAmount`: 无效金额（负数或零）
- `ErrInvalidPIN`: PIN码错误
- `ErrAccountNotFound`: 账户不存在
- `ErrCardNotFound`: 银行卡不存在

## 并发安全

系统使用以下机制确保并发安全：

1. **Account**: 使用`sync.Mutex`保护余额操作
2. **BankingService**: 使用`sync.Map`存储账户和卡片
3. **CashDispenser**: 使用`sync.Mutex`保护现金操作
4. **ATM**: 使用`atomic`包生成唯一交易ID

## 测试覆盖

测试用例涵盖：

- ✅ ATM创建
- ✅ 用户认证（正确/错误PIN，不存在的卡）
- ✅ 余额查询
- ✅ 取款（成功、余额不足、ATM现金不足、无效金额）
- ✅ 存款（成功、无效金额）
- ✅ 并发取款
- ✅ 并发存款和取款
- ✅ 交易ID生成
- ✅ 多账户操作
- ✅ 各个组件的单元测试

## 设计模式

- **接口模式**: Transaction接口支持不同类型的交易
- **组合模式**: ATM组合了BankingService和CashDispenser
- **单一职责原则**: 每个类都有明确的职责
- **线程安全**: 所有共享状态都有适当的同步机制

## 许可证

本项目仅供学习和参考使用。

