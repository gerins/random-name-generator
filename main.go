package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/labstack/echo/v4"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	City      string `json:"city"`
}

func main() {
	port := flag.Int("port", 8080, "the port to listen on")
	flag.Parse()

	e := echo.New()

	e.GET("/api/person", func(c echo.Context) error {
		person := Person{
			FirstName: randomdata.FirstName(randomdata.RandomGender),
			LastName:  randomdata.LastName(),
			Email:     randomdata.Email(),
			Country:   randomdata.Country(randomdata.FullCountry),
			City:      randomdata.City(),
		}
		return c.JSON(http.StatusOK, person)
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	address := fmt.Sprintf(":%d", *port)
	e.Logger.Fatal(e.Start(address))
}
