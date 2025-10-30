package password

import (
	"golang.org/x/crypto/argon2"
	"encoding/base64"
	"synapse/auth/config"
)

func HashPassword(password, salt string) string {
	pepper := config.GetEnv("PEPPER")
    hash := argon2.IDKey([]byte(pepper + password), []byte(salt), 2, 64*1024, 4, 32)

    return base64.RawStdEncoding.EncodeToString(hash)
}

func VerifyPassword(password, salt, savedHash string) bool {
	pepper := config.GetEnv("PEPPER")
	hash := argon2.IDKey([]byte(pepper + password), []byte(salt), 2, 64*1024, 4, 32)
	
	return base64.RawStdEncoding.EncodeToString(hash) == savedHash
}