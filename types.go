package main

type RegisterRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token string
}

type User struct {
	Name      string   `json:"name,omitempty" bson:"name,omitempty"`
	PassHash  []byte   `json:"passhash,omitempty" bson:"passhash,omitempty"`
	Salt      []byte   `json:"salt,omitempty" bson:"salt,omitempty"`
	Password  string   `json:"password,omitempty" bson:"password,omitempty"`
	Challenge []byte   `json:"challenge,omitempty" bson:"challenge,omitempty"`
}
