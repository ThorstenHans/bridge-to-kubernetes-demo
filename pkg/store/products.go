package store

import (
	"errors"
	"sort"
)

type Store struct {
	products []Product
}

func NewStore() *Store {
	return &Store{
		products: []Product{
			{Id: 1, Name: "Milk", Price: 0.99, IsInStock: true},
			{Id: 2, Name: "Butter", Price: 1.49, IsInStock: true},
			{Id: 3, Name: "Coke", Price: 1.19, IsInStock: false},
		},
	}
}

func (s *Store) CreateProduct(name string, price float32, inStock bool) *Product {

	sort.Slice(s.products, func(i, j int) bool { return s.products[i].Id < s.products[j].Id })
	nextId := s.products[len(s.products)-1].Id + 1
	p := &Product{
		Id:        nextId,
		Name:      name,
		Price:     price,
		IsInStock: inStock,
	}
	s.products = append(s.products, *p)
	return p
}

func (s *Store) GetAll() *[]Product {
	return &s.products
}

func (s *Store) GetById(id int) (*Product, error) {
	p, _ := s.findById(id)
	if p == nil {
		return nil, errors.New("Product not found")
	}
	return p, nil
}

func (s *Store) UpdateById(id int, name string, price float32, inStock bool) (*Product, error) {
	p, _ := s.findById(id)
	if p == nil {
		return nil, errors.New("Product not found")
	}
	p.Name = name
	p.IsInStock = inStock
	p.Price = price
	return p, nil
}

func (s *Store) DeleteById(id int) error {
	p, idx := s.findById(id)
	if p == nil {
		return errors.New("Product not found")
	}

	s.products = append(s.products[:idx], s.products[idx+1:]...)
	return nil
}

func (s *Store) findById(id int) (*Product, int) {
	for i, p := range s.products {
		if p.Id == id {
			return &s.products[i], i
		}
	}
	return nil, -1
}
