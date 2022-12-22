package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var path = "/"
	// Setup Main Server
	echoMainServer := echo.New()
	//This hides echo banner when commented
	//echoMainServer.HideBanner = true
	echoMainServer.Use(middleware.Logger())
	echoMainServer.GET(path, hello)

	// Create Prometheus server and a Middleware
	echoPrometheus := echo.New()
	echoPrometheus.HideBanner = true
	prom := prometheus.NewPrometheus("echo", nil)

	// Scrape metrics from Main Server
	echoMainServer.Use(prom.HandlerFunc)
	// Setup metrics endpoint at another server
	prom.SetMetricsPath(echoPrometheus)

	go func() { echoPrometheus.Logger.Fatal(echoPrometheus.Start(":9360")) }()

	echoMainServer.Logger.Fatal(echoMainServer.Start(":7090"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}


