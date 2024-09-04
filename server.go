package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
)

// type Stock struct {
// 	Ticker string
// 	Price  float64
// }

// func Fetch(ticker string) Stock {
// 	// Replace with actual implementation
// 	price := 0.0 // Replace 0.0 with the actual price
// 	return Stock{Ticker: ticker, Price: price}
// }

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/search", func(c fiber.Ctx) error {
		ticker := c.Query("ticker")
		stockData, err := fetchStockData(ticker)
		if err != nil {
			// Handle the error
			return err
		}

		return c.Render("results", fiber.Map{"Results": stockData})
	})

	app.Get("/values/:ticker", func(c fiber.Ctx) error {
		ticker := c.Params("ticker")
		Values := GetDailyValues(ticker)
		return c.Render("values", fiber.Map{
			"Ticker": ticker,
			"Values": Values,
		})
	})

	app.Listen(":3000")
}
