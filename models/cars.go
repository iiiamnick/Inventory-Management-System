package models

import (
	"database/sql"
	"fmt"

	"github.com/iiiamnick/Inventory-Management-System/config"
)

var Cars = make(map[int]Car)

type Car struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Year  int     `json:"year"`
	Price float64 `json:"price"`
}

func (c *Car) Insert() {
	query := `INSERT INTO cars (name, model, brand, year, price) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := config.DB.QueryRow(query, c.Name, c.Model, c.Brand, c.Year, c.Price).Scan(&c.ID)
	if err != nil {
		fmt.Printf("Error inserting car: %v\n", err)

	}
}

func (c *Car) Update() {
	query := `UPDATE cars SET name=$1, model=$2, brand=$3, year=$4, price =$5 WHERE ID = $5`
	_, err := config.DB.Exec(query, c.Name, c.Model, c.Brand, c.Year, c.Price, c.ID)
	if err != nil {
		fmt.Printf("Error in updating car: %v", err)

	}
}

func (c *Car) GET() error {
	query := `SELECT name, model, brand, year, price FROM cars WHERE ID = $1`
	err := config.DB.QueryRow(query, c.ID).Scan(&c.Name, &c.Model, &c.Brand, &c.Year, &c.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Error getting  car: %v", err)
			return err
		}
	}
	return nil
}

func (c *Car) DELETE() error {
	query := `DELETE FROM cars  WHERE ID = $1`
	_, err := config.DB.Exec(query, c.ID)
	if err != nil {

		fmt.Printf("Error deleting car with id : %v, error %v\n", c.ID, err)
	}
	return nil
}
