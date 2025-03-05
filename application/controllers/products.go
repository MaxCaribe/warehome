package controllers

import "net/http"

type Products struct {
}

func NewProducts() *Products {
	return &Products{}
}

func (*Products) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO WORLD"))
}
