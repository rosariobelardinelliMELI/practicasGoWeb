package internal

// var (
// 	ErrFieldRequired      = errors.New("field required")
// 	ErrFieldQuality       = errors.New("field quality")
// 	ErrMovieAlreadyExists = errors.New("movie already exists")
// )

// Interface que representa al servicio
type ProductService interface {
	Save(product *Product) (err error)
	GetAll() (slice []Product, err error)
	GetByID(id int) (product Product, err error)
	GetByPriceRange(price float64) (slice []Product, err error)
	Update(product *Product) (err error)
	UpdatePartially(product *Product) (err error)
	Delete(id int) (err error)
}
