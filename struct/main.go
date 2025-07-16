package main

import "fmt"

func main() {
	type Person struct {
		Name string
		Age  int
	}
	var p Person
	p = Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Println("Name:", p.Name)
	p = Person{
		Name: "Bob",
		Age:  25,
	}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)


	type Employee struct {
		Position string
		Salary   int
		Person   // Embedded struct
	}

	var e Employee

	e = Employee{
		Position: "Developer",
		Salary:   70000,
		Person: Person{
			Name: "Charlie",
			Age:  28,
		},
	}
	fmt.Println("Position:", e.Position)
	fmt.Println("Salary:", e.Salary)
	fmt.Println("Name:", e.Person.Name)
	fmt.Println("Age:", e.Person.Age)

	s := struct {
		Model string
		Year  int
	}{
		Model: "Toyota",
		Year:  2020,
	}
	fmt.Println("Car Model:", s.Model)
	fmt.Println("Car Year:", s.Year)
}