package atm

import (
	"sync"
	"testing"
)

// 测试创建ATM
func TestNewATM(t *testing.T) {
	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(10000)
	atm := NewATM(bankingService, cashDispenser)

	if atm == nil {
		t.Fatal("ATM创建失败")
	}
	if atm.bankingService == nil {
		t.Error("BankingService未正确初始化")
	}
	if atm.cashDispenser == nil {
		t.Error("CashDispenser未正确初始化")
	}
}

// 测试用户认证
func TestAuthenticateUser(t *testing.T) {
	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(10000)
	atm := NewATM(bankingService, cashDispenser)

	// 创建测试账户和卡片
	account := NewAccount("ACC001", 1000.0)
	bankingService.AddAccount(account)
	card := NewCard("CARD001", "1234", "ACC001")
	bankingService.AddCard(card)

	// 测试正确的PIN
	validCard, err := atm.AuthenticateUser("CARD001", "1234")
	if err != nil {
		t.Errorf("有效PIN认证失败: %v", err)
	}
	if validCard == nil {
		t.Error("返回的卡片为nil")
	}

	// 测试错误的PIN
	_, err = atm.AuthenticateUser("CARD001", "0000")
	if err != ErrInvalidPIN {
		t.Errorf("期望错误 %v, 得到 %v", ErrInvalidPIN, err)
	}

	// 测试不存在的卡片
	_, err = atm.AuthenticateUser("CARD999", "1234")
	if err != ErrCardNotFound {
		t.Errorf("期望错误 %v, 得到 %v", ErrCardNotFound, err)
	}
}

// 测试查询余额
func TestGetBalance(t *testing.T) {
	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(10000)
	atm := NewATM(bankingService, cashDispenser)

	account := NewAccount("ACC001", 1000.0)
	bankingService.AddAccount(account)
	card := NewCard("CARD001", "1234", "ACC001")
	bankingService.AddCard(card)

	// 测试正确查询
	balance, err := atm.GetBalance("CARD001", "1234")
	if err != nil {
		t.Errorf("查询余额失败: %v", err)
	}
	if balance != 1000.0 {
		t.Errorf("期望余额 1000.0, 得到 %.2f", balance)
	}

	// 测试错误的PIN
	_, err = atm.GetBalance("CARD001", "0000")
	if err != ErrInvalidPIN {
		t.Errorf("期望错误 %v, 得到 %v", ErrInvalidPIN, err)
	}
}

// 测试取款
func TestWithdrawCash(t *testing.T) {
	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(10000)
	atm := NewATM(bankingService, cashDispenser)

	account := NewAccount("ACC001", 1000.0)
	bankingService.AddAccount(account)
	card := NewCard("CARD001", "1234", "ACC001")
	bankingService.AddCard(card)

	// 测试成功取款
	err := atm.WithdrawCash("CARD001", "1234", 200.0)
	if err != nil {
		t.Errorf("取款失败: %v", err)
	}

	// 验证余额
	balance := account.GetBalance()
	if balance != 800.0 {
		t.Errorf("期望余额 800.0, 得到 %.2f", balance)
	}

	// 验证ATM现金减少
	availableCash := cashDispenser.GetAvailableCash()
	if availableCash != 9800 {
		t.Errorf("期望ATM现金 9800, 得到 %d", availableCash)
	}

	// 测试余额不足
	err = atm.WithdrawCash("CARD001", "1234", 1000.0)
	if err != ErrInsufficientFunds {
		t.Errorf("期望错误 %v, 得到 %v", ErrInsufficientFunds, err)
	}

	// 测试负数金额
	err = atm.WithdrawCash("CARD001", "1234", -100.0)
	if err != ErrInvalidAmount {
		t.Errorf("期望错误 %v, 得到 %v", ErrInvalidAmount, err)
	}

	// 测试零金额
	err = atm.WithdrawCash("CARD001", "1234", 0)
	if err != ErrInvalidAmount {
		t.Errorf("期望错误 %v, 得到 %v", ErrInvalidAmount, err)
	}
}

// 测试ATM现金不足
func TestWithdrawCashInsufficientATMCash(t *testing.T) {
	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(100)
	atm := NewATM(bankingService, cashDispenser)

	account := NewAccount("ACC001", 1000.0)
	bankingService.AddAccount(account)
	card := NewCard("CARD001", "1234", "ACC001")
	bankingService.AddCard(card)

	// 尝试取款超过ATM现金
	err := atm.WithdrawCash("CARD001", "1234", 200.0)
	if err != ErrInsufficientCashInATM {
		t.Errorf("期望错误 %v, 得到 %v", ErrInsufficientCashInATM, err)
	}

	// 验证账户余额未改变
	balance := account.GetBalance()
	if balance != 1000.0 {
		t.Errorf("账户余额不应改变，期望 1000.0, 得到 %.2f", balance)
	}
}

// 测试存款
func TestDepositCash(t *testing.T) {
	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(10000)
	atm := NewATM(bankingService, cashDispenser)

	account := NewAccount("ACC001", 1000.0)
	bankingService.AddAccount(account)
	card := NewCard("CARD001", "1234", "ACC001")
	bankingService.AddCard(card)

	// 测试成功存款
	err := atm.DepositCash("CARD001", "1234", 300.0)
	if err != nil {
		t.Errorf("存款失败: %v", err)
	}

	// 验证余额
	balance := account.GetBalance()
	if balance != 1300.0 {
		t.Errorf("期望余额 1300.0, 得到 %.2f", balance)
	}

	// 验证ATM现金增加
	availableCash := cashDispenser.GetAvailableCash()
	if availableCash != 10300 {
		t.Errorf("期望ATM现金 10300, 得到 %d", availableCash)
	}

	// 测试负数金额
	err = atm.DepositCash("CARD001", "1234", -100.0)
	if err != ErrInvalidAmount {
		t.Errorf("期望错误 %v, 得到 %v", ErrInvalidAmount, err)
	}
}

// 测试并发取款
func TestConcurrentWithdrawals(t *testing.T) {
	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(10000)
	atm := NewATM(bankingService, cashDispenser)

	account := NewAccount("ACC001", 1000.0)
	bankingService.AddAccount(account)
	card := NewCard("CARD001", "1234", "ACC001")
	bankingService.AddCard(card)

	var wg sync.WaitGroup
	successCount := 0
	var mu sync.Mutex

	// 10个并发取款，每次100元
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := atm.WithdrawCash("CARD001", "1234", 100.0)
			if err == nil {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	// 只有前10次应该成功（1000 / 100 = 10）
	if successCount != 10 {
		t.Errorf("期望10次成功取款, 得到 %d", successCount)
	}

	// 验证最终余额
	balance := account.GetBalance()
	if balance != 0 {
		t.Errorf("期望余额 0, 得到 %.2f", balance)
	}
}

// 测试并发存款和取款
func TestConcurrentDepositAndWithdraw(t *testing.T) {
	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(10000)
	atm := NewATM(bankingService, cashDispenser)

	account := NewAccount("ACC001", 1000.0)
	bankingService.AddAccount(account)
	card := NewCard("CARD001", "1234", "ACC001")
	bankingService.AddCard(card)

	var wg sync.WaitGroup

	// 5个并发存款，每次100元
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = atm.DepositCash("CARD001", "1234", 100.0)
		}()
	}

	// 5个并发取款，每次100元
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = atm.WithdrawCash("CARD001", "1234", 100.0)
		}()
	}

	wg.Wait()

	// 验证最终余额应该是1000（存500，取500）
	balance := account.GetBalance()
	if balance != 1000.0 {
		t.Errorf("期望余额 1000.0, 得到 %.2f", balance)
	}
}

// 测试交易ID生成
func TestGenerateTransactionID(t *testing.T) {
	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(10000)
	atm := NewATM(bankingService, cashDispenser)

	id1 := atm.GenerateTransactionID()
	id2 := atm.GenerateTransactionID()

	if id1 == id2 {
		t.Error("交易ID应该是唯一的")
	}

	if id1 != "TXN-1" {
		t.Errorf("期望第一个交易ID为 TXN-1, 得到 %s", id1)
	}
	if id2 != "TXN-2" {
		t.Errorf("期望第二个交易ID为 TXN-2, 得到 %s", id2)
	}
}

// 测试多个账户
func TestMultipleAccounts(t *testing.T) {
	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(10000)
	atm := NewATM(bankingService, cashDispenser)

	// 创建两个账户
	account1 := NewAccount("ACC001", 1000.0)
	account2 := NewAccount("ACC002", 500.0)
	bankingService.AddAccount(account1)
	bankingService.AddAccount(account2)

	card1 := NewCard("CARD001", "1234", "ACC001")
	card2 := NewCard("CARD002", "5678", "ACC002")
	bankingService.AddCard(card1)
	bankingService.AddCard(card2)

	// 账户1取款
	err := atm.WithdrawCash("CARD001", "1234", 200.0)
	if err != nil {
		t.Errorf("账户1取款失败: %v", err)
	}

	// 账户2存款
	err = atm.DepositCash("CARD002", "5678", 100.0)
	if err != nil {
		t.Errorf("账户2存款失败: %v", err)
	}

	// 验证余额
	balance1, _ := atm.GetBalance("CARD001", "1234")
	balance2, _ := atm.GetBalance("CARD002", "5678")

	if balance1 != 800.0 {
		t.Errorf("账户1期望余额 800.0, 得到 %.2f", balance1)
	}
	if balance2 != 600.0 {
		t.Errorf("账户2期望余额 600.0, 得到 %.2f", balance2)
	}
}

// 测试Card类
func TestCard(t *testing.T) {
	card := NewCard("CARD001", "1234", "ACC001")

	if card.GetCardNumber() != "CARD001" {
		t.Errorf("期望卡号 CARD001, 得到 %s", card.GetCardNumber())
	}

	if card.GetAccountNumber() != "ACC001" {
		t.Errorf("期望账号 ACC001, 得到 %s", card.GetAccountNumber())
	}

	if !card.ValidatePIN("1234") {
		t.Error("有效PIN验证失败")
	}

	if card.ValidatePIN("0000") {
		t.Error("无效PIN验证应该失败")
	}
}

// 测试Account类
func TestAccount(t *testing.T) {
	account := NewAccount("ACC001", 1000.0)

	if account.GetAccountNumber() != "ACC001" {
		t.Errorf("期望账号 ACC001, 得到 %s", account.GetAccountNumber())
	}

	if account.GetBalance() != 1000.0 {
		t.Errorf("期望余额 1000.0, 得到 %.2f", account.GetBalance())
	}

	// 测试Credit
	err := account.Credit(200.0)
	if err != nil {
		t.Errorf("Credit失败: %v", err)
	}
	if account.GetBalance() != 1200.0 {
		t.Errorf("期望余额 1200.0, 得到 %.2f", account.GetBalance())
	}

	// 测试Debit
	err = account.Debit(300.0)
	if err != nil {
		t.Errorf("Debit失败: %v", err)
	}
	if account.GetBalance() != 900.0 {
		t.Errorf("期望余额 900.0, 得到 %.2f", account.GetBalance())
	}

	// 测试余额不足
	err = account.Debit(1000.0)
	if err != ErrInsufficientFunds {
		t.Errorf("期望错误 %v, 得到 %v", ErrInsufficientFunds, err)
	}
}

// 测试CashDispenser
func TestCashDispenser(t *testing.T) {
	dispenser := NewCashDispenser(1000)

	if dispenser.GetAvailableCash() != 1000 {
		t.Errorf("期望现金 1000, 得到 %d", dispenser.GetAvailableCash())
	}

	// 测试分发现金
	err := dispenser.DispenseCash(300)
	if err != nil {
		t.Errorf("分发现金失败: %v", err)
	}
	if dispenser.GetAvailableCash() != 700 {
		t.Errorf("期望现金 700, 得到 %d", dispenser.GetAvailableCash())
	}

	// 测试现金不足
	err = dispenser.DispenseCash(800)
	if err != ErrInsufficientCashInATM {
		t.Errorf("期望错误 %v, 得到 %v", ErrInsufficientCashInATM, err)
	}

	// 测试添加现金
	dispenser.AddCash(500)
	if dispenser.GetAvailableCash() != 1200 {
		t.Errorf("期望现金 1200, 得到 %d", dispenser.GetAvailableCash())
	}
}

// 测试BankingService
func TestBankingService(t *testing.T) {
	service := NewBankingService()

	account := NewAccount("ACC001", 1000.0)
	service.AddAccount(account)

	// 测试获取账户
	retrievedAccount, err := service.GetAccount("ACC001")
	if err != nil {
		t.Errorf("获取账户失败: %v", err)
	}
	if retrievedAccount.GetAccountNumber() != "ACC001" {
		t.Error("获取的账户不正确")
	}

	// 测试获取不存在的账户
	_, err = service.GetAccount("ACC999")
	if err != ErrAccountNotFound {
		t.Errorf("期望错误 %v, 得到 %v", ErrAccountNotFound, err)
	}

	// 测试卡片管理
	card := NewCard("CARD001", "1234", "ACC001")
	service.AddCard(card)

	retrievedCard, err := service.GetCard("CARD001")
	if err != nil {
		t.Errorf("获取卡片失败: %v", err)
	}
	if retrievedCard.GetCardNumber() != "CARD001" {
		t.Error("获取的卡片不正确")
	}

	// 测试卡片验证
	validCard, err := service.ValidateCard("CARD001", "1234")
	if err != nil {
		t.Errorf("卡片验证失败: %v", err)
	}
	if validCard == nil {
		t.Error("返回的卡片为nil")
	}

	// 测试错误PIN
	_, err = service.ValidateCard("CARD001", "0000")
	if err != ErrInvalidPIN {
		t.Errorf("期望错误 %v, 得到 %v", ErrInvalidPIN, err)
	}
}

