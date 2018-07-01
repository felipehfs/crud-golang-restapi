package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"felipe.com/cadastroUsuario/src/model"
	"github.com/gorilla/mux"

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

// Read retrieves all data of the database
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

// Delete remove the Customer on database
func (c CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	err := model.DeleteCustomer(c.Conn, id)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "%v", err)
	}
	fmt.Fprintln(w, "Success!")
}

// Search finds the customer by ID on the database
func (c CustomerController) Search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	customer, err := model.FindCustomer(c.Conn, id)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, "Um erro ocorreu!")
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func errorHandler(w http.ResponseWriter, err *error) {

}

// Update changes the data of customer in the server
func (c CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var customer model.Customer
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &customer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}

	err = model.UpdateCustomer(c.Conn, id, customer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}

	json.NewEncoder(w).Encode(customer)
}
