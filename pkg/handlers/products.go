package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ThorstenHans/bridge-to-kubernetes-demo/pkg/store"
)

type Products struct {
	Log   *log.Logger
	Store *store.Store
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.serveAllProducts(w)
	} else if r.Method == http.MethodPost {
		p.addNewProduct(w, r)
	}
}

func (p *Products) addNewProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	model := &CreateProductModel{}
	err := decoder.Decode(model)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product := p.Store.CreateProduct(model.Name, model.Price, model.IsInStock)
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("location", fmt.Sprintf("%s://%s%s/%v", r.URL.Scheme, r.URL.Host, r.URL.Path, product.Id))
	encoder := json.NewEncoder(w)
	encoder.Encode(product)
}

func (p *Products) serveAllProducts(w http.ResponseWriter) {
	encoder := json.NewEncoder(w)
	encoder.Encode(p.Store.GetAll())
}
