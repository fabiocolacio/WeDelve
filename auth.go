package main

import (
	"crypto/rand"
	"log"
	"time"
	"golang.org/x/crypto/scrypt"
	"github.com/golang-jwt/jwt/v4"
)

const (
	SaltSize      int = 16
	ScryptN       int = 32768
	Scryptr       int = 8
	Scryptp       int = 1
	PHashSize     int = 32
	ChallengeSize int = 32
	JwtKeySize    int = 32
)

var (
	jwtKey []byte
)

func init() {
	jwtKey = make([]byte, JwtKeySize)
	if _, err := rand.Read(jwtKey); err != nil {
		log.Fatal("Failed to generate signing key for JWT tokens")
	}
}

func NewChallenge() (challenge []byte, err error) {
	challenge = make([]byte, ChallengeSize)
	_, err = rand.Read(challenge)
	return
}

func SaltAndHash(password string) (salt []byte, hash []byte, err error) {
	salt = make([]byte, SaltSize)
	if _, err = rand.Read(salt); err == nil {
		hash, err = scrypt.Key([]byte(password), salt, ScryptN, Scryptr, Scryptp, PHashSize)
	}
	return
}

func HashPassword(password string, salt []byte) (hash []byte, err error) {
	hash, err = scrypt.Key([]byte(password), salt, ScryptN, Scryptr, Scryptp, PHashSize)
	return
}

func NewToken(user User) (string, error) {
	claims := &jwt.StandardClaims{
		Subject: user.Name,
		IssuedAt: time.Now().Unix(),
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
