package app

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
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
		}
		fmt.Println(products[0].ImageUrl)
		return a.RenderView(c, views.HomeView("Home", products), "/")
	}

}

func(a *Application) AboutGetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return a. RenderView(c, views.AboutView("About"), "/about")
	}

}


func(a *Application) Admin()echo.HandlerFunc{
	return func(c echo.Context) error {
		return a.RenderView(c, views.AdminView("Admin"), "/admin")
	}
}
