package conf

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type configuration struct {
	LogLevel     logrus.Level
	LogFormatter logrus.Formatter

	APIPort int

	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret          string
	JWTExpiryInSeconds int

	SwaggerEnabled bool
}

var Configuration = configuration{
	LogLevel:     getOrDefaultLogLevel("LOG_LEVEL", logrus.DebugLevel),
	LogFormatter: getOrDefaultLogFormatter("LOG_TYPE", &logrus.TextFormatter{PadLevelText: true, DisableQuote: true}),

	APIPort: getOrDefaultInt("API_PORT", 8080),

	DBHost:     getOrDefaultString("DB_HOST", "localhost"),
	DBPort:     getOrDefaultInt("DB_PORT", 5432),
	DBUser:     getOrDefaultString("DB_USER", "user"),
	DBPassword: getOrDefaultString("DB_PASSWORD", "password"),
	DBName:     getOrDefaultString("DB_NAME", "online_book_store"),

	JWTSecret:          getOrDefaultString("JWT_SECRET", "yxolS4eb7YNyxBdLOFDeNVxQEFNiwag4lkvaKnId12uN8HAum8MGRL0WtUtzXZveb4l/T3z+EL3M"),
	JWTExpiryInSeconds: getOrDefaultInt("JWT_EXPIRY_IN_SECONDS", 3600),

	SwaggerEnabled: getOrDefaultBool("SWAGGER_ENABLED", true),
}
