package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Issues struct with array of issues
// an array of users
type Issues struct {
	Issues []Issue `json:"Issues"`
}

// Issue struct
type Issue struct {
	Severity   string `json:"severity"`
	Confidence string `json:"confidence"`
	RuleID     string `json:"rule_id"`
	File       string `json:"file"`
	Line       string `json:"line"`
}

//https://www.dotnetperls.com/between-before-after-go
func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

func main() {
	proxyStr := "http://localhost:8080"
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Println(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	urlStr := "https://sigsci-dev-test.herokuapp.com/gosec"
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
	}
	client := &http.Client{
		Transport: transport,
	}

	jsonFile, err := os.Open("results.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened results.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var issues Issues
	json.Unmarshal(byteValue, &issues)
	for i := 6; i < len(issues.Issues); i++ {
		filePath := after(issues.Issues[i].File, "signalsciences/sigsci")
		postBody, err := json.Marshal(map[string]string{
			"rule_id":    issues.Issues[i].RuleID,
			"severity":   issues.Issues[i].Severity,
			"confidence": issues.Issues[i].Confidence,
			"file_path":  "https://github.com/signalsciences/sigsci/blob/master" + filePath + "#L" + issues.Issues[i].Line,
		})
		request, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(postBody))
		request.Header.Add("Content-Type", "application/json")
		if err != nil {
			log.Println(err)
		}
		response, err := client.Do(request)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Response code: %v", response.StatusCode)
	}

}
