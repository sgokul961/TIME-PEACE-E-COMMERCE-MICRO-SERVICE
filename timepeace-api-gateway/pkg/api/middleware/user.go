package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sgokul961/timepeace-api-gateway/pkg/utils/response"
)

type CoustomUserClaimn struct {
	Id    uint
	Email string
	Role  string
	jwt.StandardClaims
}

func UserAuthorizationMiddleWare(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	var tokenString string

	if strings.HasPrefix(s, "Bearer") {
		tokenString = strings.TrimPrefix(s, "Bearer")
	} else {
		tokenString = s
	}

	token, err := validateUserToken(tokenString)

	if err != nil || !token.Valid {
		errRes := response.Responses(http.StatusUnauthorized, "not authorized", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	claims, ok := token.Claims.(*CoustomUserClaimn)

	if !ok {
		errRes := response.Responses(http.StatusUnauthorized, "not authorized", nil, fmt.Errorf("calim not retrived").Error())
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	id := claims.Id
	c.Set("userId", id)
	c.Next()
}
func validateUserToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CoustomUserClaimn{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	return token, err
}
