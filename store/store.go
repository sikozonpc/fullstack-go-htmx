package store

import (
	"database/sql"

	"github.com/sikozonpc/fullstackgo/types"
)

type Storage struct {
	db *sql.DB
}

type Store interface {
	CreateCar(car *types.Car) (*types.Car, error)
	GetCars() ([]types.Car, error)
	DeleteCar(id string) error
	FindCarsByNameMakeOrBrand(search string) ([]types.Car, error)
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) DeleteCar(id string) error {
	_, err := s.db.Exec("DELETE FROM cars WHERE id = ?", id)
	return err
}

func (s *Storage) CreateCar(c *types.Car) (*types.Car, error) {
	row, err := s.db.Exec("INSERT INTO cars (brand, make, model, year, imageURL) VALUES (?, ?, ?, ?, ?)", c.Brand, c.Make, c.Model, c.Year, c.ImageURL)
	if err != nil {
		return nil, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}
	c.ID = int(id)

	return c, nil
}

func (s *Storage) GetCars() ([]types.Car, error) {
	rows, err := s.db.Query("SELECT * FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []types.Car
	for rows.Next() {
		car, err := scanCar(rows)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (s *Storage) FindCarsByNameMakeOrBrand(search string) ([]types.Car, error) {
	rows, err := s.db.Query("SELECT * FROM cars WHERE brand LIKE ? OR model LIKE ? OR make LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []types.Car
	for rows.Next() {
		car, err := scanCar(rows)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func scanCar(row *sql.Rows) (types.Car, error) {
	var car types.Car
	err := row.Scan(&car.ID, &car.Brand, &car.Make, &car.Model, &car.Year, &car.ImageURL, &car.CreatedAt)
	return car, err
}
