package main

type User struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	PassHash []byte `json:"passhash,omitempty" bson:"passhash,omitempty"`
	Salt     []byte `json:"salt,omitempty" bson:"salt,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}
