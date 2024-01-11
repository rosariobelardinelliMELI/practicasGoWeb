package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var SliceProductos []Product

func main() {

	file, err := os.Open("./products.json")
	if err != nil {
		fmt.Println("Error al abrir el archivo JSON: ", err)
	}
	defer file.Close()

	err = cargarSlice(&SliceProductos, file)
	if err != nil {
		fmt.Println("Error decodificando el archivo JSON: ", err)
	}

	//fmt.Println(SliceProductos)

	router := chi.NewRouter()
	router.Route("/ping", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("pong"))
		})
	})
	router.Route("/products", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(SliceProductos)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			idUrl := chi.URLParam(r, "id")
			idUrlInt, _ := strconv.Atoi(idUrl)

			var product Product
			for _, p := range SliceProductos {
				if p.Id == idUrlInt {
					product = p
					break
				}
				//product = nil
				product.Id = 0
			}
			//fmt.Printf("%v", product)

			if product.Id == 0 {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte("Product not found"))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(product)
		})
		r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
			priceUrl := r.URL.Query().Get("price")
			priceUrlFloat, _ := strconv.ParseFloat(priceUrl, 64)

			var resultProducts []Product
			for _, p := range SliceProductos {
				if p.Price > priceUrlFloat {
					resultProducts = append(resultProducts, p)
				}
			}
			//fmt.Printf("%v", product)

			if len(resultProducts) == 0 {
				w.WriteHeader(http.StatusNotFound)
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte("Products not found"))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resultProducts)
		})
	})

	http.ListenAndServe(":8080", router)

}

func cargarSlice(s *[]Product, file *os.File) error {

	decoder := json.NewDecoder(file)
	err := decoder.Decode(s)
	if err != nil {
		return err
	}
	return nil
}
