package atm

import "sync"

type CashDispenser struct {
	availableCash int
	mu            sync.Mutex
}

func NewCashDispenser(availableCash int) *CashDispenser {
	return &CashDispenser{
		availableCash: availableCash,
	}
}

func (c *CashDispenser) DispenseCash(amount int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if amount > c.availableCash {
		return ErrInsufficientCashInATM
	}
	c.availableCash -= amount
	return nil
}

func (c *CashDispenser) AddCash(amount int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.availableCash += amount
}

func (c *CashDispenser) GetAvailableCash() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.availableCash
}
