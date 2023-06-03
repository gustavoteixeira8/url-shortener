package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	_, b, _, _ = runtime.Caller(0)
	currentDir = filepath.Dir(b)
)

func LoadDotenv() {
	if strings.ToLower(GetEnv("DOCKER_RUNNING")) == "true" {
		return
	}

	err := godotenv.Load(fmt.Sprintf("%s/../../.env", currentDir))

	if err != nil {
		logrus.Fatalf(err.Error())
	}

	logrus.Info("Dotenv loaded")
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
