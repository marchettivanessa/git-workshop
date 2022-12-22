package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	var path = "/"
	// Setup Main Server
	echoMainServer := echo.New()
	echoMainServer.GET(path, hello)
	//New random comment
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
