package coffeevendingmachine

type Coffee struct {
	Type   CoffeeType
	Name   string
	Price  float64
	Recipe []Ingredient
}

type Ingredient struct {
	Name     string
	Quantity int
}

func NewCoffee(coffetype CoffeeType, name string, price float64, recipe []Ingredient) *Coffee {
	return &Coffee{
		Type:   coffetype,
		Name:   name,
		Price:  price,
		Recipe: recipe,
	}
}

func (c *Coffee) GetType() CoffeeType {
	return c.Type
}

func (c *Coffee) GetName() string {
	return c.Name
}

func (c *Coffee) GetPrice() float64 {
	return c.Price
}

func (c *Coffee) GetRecipe() []Ingredient {
	return c.Recipe
}

func (c *Coffee) SetPrice(price float64) {
	c.Price = price
}

func (c *Coffee) SetRecipe(recipe []Ingredient) {
	c.Recipe = recipe
}
