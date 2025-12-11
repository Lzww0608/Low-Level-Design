package main

// Product 代表售货机中的一个产品
type Product struct {
	Name  string
	Price float64
}

// NewProduct 创建一个新产品
func NewProduct(name string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}
