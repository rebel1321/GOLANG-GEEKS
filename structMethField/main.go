package main

import "fmt"

type Car struct {
	Name  string
	Type  string
	Brand string
	Year  int
}

func (c Car) operational() {
	if c.Year < 2022 {
		fmt.Println("This car is operational.")
	} else {
		fmt.Println("This car is new and operational.")
	}
} 
func main() {
	c := Car{
		Name:  "Model S",
		Type:  "Sedan",
		Brand: "Tesla",
		Year:  2020,
	}
	fmt.Println("Car Details:")
	fmt.Println("Name:", c.Name)	
	fmt.Println("Type:", c.Type)
	fmt.Println("Brand:", c.Brand)
	fmt.Println("Year:", c.Year)

	c.operational()
}