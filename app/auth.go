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
		db_user, err := a.db.GetUserByEmail(c.Request().Context(), email)
		if err != nil {
			fmt.Println(err)
			return c.HTML(200, "Error")
		}
		
		if !utls.CheckPasswordHash(password, db_user.Password) {
			return c.HTML(200, "Invalid Credentials")
		}
		c.SetCookie(&http.Cookie{
			Name:     "auth_token",
			Value:    email,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			HttpOnly: true,
			MaxAge:   3600,
			Path:     "/",
		})

		return c.HTML(
			200,
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
		if c.FormValue("password") != c.FormValue("password1") {
			return c.HTML(200, `
			<span class="text-center text-blue-700">
				Passwordss Should be equal,
			</span>
			`)
		}

		pass, err := utls.HashPassword(c.FormValue("password"))
		if err != nil {
			return c.HTML(200, `
				<span class="text-center text-blue-700">
					something went wrong
				</span>`)
		}

		err = a.db.CreateUser(c.Request().Context(), database.CreateUserParams{
			Name:     c.FormValue("name"),
			Email:    c.FormValue("email"),
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
		c.SetCookie(&http.Cookie{
			Name:     "auth_token",
			Value:    "",
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			HttpOnly: true,
			MaxAge:   3600,
			Path:     "/",
		})
		return c.HTML(200, "Log Out")
	}
}
