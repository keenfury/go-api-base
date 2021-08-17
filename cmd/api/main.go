package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/keenfury/api/config"
	ae "github.com/keenfury/api/internal/api_error"
	m "github.com/keenfury/api/internal/middleware"
	"github.com/keenfury/api/internal/protobuf"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// --- replace server header text ---
)

func main() {
	setPidFile()

	// argument flag
	var restPort, tcpPort, runTCP string

	flag.StringVar(&restPort, "restport", "", "the port number used for the API restful listener")
	flag.StringVar(&tcpPort, "tcpport", "", "the port number used for the tcp/protobuf listener")
	flag.StringVar(&runTCP, "runtcp", "", "toggle to start up the tcp/protobuf listener (true/false)")

	flag.Parse()

	if restPort == "" {
		restPort = config.RestPort
	}
	if tcpPort == "" {
		tcpPort = config.TCPPort
	}
	if runTCP == "" {
		runTCP = config.RunTCPPort
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

	// start tcp server if applicable
	if runTCP == "true" {
		protobuf.StartTCPPort(tcpPort)
	}

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
	return c.String(http.StatusOK, "Welcome to the API")
}

func ServerStatus(c echo.Context) error {
	c.Response().Header().Add("API_SERVICE", config.AppVersion)
	c.Response().WriteHeader(http.StatusOK)
	return nil
}

func InitializeRoutes(e *echo.Echo) {
	// initialize all routes here
	routeGroup := e.Group("v1") // may want to add "v1" if doing versioning
	// --- replace server text - do not remove ---

}
