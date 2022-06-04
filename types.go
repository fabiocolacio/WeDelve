package main

type User struct {
	Id        RecordId `json:"_id,omitempty" bson"_id,omitempty"`
	Name      string   `json:"name,omitempty" bson:"name,omitempty"`
	PassHash  []byte   `json:"passhash,omitempty" bson:"passhash,omitempty"`
	Salt      []byte   `json:"salt,omitempty" bson:"salt,omitempty"`
	Password  string   `json:"password,omitempty" bson:"password,omitempty"`
	Challenge []byte   `json:"challenge,omitempty" bson:"challenge,omitempty"`
}
