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

const PoligonPath = "https://api.polygon.io/"
const ApiKey = "vLcR3SaJGTOsLucDg_s2E2BDfjnjaVyO"

const TickerPath = PoligonPath + "v2/aggs/ticker/"
const DailyValuesPath = PoligonPath + "v1/open-close/"

type Stock struct {
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
}

type SearchResult struct {
	Results []Stock `json:"results"`
}

type Values struct {
	Symbol     string  `json:"symbol"`
	Open       float64 `json:"open"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Close      float64 `json:"close"`
	AfterHours float64 `json:"afterHours"`
}

func getCurrentDate() string {
	currentTime := time.Now()
	dayBefore := currentTime.AddDate(0, 0, -1)
	return dayBefore.Format("2006-01-02")
}

func Fetch(path string) string {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func SearchTicker(ticker string) []Stock {
	// body := Fetch(TickerPath + "?" + ApiKey + "&ticker=" + strings.ToUpper(ticker))
	// fmt.Println(TickerPath + strings.ToUpper(ticker) + "/range/1/day/" + getCurrentDate() + "/" + getCurrentDate() + "/?apiKey=" + ApiKey)
	body := Fetch(TickerPath + strings.ToUpper(ticker) + "/range/1/day/" + getCurrentDate() + "/" + getCurrentDate() + "/?apiKey=" + ApiKey)
	fmt.Println("Line 62: " + body)
	data := SearchResult{}

	// json.Unmarshal([]byte(string(body)), &data)
	json.Unmarshal([]byte(string(body)), &data)

	fmt.Println("Line 68: ", data)
	return data.Results
}

func GetDailyValues(ticker string) Values {
	// body := Fetch(DailyValuesPath + "/" + strings.ToUpper(ticker) + "/2023-09-15/?" + ApiKey)
	body := Fetch(DailyValuesPath + strings.ToUpper(ticker) + "/" + getCurrentDate() + "/?adjusted=true&apiKey=" + ApiKey)
	fmt.Println(DailyValuesPath + strings.ToUpper(ticker) + "/" + getCurrentDate() + "/?adjusted=true&apiKey=" + ApiKey)
	fmt.Println("Line 76: " + body)
	data := Values{}
	json.Unmarshal([]byte(string(body)), &data)
	// fmt.Println(data)
	fmt.Println("Line 80: " + fmt.Sprintf("%v", data))
	return data
}
