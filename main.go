package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var firstNames = []string{
	"John",
	"Jane",
	"Alice",
	"Bob",
	"Emma",
	"David",
	"Grace",
	"Harry",
	"Olivia",
	"Peter",
}

var lastNames = []string{
	"Smith",
	"Johnson",
	"Brown",
	"Taylor",
	"Davis",
	"Wilson",
	"Lee",
	"Jones",
	"Garcia",
	"Miller",
}

func main() {
	port := flag.Int("port", 8080, "the port to listen on")
	flag.Parse()

	e := echo.New()

	e.GET("/api/person", func(c echo.Context) error {
		person := Person{
			FirstName: firstNames[rand.Intn(len(firstNames))],
			LastName:  lastNames[rand.Intn(len(lastNames))],
		}
		return c.JSON(http.StatusOK, person)
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	address := fmt.Sprintf(":%d", *port)
	e.Logger.Fatal(e.Start(address))
}
