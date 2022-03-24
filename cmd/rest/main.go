package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/keenfury/go-api-base/config"
	ae "github.com/keenfury/go-api-base/internal/api_error"
	m "github.com/keenfury/go-api-base/internal/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	// --- replace server header text ---
)

func main() {
	setPidFile()

	// argument flag
	var restPort string
	flag.StringVar(&restPort, "restPort", "", "the port number used for the REST listener")

	flag.Parse()

	if restPort == "" {
		restPort = config.RestPort
	}

	e := echo.New()
	e.HTTPErrorHandler = ae.ErrorHandler // set echo's error handler
	if !strings.Contains(config.Env, "prod") {
		log.Infoln("Logging set to debug...")
		e.Debug = true
		e.Use(m.DebugHandler)
	}
	e.Use(
		middleware.Recover(),
		m.Handler,
	)

	// set output based on config value
	m.SetLogOutput(config.LogOutput)

	// set all non-endpoints here
	e.GET("/", Index)
	e.HEAD("/server_status", ServerStatus)
	e.HEAD("/status", ServerStatus)

	InitializeRoutes(e)

	e.Start(fmt.Sprintf(":%s", restPort))
}

func setPidFile() {
	// purpose: to set the starting applications pid number to file
	if pidFile, err := os.Create(config.PidPath); err != nil {
		log.Panicln("Unable to create pid file...")
	} else if _, err := pidFile.Write([]byte(fmt.Sprintf("%d", os.Getpid()))); err != nil {
		log.Panicln("Unable to write pid to file...")
	}
}

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the GO_API_BASE API")
}

func ServerStatus(c echo.Context) error {
	c.Response().Header().Add("GO_API_BASE_SERVICE", config.AppVersion)
	c.Response().WriteHeader(http.StatusOK)
	return nil
}

func InitializeRoutes(e *echo.Echo) {
	// initialize all routes here
	// --- replace server once text - do not remove ---
	// --- replace server text - do not remove ---

}
