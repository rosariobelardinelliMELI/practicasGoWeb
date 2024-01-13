package internal

import (
	"errors"
)

var (
	ErrCodeValueAlreadyExists  = errors.New("Code Value already exists, must be new")
	ErrProductDoesntExists     = errors.New("There isn't a product with that Id")
	ErrNoProductsWithThatPrice = errors.New("There aren't any products which price is higher than that")
)

// Es la interface que representa al repositorio
type ProductRepository interface {
	LoadProducts(filepath string) (err error)
	Save(product *Product) (err error)
	GetAll() (slice []Product, err error)
	GetByID(id int) (product Product, err error)
	GetByPriceRange(price float64) (slice []Product, err error)
	Update(product *Product) (err error)
	UpdatePartially(product *Product) (err error)
	Delete(id int) (err error)
}
