package utls

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/xnpltn/azshop/database"
	"golang.org/x/crypto/bcrypt"
)



func CheckCredentials(c echo.Context, db *database.Queries) ( *database.User, bool){
	authCookie, err := c.Cookie("auth_token")
	if err!= nil{
		fmt.Println(err)
		return nil, false
	}

	user, err := db.GetUserByEmail(c.Request().Context(), authCookie.Value)
	if err != nil{
		fmt.Println(err)
		return nil, false
	}
	return &user, true

}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}