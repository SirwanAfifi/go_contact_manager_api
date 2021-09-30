package main

import (
	"net/http"

	"github.com/SirwanAfifi/go_server/api"
)

func main()  {
 	server := api.NewServer()
	http.ListenAndServe(":8080", server)
}