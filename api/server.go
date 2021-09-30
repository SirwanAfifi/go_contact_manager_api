package api

import (
	"net/http"

	"github.com/SirwanAfifi/go_server/database"
	"github.com/SirwanAfifi/go_server/services"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)


type Server struct {
	*mux.Router
	Db *gorm.DB
	AllUsers http.Handler
	NewUser http.Handler
	DeleteUser http.Handler
	UpdateUser http.Handler
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.Db = database.InitDb()
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/api/users", services.AllUsers).Methods("GET")
	s.HandleFunc("/api/users", services.NewUser).Methods("POST")
	s.HandleFunc("/api/users", services.DeleteUser).Methods("DELETE")
	s.HandleFunc("/api/users", services.UpdateUser).Methods("PUT")
}
