package repository

import (
	"encoding/json"
	"fmt"
	"goweb/internal"
	"os"
)

func NewProductosMap(db map[int]internal.Product, lastId int) *ProductosMap {
	// default config / values
	// ...

	return &ProductosMap{
		db:     db,
		lastId: lastId,
	}
}

type ProductosMap struct {
	db     map[int]internal.Product
	lastId int
}

func (s *ProductosMap) LoadProducts(filepath string) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error al abrir el archivo JSON: ", err)
	}
	defer file.Close()

	err = cargarMap(s, file)
	if err != nil {
		fmt.Println("Error decodificando el archivo JSON: ", err)
	}
	return

}

func cargarMap(s *ProductosMap, file *os.File) error {

	decoder := json.NewDecoder(file)
	var productos []internal.Product

	err := decoder.Decode(&productos)
	if err != nil {
		return err
	}
	// Tengo que construir el mapa a partir del slice
	s.db = make(map[int]internal.Product)
	for _, p := range productos {
		s.db[p.Id] = p
	}

	s.lastId = len(s.db)
	return nil
}

func (s *ProductosMap) Save(product *internal.Product) (err error) {

	// Validar que no tenga el mismo code value
	for _, p := range s.db {
		if p.CodeValue == product.CodeValue {
			return internal.ErrCodeValueAlreadyExists
		}
	}

	// Agregar el producto al slice
	(*s).lastId++
	(*product).Id = (*s).lastId
	s.db[s.lastId] = *product
	//s.db = append(s.db, (*product))
	return
}

func (s *ProductosMap) GetAll() (slice []internal.Product, err error) {
	slice = make([]internal.Product, 0, len(s.db))
	for k, v := range s.db {
		// serialization
		p := internal.Product{
			Id:          k,
			Name:        v.Name,
			Quantity:    v.Quantity,
			CodeValue:   v.CodeValue,
			IsPublished: v.IsPublished,
			Expiration:  v.Expiration,
			Price:       v.Price,
		}
		slice = append(slice, p)
	}
	return slice, nil
}

func (s *ProductosMap) GetByID(id int) (product internal.Product, err error) {
	product, ok := s.db[id]
	if !ok {
		err = internal.ErrProductDoesntExists
	}
	return
}

func (s *ProductosMap) GetByPriceRange(price float64) (slice []internal.Product, err error) {
	for _, p := range s.db {
		if p.Price > price {
			slice = append(slice, p)
		}
	}
	if len(slice) == 0 {
		err = internal.ErrNoProductsWithThatPrice
	}
	return
}

func (s *ProductosMap) Update(product *internal.Product) (err error) {

	// Validar que no tenga el mismo code value
	for _, p := range s.db {
		if p.CodeValue == product.CodeValue {
			return internal.ErrCodeValueAlreadyExists
		}
	}

	// Agregar el producto al slice
	// (*s).lastId++
	// (*product).Id = (*s).lastId
	// s.db = append(s.db, (*product))
	s.db[product.Id] = *product
	return
}

func (s *ProductosMap) UpdatePartially(product *internal.Product) (err error) {

	// Agregar el producto al slice
	// (*s).lastId++
	// (*product).Id = (*s).lastId
	// s.db = append(s.db, (*product))
	s.db[product.Id] = *product
	return
}

func (s *ProductosMap) Delete(id int) (err error) {

	_, ok := s.db[id]
	if !ok {
		err = internal.ErrProductDoesntExists
	}
	delete(s.db, id)
	return
}
