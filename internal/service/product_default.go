package service

import "goweb/internal"

func NewProductDefaultService(rp internal.ProductRepository) *ProductDefaultService {
	return &ProductDefaultService{
		rp: rp,
	}
}

// Es un struct que representa la implemetanción de un servicio
type ProductDefaultService struct {
	// rp es un repositorio de productos
	rp internal.ProductRepository
}

func (p *ProductDefaultService) Save(product *internal.Product) (err error) {
	// Validar el producto, no tenemos restricción
	// más que el code value, así que no va nada

	// tengo que chequear que la fecha tenga un formato

	err = p.rp.Save(product)
	if err != nil {
		return
	}
	return
}

func (p *ProductDefaultService) GetAll() (slice []internal.Product, err error) {

	return p.rp.GetAll()
}

func (p *ProductDefaultService) GetByID(id int) (product internal.Product, err error) {

	product, err = p.rp.GetByID(id)
	if err != nil {
		return
	}
	return
}

func (p *ProductDefaultService) GetByPriceRange(price float64) (slice []internal.Product, err error) {

	slice, err = p.rp.GetByPriceRange(price)
	if err != nil {
		return
	}
	return
}

func (p *ProductDefaultService) Update(product *internal.Product) (err error) {
	// Validar el producto, no tenemos restricción
	// más que el code value, así que no va nada

	// tengo que chequear que la fecha tenga un formato

	err = p.rp.Update(product)
	if err != nil {
		return
	}
	return
}

func (p *ProductDefaultService) UpdatePartially(product *internal.Product) (err error) {
	// Validar el producto, no tenemos restricción
	// más que el code value, así que no va nada

	// tengo que chequear que la fecha tenga un formato

	err = p.rp.UpdatePartially(product)
	if err != nil {
		return
	}
	return
}

func (p *ProductDefaultService) Delete(id int) (err error) {
	// Validar el producto, no tenemos restricción
	// más que el code value, así que no va nada

	// tengo que chequear que la fecha tenga un formato

	err = p.rp.Delete(id)
	if err != nil {
		return
	}
	return
}
