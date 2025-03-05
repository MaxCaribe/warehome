package application

import (
	"github.com/MaxCaribe/warehome/application/controllers"
	"net/http"
)

type Server struct {
	Mux      *http.ServeMux
	products *controllers.Products
}

func NewServer() *Server {
	return &Server{
		Mux:      http.NewServeMux(),
		products: controllers.NewProducts(),
	}
}

func (s *Server) Run() {
	s.Mux.HandleFunc("/products", s.products.List)
}
