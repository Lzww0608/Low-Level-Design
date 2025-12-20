package atm

import "fmt"

func RunATMDemo() {
	// 初始化银行服务
	bankingService := NewBankingService()

	// 创建账户
	account1 := NewAccount("ACC001", 1000.0)
	account2 := NewAccount("ACC002", 500.0)
	bankingService.AddAccount(account1)
	bankingService.AddAccount(account2)

	// 创建卡片
	card1 := NewCard("CARD001", "1234", "ACC001")
	card2 := NewCard("CARD002", "5678", "ACC002")
	bankingService.AddCard(card1)
	bankingService.AddCard(card2)

	// 初始化ATM，现金容量为10000
	cashDispenser := NewCashDispenser(10000)
	atm := NewATM(bankingService, cashDispenser)

	fmt.Println("=== ATM系统演示 ===")
	fmt.Println()

	// 场景1：查询余额
	fmt.Println("场景1：查询余额")
	balance, err := atm.GetBalance("CARD001", "1234")
	if err != nil {
		fmt.Printf("查询余额失败: %v\n", err)
	} else {
		fmt.Printf("账户ACC001余额: %.2f\n", balance)
	}
	fmt.Println()

	// 场景2：取款
	fmt.Println("场景2：取款200元")
	err = atm.WithdrawCash("CARD001", "1234", 200.0)
	if err != nil {
		fmt.Printf("取款失败: %v\n", err)
	} else {
		fmt.Println("取款成功")
		balance, _ := atm.GetBalance("CARD001", "1234")
		fmt.Printf("当前余额: %.2f\n", balance)
	}
	fmt.Println()

	// 场景3：存款
	fmt.Println("场景3：存款300元")
	err = atm.DepositCash("CARD001", "1234", 300.0)
	if err != nil {
		fmt.Printf("存款失败: %v\n", err)
	} else {
		fmt.Println("存款成功")
		balance, _ := atm.GetBalance("CARD001", "1234")
		fmt.Printf("当前余额: %.2f\n", balance)
	}
	fmt.Println()

	// 场景4：PIN错误
	fmt.Println("场景4：使用错误的PIN")
	_, err = atm.GetBalance("CARD001", "0000")
	if err != nil {
		fmt.Printf("操作失败: %v\n", err)
	}
	fmt.Println()

	// 场景5：余额不足
	fmt.Println("场景5：尝试取款超过余额")
	err = atm.WithdrawCash("CARD002", "5678", 1000.0)
	if err != nil {
		fmt.Printf("取款失败: %v\n", err)
	}
	fmt.Println()

	// 场景6：ATM现金不足
	fmt.Println("场景6：ATM现金不足（尝试取款11000元）")
	err = atm.WithdrawCash("CARD001", "1234", 11000.0)
	if err != nil {
		fmt.Printf("取款失败: %v\n", err)
	}
	fmt.Println()

	// 场景7：无效金额
	fmt.Println("场景7：尝试取款负数金额")
	err = atm.WithdrawCash("CARD001", "1234", -100.0)
	if err != nil {
		fmt.Printf("取款失败: %v\n", err)
	}
	fmt.Println()

	fmt.Printf("ATM剩余现金: %d\n", cashDispenser.GetAvailableCash())
}
