package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/xnpltn/azshop/database"
)


type Application struct{
	db *database.Queries
	Router *echo.Echo
}

func NewApp() *Application{
	r := echo.New()

	return &Application{
		Router: r,
	}
}


func(a *Application) Start(port uint32) error{
	conn, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5433/azshop?sslmode=disable")
	if err != nil{
		log.Fatal("Can't connect to databse: ", err)
	}
	a.db  = database.New(conn)
	a.Load()
	err = a.Router.Start(fmt.Sprintf(":%d", port))
	return err
}
func(a *Application)Stop()error{
	return nil
}


func(a *Application) Load(){
	
	
	// views
	a.Router.Static("/static", "assets")
	a.Router.GET("/", a.HomeGetHandler())
	a.Router.GET("/about", a.AboutGetHandler())
	a.Router.GET("/admin", a.Admin())

	

	// api
	admin := a.Router.Group("/admin")
	admin.POST("/product", a.CreateProduct)
}


