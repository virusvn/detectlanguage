package detectlanguage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Detections language
type Detections struct {
	Confidence float64 `json:"confidence"`
	IsReliable bool    `json:"isReliable"`
	Language   string  `json:"language"`
}

type DataDetections struct {
	Data [][]Detections
}
type DetectionLanguage struct {
	Data struct {
		Detections []Detections `json:"detections"`
	} `json:"data"`
}
type DetectionLanguages struct {
	Data struct {
		Detections [][]Detections `json:"detections"`
	} `json:"data"`
}

// Status of account
type AccountStatus struct {
	Bytes              int         `json:"bytes"`
	DailyBytesLimit    int         `json:"daily_bytes_limit"`
	DailyRequestsLimit int         `json:"daily_requests_limit"`
	Date               string      `json:"date"`
	Plan               string      `json:"plan"`
	PlanExpires        interface{} `json:"plan_expires"`
	Requests           int         `json:"requests"`
	Status             string      `json:"status"`
}

// Language supported
type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

const detectLanguageEndpoint = "ws.detectlanguage.com/0.2/"

func getServiceURL(useHTTPS bool) string {
	if useHTTPS {
		return "https://" + detectLanguageEndpoint
	}
	return "http://" + detectLanguageEndpoint
}

// Detect language by text
func Detect(key string, text string, useHTTPS bool) DetectionLanguage {
	serviceURL := getServiceURL(useHTTPS)
	escapedText := url.QueryEscape(text)
	fullURL := serviceURL + fmt.Sprintf("detect?q=%s&key=%s", escapedText, key)
	resp := fetch(fullURL)
	var myJSON DetectionLanguage
	json.Unmarshal(resp, &myJSON)
	fmt.Println(myJSON)
	return myJSON
}

// Detect multi sentences
func DetectMulti(key string, text []string, useHTTPS bool) DetectionLanguages {
	serviceURL := getServiceURL(useHTTPS)
	fullURL := serviceURL + fmt.Sprintf("detect?key=%s", key)
	for _, t := range text {
		escapedText := url.QueryEscape(t)
		fullURL += "&q[]=" + escapedText
	}
	resp := fetch(fullURL)
	var myJSON DetectionLanguages
	json.Unmarshal(resp, &myJSON)
	fmt.Println(myJSON)
	return myJSON
}

// Get Account Status
func GetAccountStatus(key string, useHTTPS bool) AccountStatus {
	serviceURL := getServiceURL(useHTTPS)
	fullURL := serviceURL + fmt.Sprintf("user/status?key=%s", key)
	resp := fetch(fullURL)
	var myJSON AccountStatus
	json.Unmarshal(resp, &myJSON)
	return myJSON
}

// Get languages supported
func GetLanguagesSupported(useHTTPS bool) []Language {
	serviceURL := getServiceURL(useHTTPS)
	fullURL := serviceURL + fmt.Sprintf("languages")
	var myJSON []Language
	resp := fetch(fullURL)
	json.Unmarshal(resp, &myJSON)
	return myJSON
}

// Fetch url
func fetch(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bytes))
		return bytes
	}
	return []byte("")
}

// Is language
func Is(languageShortCode string, key string, text string, useHTTPS bool) bool {
	languages := Detect(key, text, useHTTPS)
	for _, language := range languages.Data.Detections {
		if language.Language == languageShortCode && language.IsReliable {
			return true
		}
	}
	return false
}
