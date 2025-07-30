package decorator

type Coffee interface {
    Cost() float64
    Ingredients() string
}

type BasicCoffee struct{}

func (c *BasicCoffee) Cost() float64 {
   return 2.00 // Base cost of the coffee
}

func (c *BasicCoffee) Ingredients() string {
    return "Coffee"
}
type WithMilk struct {
    Coffee Coffee
}

func (m *WithMilk) Cost() float64 {
    return m.Coffee.Cost() + 0.50 // Adding cost of milk
}

func (m *WithMilk) Ingredients() string {
    return m.Coffee.Ingredients() + ", Milk"
}

type WithSugar struct {
    Coffee Coffee
}

func (s *WithSugar) Cost() float64 {
    return s.Coffee.Cost() + 0.25 // Adding cost of sugar
}

func (s *WithSugar) Ingredients() string {
    return s.Coffee.Ingredients() + ", Sugar"
}