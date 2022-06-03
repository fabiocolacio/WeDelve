package main

import (
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
)

const (
	SALT_SIZE int = 16
	SCRYPT_N int = 32768
	SCRYPT_r int = 8
	SCRYPT_P int = 1
	PHASH_SIZE = 32
)

func SaltAndHash(password string) (salt []byte, hash []byte, err error) {
	salt = make([]byte, 16)
	if _, err = rand.Read(salt); err == nil {
		hash, err = scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	}
	return
}
