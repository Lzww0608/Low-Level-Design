package coffeevendingmachine

type CoffeeType int

const (
	ESPRESSO CoffeeType = iota
	CAPPUCCINO
	LATTE
	AMERICANO
	MOCHA
	LATTE_MACCHIATO
	IRISH_COFFEE
)

func (c CoffeeType) String() string {
	return []string{"Espresso", "Cappuccino", "Latte", "Americano", "Mocha", "Latte Macchiato", "Irish Coffee"}[c]
}
