package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sgokul961/timepeace-api-gateway/pkg/utils/response"
)

type CoustomAdminClaim struct {
	ID    uint
	Email string
	Role  string
	jwt.StandardClaims
}

func AdminAuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	var tokenString string
	if strings.HasPrefix(s, "Bearer") {
		tokenString = strings.TrimPrefix(s, "Bearer")
	} else {
		tokenString = s
	}

	token, err := validateToken(tokenString)

	if err != nil || !token.Valid {
		errRes := response.Responses(http.StatusUnauthorized, "not authorized", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(*CoustomAdminClaim)

	if !ok {
		errRes := response.Responses(http.StatusUnauthorized, "not authorized", nil, fmt.Errorf("cliam not retrived"))
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	if claims.Role == "user" {
		errRes := response.Responses(http.StatusUnauthorized, "no admin toekn ", nil, errors.New("access restricted to only admin"))
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	id := claims.ID
	c.Set("adminId", id)
	c.Next()
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CoustomAdminClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	return token, err
}
