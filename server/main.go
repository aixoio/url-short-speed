package main

import (
	"github.com/aixoio/url-short-speed/server/db"
	"github.com/aixoio/url-short-speed/server/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.Connect()
	router.StartRouter()
}
