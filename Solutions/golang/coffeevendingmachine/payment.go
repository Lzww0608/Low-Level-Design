package coffeevendingmachine

import "fmt"

type Payment struct {
	Amount float64
}

func NewPayment(amount float64) *Payment {
	return &Payment{
		Amount: amount,
	}
}

// ProcessPayment 处理支付并返回找零
func (p *Payment) ProcessPayment(price float64) (float64, error) {
	if p.Amount < price {
		return 0, fmt.Errorf("insufficient payment: paid %.2f, required %.2f", p.Amount, price)
	}
	change := p.Amount - price
	return change, nil
}

// GetAmount 获取支付金额
func (p *Payment) GetAmount() float64 {
	return p.Amount
}
