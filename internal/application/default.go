package application

import (
	"goweb/internal"
	"goweb/internal/handler"
	"goweb/internal/repository"
	"goweb/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewDefaultHTTP(address string) *DefaultHTTP {
	return &DefaultHTTP{
		addr: address,
	}
}

type DefaultHTTP struct {
	addr string
}

func (h *DefaultHTTP) Run() (err error) {
	// initialize dependencies
	// - repository
	var Slice map[int]internal.Product
	rp := repository.NewProductosMap(Slice, 0)
	// load repository
	rp.LoadProducts("../products.json")
	// - service
	sv := service.NewProductDefaultService(rp)
	// - handler
	hd := handler.NewProductDefaultHandler(sv)
	// - router
	rt := chi.NewRouter()
	//   endpoints
	// rt.Get("/movies/{id}", hd.GetByID())
	// rt.Post("/movies", hd.Create())
	// rt.Put("/movies/{id}", hd.Update())

	rt.Route("/products", func(rt chi.Router) {
		rt.Get("/", hd.GetAll())
		rt.Get("/{id}", hd.GetByID())
		rt.Post("/", hd.Create())
		rt.Get("/search", hd.GetByPriceRange())
		rt.Put("/{id}", hd.Update())
		rt.Patch("/{id}", hd.UpdatePartially())
		rt.Delete("/{id}", hd.Delete())
	})
	// run http server
	err = http.ListenAndServe(h.addr, rt)
	return
}
