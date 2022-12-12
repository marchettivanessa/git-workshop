package main

import "net/http"

func main() {
	var path = "/"
	// Setup Main Server
	echoMainServer := echo.New()
	echoMainServer.GET(path, hello)
}

func hello(c echo.Context) error {
	
	return c.String(http.StatusOK, "Hello, World!")
}

