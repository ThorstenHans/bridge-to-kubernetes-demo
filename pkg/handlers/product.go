package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ThorstenHans/bridge-demo/pkg/store"
)

type Product struct {
	Log   *log.Logger
	Store *store.Store
}

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) != 4 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(segments[len(segments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		p.getProductById(id, w)
	case http.MethodPut:
		p.updateProductById(id, w, r)
	case http.MethodDelete:
		p.deleteProductById(id, w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Product) updateProductById(id int, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	model := &UpdateProductModel{}
	err := decoder.Decode(model)
	if err != nil {
		p.Log.Printf("error: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := p.Store.UpdateById(id, model.Name, model.Price, model.IsInStock)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(product)
}

func (p *Product) deleteProductById(id int, w http.ResponseWriter) {
	err := p.Store.DeleteById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (p *Product) getProductById(id int, w http.ResponseWriter) {

	product, err := p.Store.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(product)
}
