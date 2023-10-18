package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
)

//
// ParseRsaPublicKeyFromPemStr parses a PEM encoded RSA public key string.
// It returns the parsed key or an error if the key could not be parsed.
//
func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	//
	// Decode the PEM block
	//
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	//
	// Parse the public key
	//
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	//
	// Check if the key is of type RSA
	//
	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}

//
// RSA_PKCS_Encrypt encrypts a secret message using RSA PKCS1v15.
// It returns an error and the encrypted message as a base64 encoded string.
//
func RSA_PKCS_Encrypt(secretMessage string, key rsa.PublicKey) (err error, encrypted string) {
	rng := rand.Reader
	ciphertext, err := rsa.EncryptPKCS1v15(rng, &key, []byte(secretMessage))
	if err == nil {
		encrypted = base64.StdEncoding.EncodeToString(ciphertext)
	}

	return err, encrypted
}

//
// GenerateEncryptedRequest generates an encrypted request using a public key.
// It returns an error and the encrypted data as a base64 encoded string.
//
func GenerateEncryptedRequest(response map[string]interface{}, publickeypem string) (err error, encdata string) {
	err = nil
	encdata = ""
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publickeypem)

	if err != nil {
		fmt.Println("Unable to decode public key: " + err.Error())
		return err, ""
	}

	//
	// Convert the response map to JSON
	//
	jsonStr, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Unable to generate json from realtime valudation map: " + err.Error())
		return err, ""
	}

	//
	// Parse the public key
	//
	key, err := ParseRsaPublicKeyFromPemStr(string(decodedPublicKey))
	if err != nil {
		fmt.Println("Unable to load public key: " + err.Error())
		return err, ""
	}

	//
	// Encrypt the JSON string and store it in the response map
	//
	err, encdata = RSA_PKCS_Encrypt(string(jsonStr), *key)
	if err != nil {
		fmt.Println("Unable to encrypt the payload: " + err.Error())
		return err, ""
	}
	return err, encdata
}
