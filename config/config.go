package config

import (
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/client9/reopen"
	"github.com/kardianos/osext"
)

var (
	AppName         = getEnvOrDefault("APP_NAME", "your_app")
	AppVersion      = getEnvOrDefault("APP_VERSION", "1.0.0")
	RestPort        = getEnvOrDefault("APP_PORT", "12580")
	TCPPort         = getEnvOrDefault("TCP_PORT", "12573")
	RunTCPPort      = getEnvOrDefault("RUN_TCP_PORT", "false")
	PidPath         = getEnvOrDefault("PID_PATH", fmt.Sprintf("/tmp/%s.pid", AppName))
	DBUser          = getEnvOrDefault("DB_USER", "root")
	DBPass          = getEnvOrDefault("DB_PASS", "")
	DBDB            = getEnvOrDefault("DB_DB", "scaffold_test")
	DBHost          = getEnvOrDefault("DB_HOST", "")
	LogPath         = getEnvOrDefault("LOG_PATH", fmt.Sprintf("/tmp/%s.out", AppName))
	LogOutput       *reopen.FileWriter
	ExecDir         = ""
	Env             = getEnvOrDefault("ENV", "dev")
	StorageSQL      = true
	StoragePsql     = true
	StorageSqlite   = false
	SqlitePath      = "/tmp/scaffold_test.db"
	StorageMysql    = false
	StorageGorm     = true
	StorageFile     = false
	StorageFilePath = "/tmp/"
	StorageMongo    = true
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
