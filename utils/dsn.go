package utils

import (
	"fmt"

	"github.com/rajpatelbot/url-shortner/config"
)

// This function is responsible to generate a DSN (Data Source Name) string for Postgres connection.
func GenerateDSN() string {
	// Fetch environment configuration
	env := config.EnvConfig

	// Construct DSN string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		env.HOST,
		env.USER,
		env.PASSWORD,
		env.DB_NAME,
		env.DB_PORT,
		env.SSL_MODE,
		env.TIMEZONE,
	)

	return dsn
}
