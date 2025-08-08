package models

import (
	"car/config"
	"fmt"

)

var Cars = make(map[int]Car)

type Car struct {
	Id    int     `json:"id" db:"id"`
	Name  string  `json:"name" db:"name"`
	Model string  `json:"model" db:"model"`
	Brand string  `json:"brand" db:"brand"`
	Year  int     `json:"year" db:"year"`
	Price float64 `json:"price" db:"price"`
}

func (c *Car) Insert() {
  query := `INSERT INTO cars (name, model, brand, year, price) VALUES (:name, :model, :brand, :year, :price) RETURNING id`
	row, err := config.DB.NamedQuery(query, c)
	if err != nil {
		fmt.Printf("Error inserting car: %v\n", err)
	}
	defer row.Close()
	if row.Next() {
		if err := row.Scan(&c.Id); err != nil {
			fmt.Printf("Error scanning car ID: %v\n", err)
		}
	}


}
func (c *Car) Get() error {
	query := `SELECT id, name, model, brand, year, price FROM cars WHERE id = $1`
	if err := config.DB.Get(c, query, c.Id); err != nil {
		fmt.Printf("Error getting car: %v\n", err)
		return err
	}
	return nil
}


func (c *Car) Delete() error {
	query := `DELETE FROM cars WHERE id = $1`
	if _, err := config.DB.Exec(query, c.Id); err != nil {
		fmt.Printf("Error deleting car: %v\n", err)
		return err
	}
	
	return nil
}
func (c *Car) Update() {
	query := `UPDATE cars SET name = :name, model = :model, brand = :brand, year = :year, price = :price WHERE id = :id`
	_, err := config.DB.NamedExec(query, c)
	if err != nil {
		fmt.Printf("Error updating car: %v\n", err)
	}
}