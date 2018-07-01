package model

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// Customer represents the personal info of clients
type Customer struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name,omitempty" bson:"name,omitempty"`
	Email    string        `json:"email" bson:"email,omitempty"`
	Phone    string        `json:"phone,omitempty" bson:"phone,omitempty"`
	ZipCode  string        `json:"zipcode,omitempty" bson:"zipcode,omitempty"`
	Street   string        `json:"street,omitempty" bson:"street,omitempty"`
	Neighbor string        `json:"neighbor,omitempty" bson:"neighbor,omitempty"`
}

func (c Customer) String() string {
	return fmt.Sprintf("%v %v %v %v %v %v", c.Name, c.Email, c.Phone,
		c.ZipCode, c.Street, c.Neighbor)
}
