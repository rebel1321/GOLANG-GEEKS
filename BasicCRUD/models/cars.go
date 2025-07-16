package models

import (
	"car/config"
	"fmt"
)

var Cars = make(map[int]Car)

type Car struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Year  int     `json:"year"`
	Price float64 `json:"price"`
}

func (c *Car) Insert() {
	query := `INSERT INTO cars (name, model, brand, year, price) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	if err := config.DB.QueryRow(query, c.Name, c.Model, c.Brand, c.Year, c.Price).Scan(&c.Id); err != nil {
		fmt.Errorf("Error inserting car: %v\n", err)
	
	}

}
func (c *Car) Get() error {
	query := `SELECT id, name, model, brand, year, price FROM cars WHERE id = $1`
	err := config.DB.QueryRow(query, c.Id).Scan(&c.Id, &c.Name, &c.Model, &c.Brand, &c.Year, &c.Price)
	if err != nil {
		return fmt.Errorf("car not found: %v", err)
	}
	return nil
}


func (c *Car) Delete() error {
	query := `DELETE FROM cars WHERE id = $1`
	result, err := config.DB.Exec(query, c.Id)
	if err != nil {
		return fmt.Errorf("error deleting car: %v", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("car with id %d not found", c.Id)
	}
	return nil
}
func (c *Car) Update() {
	query := `UPDATE cars SET name = $1, model = $2, brand = $3, year = $4, price = $5 WHERE id = $6`
	if _, err := config.DB.Exec(query, c.Name, c.Model, c.Brand, c.Year, c.Price, c.Id); err != nil {
		fmt.Errorf("Error updating car: %v\n", err)

	}

}