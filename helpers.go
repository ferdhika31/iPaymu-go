package ipaymu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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

// from https://gist.github.com/bxcodec/c2a25cfc75f6b21a0492951706bc80b8#gistcomment-3381061
func structToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}
