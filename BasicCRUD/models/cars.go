package models

import (
	"car/config"
	"fmt"

	"gorm.io/gorm"
)

var Cars = make(map[int]Car)

type Car struct {
	Id    int     `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Year  int     `json:"year"`
	Price float64 `json:"price"`
}

func (c *Car) Insert() {

	if err := config.DB.Create(&c).Error; err != nil {
		fmt.Errorf("Error inserting car: %v\n", err)
	}
}
func (c *Car) Get() error {

	if err := config.DB.First(c, c.Id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("Error getting car: %v\n", err)
			return err 
		}
	}
	return nil
}


func (c *Car) Delete() error {
	if err := config.DB.Delete(c).Error; err != nil {
		fmt.Printf("Error deleting car: %v\n", err)
		return err
	}
	return nil
}
func (c *Car) Update() {
	if err := config.DB.Save(c).Error; err != nil {
		fmt.Printf("Error updating car: %v\n", err)
	}

}