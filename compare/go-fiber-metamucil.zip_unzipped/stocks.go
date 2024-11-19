package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

const (
	PoligonPath = "https://api.polygon.io/"
	ApiKey      = "vLcR3SaJGTOsLucDg_s2E2BDfjnjaVyO"
	// TickerPath      = PoligonPath + "v2/aggs/?ticker="
	TickerPath      = PoligonPath + "v3/reference/tickers/"
	DailyValuesPath = PoligonPath + "v2/aggs/ticker/"
)

type Stock struct {
	Result struct {
		Ticker string `json:"ticker"`
		Name   string `json:"name"`
		// Price  float64 `json:"price"`
	} `json:"results"`
}

type SearchResult struct {
	Stocks []Stock `json:"stocks"`
}

type Values struct {
	Ticker string    `json:"ticker"`
	Date   time.Time `json:"results[0].t"`
	Open   float64   `json:"results[0].o"`
	Close  float64   `json:"results[0].c"`
	High   float64   `json:"results[0].h"`
	Low    float64   `json:"results[0].l"`
	Volume int       `json:"results[0].v"`
}

func getCurrentDate() string {
	currentTime := time.Now()
	dayBefore := currentTime.AddDate(0, 0, -3)
	return dayBefore.Format("2006-01-02")
}

func Fetch(path string) string {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal("Got Error doing a REST GET with that URL path with this: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Got Error on returned Response Body with: %v", err)
	}

	return string(body)
}

func SearchTicker(ticker string) Stock {
	fmt.Println(TickerPath + strings.ToUpper(ticker) + "?apiKey=" + ApiKey + "\n")
	fmt.Println(DailyValuesPath + strings.ToUpper(ticker) + "/range/1/day/" + getCurrentDate() + "/" + getCurrentDate() + "?apiKey=" + ApiKey + "\n")
	body := Fetch(TickerPath + strings.ToUpper(ticker) + "?apiKey=" + ApiKey)

	fmt.Printf("Line 64: %v\n", body)

	var data Stock

	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatal("Error while trying to unmarshal the JSON: %v", err)
	}

	fmt.Printf("Unmarshalled Data Line 68: %#v", data)
	return data
}

func GetDailyValues(ticker string) []Values {
	// https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/2024-08-30/2024-08-30?apiKey=vLcR3SaJGTOsLucDg_s2E2BDfjnjaVyO
	fmt.Println(DailyValuesPath + strings.ToUpper(ticker) + "/range/1/day/" + getCurrentDate() + "/" + getCurrentDate() + "?apiKey=" + ApiKey)
	body := Fetch(DailyValuesPath + strings.ToUpper(ticker) + "/range/1/day/" + getCurrentDate() + "/" + getCurrentDate() + "?apiKey=" + ApiKey)
	fmt.Printf("Line 78: %v\n", body)
	var data []Values
	json.Unmarshal([]byte(body), &data)
	fmt.Printf("Line 81: %v\n", data)
	return data
}
