package handler

import (
	"encoding/json"
	"errors"
	"goweb/internal"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func NewProductDefaultHandler(sv internal.ProductService) *ProductDefaultHandler {
	return &ProductDefaultHandler{
		sv: sv,
	}
}

// Es una implemetación con handlers para el storage del producto
type ProductDefaultHandler struct {
	// sv is a movie service
	sv internal.ProductService
}

type ProductBodyJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (h *ProductDefaultHandler) Create() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Leyendo el JSON que viene en el body
		// tengo que chequear si está vacío algún campo
		var body ProductBodyJSON
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid body"))
			return
		}

		newProduct := internal.Product{
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		if err := h.sv.Save(&newProduct); err != nil {
			switch {
			case errors.Is(err, internal.ErrCodeValueAlreadyExists):
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
			}
			return

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newProduct)
	}
}

func (h *ProductDefaultHandler) GetAll() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		slice, err := h.sv.GetAll()
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		json.NewEncoder(w).Encode(slice)

	}
}

func (h *ProductDefaultHandler) GetByID() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		idUrl := chi.URLParam(r, "id")
		idUrlInt, err := strconv.Atoi(idUrl)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid id"))
			return
		}

		product, err := h.sv.GetByID(idUrlInt)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

func (h *ProductDefaultHandler) GetByPriceRange() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		priceUrl := r.URL.Query().Get("price")
		priceUrlFloat, err := strconv.ParseFloat(priceUrl, 64)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid price"))
			return
		}

		slice, err := h.sv.GetByPriceRange(priceUrlFloat)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(slice)
	}
}

func (h *ProductDefaultHandler) Update() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		idUrl := chi.URLParam(r, "id")
		idUrlInt, err := strconv.Atoi(idUrl)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid id"))
			return
		}

		// Leyendo el JSON que viene en el body
		// tengo que chequear si está vacío algún campo
		var body ProductBodyJSON
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid body"))
			return
		}

		newProduct := internal.Product{
			Id:          idUrlInt,
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		if err := h.sv.Update(&newProduct); err != nil {
			switch {
			case errors.Is(err, internal.ErrCodeValueAlreadyExists):
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
			}
			return

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newProduct)
	}
}

func (h *ProductDefaultHandler) UpdatePartially() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		idUrl := chi.URLParam(r, "id")
		idUrlInt, err := strconv.Atoi(idUrl)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid id"))
			return
		}
		// Traigo el producto completo para después pisarle los campos
		product, err := h.sv.GetByID(idUrlInt)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// Leyendo el JSON que viene en el body
		// tengo que chequear si está vacío algún campo
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid body"))
			return
		}

		if err := h.sv.UpdatePartially(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrCodeValueAlreadyExists):
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
			}
			return

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

func (h *ProductDefaultHandler) Delete() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		idUrl := chi.URLParam(r, "id")
		idUrlInt, err := strconv.Atoi(idUrl)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid id"))
			return
		}

		if err := h.sv.Delete(idUrlInt); err != nil {
			switch {
			case errors.Is(err, internal.ErrCodeValueAlreadyExists):
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
			}
			return

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product deleted"))
	}
}
