package config

import "os"

var (
	DB_URL = os.Getenv("DB_URL") // DATABASE's URL.

	API_PORT = os.Getenv("PORT") // API PORT.

	SECRET_KEY = os.Getenv("SECRET_KEY") // JWT SECRET KEY.

	HTTP_ORIGINS = os.Getenv("HTTP_ORIGINS")

	HASH_COST = os.Getenv("HASH_COST") // BCRYPT SALT OR COST.
)
