package conf

import (
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func getOrDefaultLogLevel(key string, defaultValue logrus.Level) logrus.Level {
	value, err := logrus.ParseLevel(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return value
}

func getOrDefaultLogType(key string, defaultValue LogType) LogType {
	value := os.Getenv(key)
	switch strings.ToLower(value) {
	case strings.ToLower(string(LogTypeJSON)):
		return LogTypeJSON
	case strings.ToLower(string(LogTypeText)):
		return LogTypeText
	default:
		return defaultValue
	}
}

func getOrDefaultLogFormatter(key string, defaultValue logrus.Formatter) logrus.Formatter {
	switch getOrDefaultLogType(key, "") {
	case LogTypeJSON:
		return &logrus.JSONFormatter{}
	case LogTypeText:
		return &logrus.TextFormatter{PadLevelText: true, DisableQuote: true}
	default:
		return defaultValue
	}
}

func getOrDefaultString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getOrDefaultInt(key string, defaultValue int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return value
}

func getOrDefaultBool(key string, defaultValue bool) bool {
	value, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return value
}
