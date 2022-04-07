package config

import (
	"fmt"
	"os"
	"time"

	"github.com/client9/reopen"
	"github.com/kardianos/osext"
	log "github.com/sirupsen/logrus"
)

var (
	AppName       = "go-api-base"
	AppVersion    = getEnvOrDefault("GO_API_BASE_APP_VERSION", "1.0.0")
	RestPort      = getEnvOrDefault("GO_API_BASE_REST_PORT", "12580")
	GrpcPort      = getEnvOrDefault("GO_API_BASE_GRPC_PORT", "12581")
	PidPath       = getEnvOrDefault("GO_API_BASE_PID_PATH", fmt.Sprintf("/tmp/%s.pid", AppName))
	LogPath       = getEnvOrDefault("GO_API_BASE_LOG_PATH", fmt.Sprintf("/tmp/%s.out", AppName))
	UseMigration  = true
	MigrationPath = getEnvOrDefault("GO_API_BASE_MIGRATION_PATH", "")
	LogOutput     *reopen.FileWriter
	ExecDir       = ""
	Env           = getEnvOrDefault("ENV", "dev")
)

func init() {
	ExecDir, _ = osext.ExecutableFolder()

	InitializeLogging()
}

func getEnvOrDefault(envVar string, defEnvVar string) (newEnvVar string) {
	if newEnvVar = os.Getenv(envVar); len(newEnvVar) == 0 {
		return defEnvVar
	} else {
		return newEnvVar
	}
}

func InitializeLogging() {
	var err error
	if LogOutput == nil {
		LogOutput, err = reopen.NewFileWriter(LogPath)
		if err != nil {
			log.Fatalf("Log output file was not set: %s", err)
		}

		// set up log format
		logFormat := &log.JSONFormatter{}
		logFormat.TimestampFormat = time.RFC3339Nano

		log.SetOutput(LogOutput)
		log.SetFormatter(logFormat)
	}
}
