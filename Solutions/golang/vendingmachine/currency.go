package main

// Coin 代表硬币面额
type Coin int

const (
	Penny      Coin = 1   // 1分
	Nickel     Coin = 5   // 5分
	Dime       Coin = 10  // 10分 (1角)
	Quarter    Coin = 25  // 25分
	HalfDollar Coin = 50  // 50分 (半美元)
	Dollar     Coin = 100 // 100分 (1美元)
)

// Value 返回硬币的价值（以元为单位）
func (c Coin) Value() float64 {
	return float64(c) / 100.0
}

// Note 代表纸币面额
type Note int

const (
	One     Note = 100   // 1元
	Five    Note = 500   // 5元
	Ten     Note = 1000  // 10元
	Twenty  Note = 2000  // 20元
	Fifty   Note = 5000  // 50元
	Hundred Note = 10000 // 100元
)

// Value 返回纸币的价值（以元为单位）
func (n Note) Value() float64 {
	return float64(n) / 100.0
}
