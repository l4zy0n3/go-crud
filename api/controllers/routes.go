package controllers

import "github.com/l4zy0n3/go-crud/api/middlewares"

func (s *Server) initializeRoutes() {

	//Book routes
	s.Router.HandleFunc("/books", middlewares.SetMiddlewareJSON(s.GetBooks)).Methods("GET")
}
