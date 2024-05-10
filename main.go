package main

import (
	"log"
	_ "github.com/lib/pq"
	"github.com/xnpltn/azshop/app"
)

const PORT = 6969


func main(){
	app := app.NewApp()
	log.Fatal(app.Start(PORT))
}