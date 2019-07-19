package conekta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type ConektaError struct {
	Object  string   `json:"object,omitempty"`
	Type    string   `json:"type,omitempty"`
	LogId   string   `json:"log_id,omitempty"`
	Details []Detail `json:"details,omitempty"`
}

type Detail struct {
	Debug_message string `json:"debug_message,omitempty"`
	Message       string `json:"message,omitempty"`
	Code          string `json:"code,omitempty"`
}

type body map[string]interface{}

type CustomerInfo struct {
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Name       string `json:"name,omitempty"`
	Corporate  *bool  `json:"corporate,omitempty"`
	CustomerID string `json:"customer_id,omitempty"`
	Object     string `json:"object,omitempty"`
}

type AntifraudInfo map[string]string

var (
	ApiKey, ApiVersion = "", "2.0.0"
)

const (
	conektaURL = "https://api.conekta.io"
)

func request(method, path string, v interface{}) (statusCode int, response []byte) {

	var payload bytes.Reader
	if v != nil {
		jsonPayload, err := json.Marshal(v)
		if err != nil {
			return
		}
		payload = *bytes.NewReader(jsonPayload)
	}

	req, err := http.NewRequest(method, conektaURL+path, &payload)
	if err != nil {
		return
	}
	req.Header.Add("accept", "application/vnd.conekta-v"+ApiVersion+"+json")
	req.SetBasicAuth(ApiKey, "")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return res.StatusCode, body
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("There's an error in Conekta Wrapper: %v\n", err)
	}
}

func ConektaFormatAmount(value float64) (formatted int64, err error) {
	strnum := fmt.Sprintf("%.2f", value)
	strnum = strings.Replace(strnum, ".", "", -1)
	formatted, err = strconv.ParseInt(strnum, 10, 64)
	return
}

func ConektaFormatToFloat64(conektaFormatted int64) float64 {
	return float64(conektaFormatted) / 100
}
