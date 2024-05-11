package app

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/xnpltn/azshop/database"
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
		if db_user.Password != password {
			return c.HTML(200, "Invalid Credentials")
		}
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
		password := c.FormValue("password")
		passwordConfirm := c.FormValue("password1")
		if password != passwordConfirm {
			return c.HTML(200, `
			<span class="text-center text-blue-700">
				Passwordss Should be equal,
			</span>
			`)
		}

		err := a.db.CreateUser(c.Request().Context(), database.CreateUserParams{
			Name:     c.FormValue("name"),
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		})

		if err != nil {
			if err.Error() == ErrDuplicate.Error(){
				return c.HTML(200, `
				<span class="text-center text-blue-700">
					Email already exit
				</span>`)
			}else {
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
