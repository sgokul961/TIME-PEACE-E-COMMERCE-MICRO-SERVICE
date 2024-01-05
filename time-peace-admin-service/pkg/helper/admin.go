package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sgokul961/time-peace-admin-service/pkg/pb"
)

type CustomAdminClaim struct {
	ID    uint64
	Email string
	Role  string
	jwt.StandardClaims
}

func GenerateAdminToken(admin *pb.AdminDetails) (string, error) {
	claims := &CustomAdminClaim{
		ID:    admin.Id,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
