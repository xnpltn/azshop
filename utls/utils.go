package utls

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/xnpltn/azshop/database"
	"golang.org/x/crypto/bcrypt"
)



func CheckCredentials(c echo.Context, db *database.Queries) (*database.User, error) {
	authCookie, err := c.Cookie("auth_token")
	if err != nil {
		return nil, err
	}
	userJWT, err := VerifyJWT(authCookie.Value)
	if err!= nil{
		return nil, err
	}
	_, err = db.GetUserByID(c.Request().Context(), userJWT.ID)
	if err != nil {
		return nil, err
	}
	return userJWT, nil

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

func NewToken(user database.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyJWT(token string) (*database.User, error) {
	parsedToken, err := jwt.Parse(string(token), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	id, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		return nil, err
	}
	user := &database.User{
		ID:   id,
		Name: claims["name"].(string),
	}
	return user, nil
}
