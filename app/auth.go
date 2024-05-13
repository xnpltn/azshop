package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xnpltn/azshop/database"
	"github.com/xnpltn/azshop/utls"
)

var ErrDuplicate = errors.New(`pq: duplicate key value violates unique constraint "unique_email"`)

func (a *Application) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		user, err := a.db.GetUserByEmail(c.Request().Context(), email)
		if err != nil {
			return c.HTML(http.StatusOK, "Error")
		}
		token, err := utls.NewToken(user)
		if err != nil {
			return c.HTML(http.StatusOK, "something went wrong")
		}
		if !utls.CheckPasswordHash(password, user.Password) {
			return c.HTML(http.StatusOK, "Invalid Credentials")
		}
		c.SetCookie(&http.Cookie{
			Name:     "auth_token",
			Value:    token,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			HttpOnly: true,
			MaxAge:   3600,
			Path:     "/",
		})

		return c.HTML(
			http.StatusOK,
			`
			<span class="text-center text-blue-700">
				SUCCESS | click <a href="/" class="underline text-red-500">here</a> to get home,
			</span>
		`,
		)
	}
}

func (a *Application) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.FormParams())
		if c.FormValue("password-s") != c.FormValue("password-s-1") {
			return c.HTML(200, `
			<span class="text-center text-blue-700">
				Passwordss Should be equal,
			</span>
			`)
		}

		pass, err := utls.HashPassword(c.FormValue("password-s"))
		if err != nil {
			return c.HTML(200, `
				<span class="text-center text-blue-700">
					something went wrong
				</span>`)
		}

		err = a.db.CreateUser(c.Request().Context(), database.CreateUserParams{
			Name:     c.FormValue("name-s"),
			Email:    c.FormValue("email-s"),
			Password: pass,
		})

		if err != nil {
			if err.Error() == ErrDuplicate.Error() {
				return c.HTML(200, `
				<span class="text-center text-blue-700">
					Email already exit
				</span>`)
			} else {
				return c.HTML(200, `
				<span class="text-center text-blue-700">
					something went wrong
				</span>`)
			}
		}

		return c.HTML(200, `
			<span class="text-center text-blue-700">
				SUCCESS | click <a href="/login" class="underline text-red-500">here</a> to login,
			</span>
		`)
	}
}

func (a *Application) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("HX-Redirect", "/")
		c.SetCookie(&http.Cookie{
			Name:     "auth_token",
			Value:    "",
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			HttpOnly: true,
			MaxAge:   3600,
			Path:     "/",
		})
		return c.HTML(http.StatusOK, "Log Out")
	}
}
