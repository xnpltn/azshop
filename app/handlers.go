package app

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/xnpltn/azshop/database"
	"github.com/xnpltn/azshop/templates"
	"github.com/xnpltn/azshop/templates/views"
	"github.com/xnpltn/azshop/utls"
)

func (a *Application) RenderView(c echo.Context, view templ.Component, layoutPath string) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return view.Render(c.Request().Context(), c.Response().Writer)
	}
	_, err := utls.CheckCredentials(c, a.db)
	if err != nil {
		return templates.Layout(layoutPath, false).Render(c.Request().Context(), c.Response().Writer)
	}
	return templates.Layout(layoutPath, true).Render(c.Request().Context(), c.Response().Writer)
}

func (a *Application) HomeGetHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		products, err := a.db.GetAllProducts(c.Request().Context())
		if err != nil {
			return a.RenderView(c, views.HomeView("Home", []database.Product{}, false), "/")
		}
		_, err = utls.CheckCredentials(c, a.db)
		if err != nil {
			return a.RenderView(c, views.HomeView("Home", products, false), "/")
		}
		return a.RenderView(c, views.HomeView("Hello", products, true), "/")
	}

}

func (a *Application) AboutGetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return a.RenderView(c, views.AboutView("About"), "/about")
	}

}

func (a *Application) LoginGetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return a.RenderView(c, views.LoginView("Login"), "/login")
	}

}

func (a *Application) Admin() echo.HandlerFunc {
	return func(c echo.Context) error {
		products, err := a.db.GetAllProducts(c.Request().Context())
		if err != nil {
			return a.RenderView(c, views.AdminView("Admin", []database.Product{}), "/admin")
		}
		return a.RenderView(c, views.AdminView("Admin", products), "/admin")
	}
}

func (a *Application) Product() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		id_, err := uuid.Parse(id)
		if err != nil {
			fmt.Println(err)
		}
		product, err := a.db.GetProductByID(c.Request().Context(), id_)
		if err != nil {
			return a.RenderView(c, views.Product(database.Product{}), "/product/1")
		}
		return a.RenderView(c, views.Product(product), fmt.Sprintf("/product/%s", id_.String()))
	}
}

func (a *Application) Cart() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, _ := utls.CheckCredentials(c, a.db)
		cart, err := a.db.GetUserCart(c.Request().Context(), user.ID)
		if err != nil {
			fmt.Println(err)
			return a.RenderView(c, views.Cart([]database.GetUserCartRow{}), "/cart")
		}
		return a.RenderView(c, views.Cart(cart), "/cart")
	}
}

func (a *Application) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userJWT, _ := utls.CheckCredentials(c, a.db)
		user, _ := a.db.GetUserByID(c.Request().Context(), userJWT.ID)
		return a.RenderView(c, views.ProfileView("Profile", user), "/profile")
	}
}
