package model

import (
	mgo "gopkg.in/mgo.v2"
)

// NewSession instantiate the Session
func NewSession() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)

	if err != nil {
		return nil, err
	}
	return session, nil
}
