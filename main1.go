package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type quote struct {
	Title  string
	Author string
}

//var quotes []quote = make([]quote, 0)

var quotes []quote = []quote{
	{"Blah", "blah2"},
	{"lead to learn", "svit"},
}

func main() {
	rand.Seed(time.Now().Unix())

	api := echo.New()

	api.GET("/", getQuotes)
	api.GET("/quotes", getRandomQuote)

	port := os.Getenv("PORT")
	if port == "" {
		port = "32445"
	}

	api.Start(":" + port)
}

func getRandomQuote(c echo.Context) error {
	index := rand.Intn(len(quotes))
	return c.JSON(200, quotes[index])
}

func getQuotes(c echo.Context) error {
	return c.JSON(200, quotes)
}
