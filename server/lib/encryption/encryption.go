package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"os"

	"github.com/tfkhdyt/fiber-toolbox/exception"
)

var encryptionKey = os.Getenv("ENCRYPTION_KEY")

func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, exception.NewInternalServerError("failed to generate keys", err)
	}
	return privateKey, &privateKey.PublicKey, nil
}

func EncryptPrivateKey(privateKey *rsa.PrivateKey) (string, error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	encryptedKey := make([]byte, len(privateKeyBytes))
	for i := range privateKeyBytes {
		encryptedKey[i] = privateKeyBytes[i] ^ encryptionKey[i%len(encryptionKey)]
	}
	return base64.StdEncoding.EncodeToString(encryptedKey), nil
}

func DecryptPrivateKey(encryptedKey string) (*rsa.PrivateKey, error) {
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedKey)
	if err != nil {
		return nil, exception.NewInternalServerError("failed to decrypt private key", err)
	}
	decryptedKey := make([]byte, len(encryptedBytes))
	for i := range encryptedBytes {
		decryptedKey[i] = encryptedBytes[i] ^ encryptionKey[i%len(encryptionKey)]
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(decryptedKey)
	if err != nil {
		return nil, exception.NewInternalServerError("failed to parse private key", err)
	}
	return privateKey, nil
}

func MarshalPublicKey(publicKey *rsa.PublicKey) (string, error) {
	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyStr := base64.StdEncoding.EncodeToString(publicKeyBytes)
	return publicKeyStr, nil
}

func MarshalPrivateKey(privateKey *rsa.PrivateKey) (string, error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyStr := base64.StdEncoding.EncodeToString(privateKeyBytes)
	return privateKeyStr, nil
}
