package password

import(
	"crypto/rand"
	"encoding/base64"
)

func GenerateSalt(length int) (string, error) {
    salt := make([]byte, length)
    _, err := rand.Read(salt)
    if err != nil {
        return "", err
    }
    return base64.StdEncoding.EncodeToString(salt), nil
}