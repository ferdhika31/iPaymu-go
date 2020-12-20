package iPaymu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"strings"
)

func StringSha256(s string) []byte {
	input := strings.NewReader(s)
	h := sha256.New()

	if _, err := io.Copy(h, input); err != nil {
		log.Fatal(err)
	}
	sum := h.Sum(nil)

	return sum
}

func StringHMACSha256(data string, secret string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
