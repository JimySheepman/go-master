package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func httpParam() {
	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 15 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 15 * time.Second,
	}

	var client = &http.Client{
		Timeout:   time.Second * 30,
		Transport: transport,
	}

	uri := fmt.Sprintf("%s%s/%s/transactions", "https://api.uat.test.com.tr", "/api/test/v2/accounts", "4412341234")
	log.Println("URL: ", uri)
	bearerToken := fmt.Sprintf("Bearer %s", "test")

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Println("Error: ", err)
	}
	req.Header.Set("Authorization", bearerToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Client-Id", "test")
	req.Header.Set("X-Client-Secret", "test")
	req.Header.Set("X-Client-Certificate", "test")

	// GET /api/test/v2/accounts/{account-id}/transactions?min={min-amount}&max={max-amount}&from={from-date}&to={to-date}
	// GET /api/test/v1/accounts/44235532/transactions?from=2023-01-09T00:00:00&to=2023-01-10T00:00:00&min=1&max=10000
	a := req.URL.Query()
	a.Add("min", "test")
	a.Add("max", "test")
	a.Add("from", "test")
	a.Add("to", "test")

	req.URL.RawQuery = a.Encode()
	log.Println(req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error: ", err)
	}
	defer resp.Body.Close()
}

func urlCheck() {
	uri := fmt.Sprintf("%s%s/%s/transactions", "https://api.uat.isbank.com.tr", "/api/isbank/v2/accounts", "4412341234")

	req, _ := http.NewRequest("GET", uri, nil)
	a := req.URL.Query()
	a.Add("min", "0")
	a.Add("max", "1000000000000")
	a.Add("from", "2006-01-02T15:04:05")
	a.Add("to", "2016-01-02T15:04:05")

	req.URL.RawQuery = a.Encode()
	log.Println(req.URL.RawQuery)
}

func urlencodeTest() {
	Scopes := "key1:value1 key2:value2 key3:value3"

	stepOne := strings.Replace(Scopes, ":", "%3A", -1)
	stepTwo := strings.Replace(stepOne, " ", "%20", -1)

	a := fmt.Sprintf("scope=%s", stepTwo)

	params := url.Values{}
	params.Add("scope", Scopes)
	urlEncode := params.Encode()

	fmt.Println("a1= ", a)
	fmt.Println("a2= ", urlEncode)
	fmt.Println("Test: ", a == urlEncode)
}
