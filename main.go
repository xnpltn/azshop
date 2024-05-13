package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/xnpltn/azshop/app"
)

const PORT = 6969

func init() {
	godotenv.Load()
}

func main() {
	app := app.NewApp()
	log.Fatal(app.Start(PORT))
}
