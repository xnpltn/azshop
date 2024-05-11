package app

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/xnpltn/azshop/database"
	"github.com/xnpltn/azshop/templates"
	"github.com/xnpltn/azshop/templates/views"
)

func(a *Application) RenderView(c echo.Context, view templ.Component, layoutPath string) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return view.Render(c.Request().Context(), c.Response().Writer)
	}
	 return templates.Layout(layoutPath, false).Render(c.Request().Context(), c.Response().Writer)
}

func(a *Application) HomeGetHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {
		products, err := a.db.GetAllProducts(c.Request().Context())
		if err!= nil{
			fmt.Println(err)
			return a.RenderView(c, views.HomeView("Home", []database.Product{}), "/")
		}

		return a.RenderView(c, views.HomeView("Hello", products), "/")
	}

}

func(a *Application) AboutGetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return a. RenderView(c, views.AboutView("About"), "/about")
	}

}

func(a *Application) LoginGetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return a. RenderView(c, views.LoginView("Login"), "/login")
	}

}


func(a *Application) Admin()echo.HandlerFunc{
	return func(c echo.Context) error {
		return a.RenderView(c, views.AdminView("Admin"), "/admin")
	}
}

func(a *Application) Product()echo.HandlerFunc{
	return func(c echo.Context) error {
		id:= c.Param("id")
		id_, err := uuid.Parse(id)
		if err!= nil{
			fmt.Println(err)
		}		
		product, err := a.db.GetProductByID(c.Request().Context(), id_)
		if err!= nil{
			fmt.Println(err)
			return a.RenderView(c, views.Product(database.Product{}), "/product/1")
		}
		return a.RenderView(c, views.Product(product), fmt.Sprintf("/product/%s", id_.String()))
	}
}
