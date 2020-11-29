package util

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func UsernameFromToken(c echo.Context) (string, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	if username == "" {
		return "", errors.New("failed to get username from token claims")
	}
	return username, nil
}
