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

// DeleteCustomer Removes the Customer by ID
func DeleteCustomer(s *mgo.Session, id string) error {
	session := s.Copy()
	col := session.DB("cadastro").C("customers")
	err := col.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	session.Close()
	return err
}

// FindCustomer search the customer by ID
func FindCustomer(s *mgo.Session, id string) (*Customer, error) {
	session := s.Copy()
	col := session.DB("cadastro").C("customers")
	var customer Customer
	err := col.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&customer)

	if err != nil {
		return nil, err
	}

	session.Close()
	return &customer, nil
}

// UpdateCustomer update the customer by ID
func UpdateCustomer(s *mgo.Session, id string, customer Customer) error {
	session := s.Copy()
	col := session.DB("cadastro").C("customers")
	err := col.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &customer)
	session.Close()
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
