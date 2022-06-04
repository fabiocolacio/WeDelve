package main

import (
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
)

const (
	SALT_SIZE      int = 16
	SCRYPT_N       int = 32768
	SCRYPT_r       int = 8
	SCRYPT_p       int = 1
	PHASH_SIZE     int = 32
	CHALLENGE_SIZE int = 32
)

func NewChallenge() (challenge []byte, err error) {
	challenge = make([]byte, CHALLENGE_SIZE)
	_, err = rand.Read(challenge)
	return
}

func SaltAndHash(password string) (salt []byte, hash []byte, err error) {
	salt = make([]byte, SALT_SIZE)
	if _, err = rand.Read(salt); err == nil {
		hash, err = scrypt.Key([]byte(password), salt, SCRYPT_N, SCRYPT_r, SCRYPT_p, PHASH_SIZE)
	}
	return
}
