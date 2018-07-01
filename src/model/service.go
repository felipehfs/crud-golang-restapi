package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CreateCustomer create the Entity
func CreateCustomer(s *mgo.Session, c Customer) error {
	session := s.Copy()
	collection := session.DB("cadastro").C("customers")
	defer session.Close()
	err := collection.Insert(c)
	return err
}

// ReadCustomer retrieves all data
func ReadCustomer(s *mgo.Session) ([]Customer, error) {
	session := s.Copy()
	col := session.DB("cadastro").C("customers")
	var customers []Customer
	err := col.Find(bson.M{}).All(&customers)
	session.Close()

	if err != nil {
		return customers, err
	}
	return customers, nil
}
