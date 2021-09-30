package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SirwanAfifi/go_server/models"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	
    var users []models.User

		for i := 0; i < 100; i++ {
			users = append(users, models.User {
				Name: fmt.Sprintf("User FirstName %d", i), 
				LastName: fmt.Sprintf("User LastName %d", i), 
				Email: fmt.Sprintf("user%d@example.com", i) ,
				PhoneNumber: fmt.Sprintf("%d%d-%d%d", i, i + 21, i * 2, i + 4) })
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	
    fmt.Fprintf(w, "New User Endpoint Hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Update User Endpoint Hit")
}