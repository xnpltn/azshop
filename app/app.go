package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/xnpltn/azshop/database"
)

type Application struct {
	db     *database.Queries
	Router *echo.Echo
}

func NewApp() *Application {
	r := echo.New()

	return &Application{
		Router: r,
	}
}

func (a *Application) Start(port uint32) error {

	conn, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Can't connect to databse: ", err)
	}
	a.db = database.New(conn)
	a.Load()
	err = a.Router.Start(fmt.Sprintf(":%d", port))
	return err
}
func (a *Application) Stop() error {
	return nil
}

func (a *Application) Load() {

	// views
	a.Router.Static("/static", "assets")
	a.Router.GET("/", a.HomeGetHandler())
	a.Router.GET("/about", a.AboutGetHandler())
	a.Router.GET("/login", a.LoginGetHandler())
	a.Router.GET("/admin", a.Admin())
	a.Router.GET("/cart", a.Cart())
	a.Router.GET("/profile", a.Profile())

	// auth
	auth := a.Router.Group("/auth")
	auth.POST("/login", a.Login())
	auth.POST("/signup", a.Signup())
	auth.POST("/logout", a.Logout())

	// api
	admin := a.Router.Group("/admin")
	admin.POST("/product", a.CreateProduct)
	admin.POST("/product/delete", a.DeleteProduct)


	// product
	product := a.Router.Group("/product")
	product.GET("/:id", a.Product())
	product.POST("/cart", a.AddToCarT())
}
