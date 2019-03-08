package hmacutil

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

const (
	// MD5 algorithm
	MD5 string = "MD5"
	// SHA1 algorithm
	SHA1 string = "SHA1"
	// SHA256 algorithm
	SHA256 string = "SHA256"
	// SHA512 algorithm
	SHA512 string = "SHA512"
)

var algos = make(map[string]func() hash.Hash)

func init() {
	algos[MD5] = md5.New
	algos[SHA1] = sha1.New
	algos[SHA256] = sha256.New
	algos[SHA512] = sha512.New
}

// Encode computes a HMAC using secret key
func Encode(algorithm string, key string, data string) []byte {
	mac := hmac.New(algos[algorithm], []byte(key))
	mac.Write([]byte(data))
	return mac.Sum(nil)
}

// Base64Encode return HMAC in base64 encoded
func Base64Encode(algorithm string, key string, data string) string {
	hmac := Encode(algorithm, key, data)
	return base64.StdEncoding.EncodeToString(hmac)
}

// HexStringEncode return HMAC in hex string encoded
func HexStringEncode(algorithm string, key string, data string) string {
	hmac := Encode(algorithm, key, data)
	return hex.EncodeToString(hmac)
}
