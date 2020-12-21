package ipaymu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func GenerateSignature(body string, method string, cfg Config) string {
	va := cfg.VirtualAccount
	secret := cfg.ApiSecret

	requestBody := fmt.Sprintf("%x", StringSha256(body))

	stringToSign := strings.ToUpper(method) + ":" + va + ":" + requestBody + ":" + secret

	return StringHMACSha256(stringToSign, secret)
}

func Request(method string, url string, body map[string]interface{}, cfg Config) map[string]interface{} {

	var API_URL = "https://my.ipaymu.com"
	if strings.ToLower(cfg.Env) == "sandbox" {
		API_URL = "http://sandbox.ipaymu.com"
	}
	var getURL = API_URL + url

	var jsonBody []byte
	jsonBody, err := json.Marshal(body)

	if err != nil {
		log.Println(err)
	}

	req, _ := http.NewRequest(method, getURL,
		bytes.NewBuffer([]byte(string(jsonBody))))

	t := time.Now()
	timestamp := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), method, cfg))

	// fmt.Println("Signature : " + signature)
	// fmt.Println("timestamps : " + string(timestamp))

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("va", cfg.VirtualAccount)
	req.Header.Add("signature", string(signature))
	req.Header.Add("timestamp", string(timestamp))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf(fmt.Sprint(err))
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s\n", err)
	}

	var data map[string]interface{}

	err = json.Unmarshal(b, &data)

	if err != nil {
		log.Printf("Failed unmarshaling: %s\n", err)
	}

	return data
}
