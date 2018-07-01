package main

import (
	"log"
	"net/http"

	"felipe.com/cadastroUsuario/src/controller"

	"felipe.com/cadastroUsuario/src/model"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	conn, err := model.NewSession()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	r := mux.NewRouter()
	customerCtrl := controller.CustomerController{Conn: conn}

	r.HandleFunc("/customers/", customerCtrl.Create).Methods("POST")
	r.HandleFunc("/customers/", customerCtrl.Read).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"}),
		handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "HEAD", "TRACE", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r)))
}