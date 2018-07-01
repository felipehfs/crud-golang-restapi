package test

import (
	"testing"

	"felipe.com/cadastroUsuario/src/model"
)

func TestCreateCustomer(t *testing.T) {
	con, _ := model.NewSession()
	c := model.Customer{}
	c.Name = "Marcos teste"
	c.Email = "marcosTeste@gmail.com"

	e := model.CreateCustomer(con, c)

	if e != nil {
		t.Error(e)
	}
	defer con.Close()
}

func TestReadCustomer(t *testing.T) {
	con, _ := model.NewSession()
	_, err := model.ReadCustomer(con)
	if err != nil {
		t.Error(err)
	}
	defer con.Close()
}
