package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"felipe.com/cadastroUsuario/src/model"

	"gopkg.in/mgo.v2"
)

// CustomerController manipulates de several useful operation
// over database
type CustomerController struct {
	Conn *mgo.Session
}

// Create inserts the new customer register
func (c CustomerController) Create(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer
	body, err := ioutil.ReadAll(r.Body)

	// Verfify if exists reader error
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}

	// Transforms the json of body request into object
	json.Unmarshal(body, &customer)
	err = model.CreateCustomer(c.Conn, customer)

	// Case something else wrong on customer creation
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}

	json, _ := json.Marshal(customer)
	w.Write([]byte(json))
}

func (c CustomerController) Read(w http.ResponseWriter, r *http.Request) {
	customers, err := model.ReadCustomer(c.Conn)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "Um erro ocorreu")
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
