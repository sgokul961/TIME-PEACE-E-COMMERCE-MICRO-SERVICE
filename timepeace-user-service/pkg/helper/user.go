package helper

import (
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sgokul961/timepeace-user-service/pkg/pb"
)

type CustomUserClaim struct {
	ID    uint64
	Email string
	Role  string
	jwt.StandardClaims
}

func GenerateUserToken(user *pb.UserDetails) (string, error) {
	claims := &CustomUserClaim{
		ID:    user.Id,
		Email: user.Email,
		Role:  "user",
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

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}
func IsValidPhoneNumber(phoneNumber string) bool {
	// Define a regex pattern for the format "+919847256365" (plus sign and 12 digits)
	pattern := `^\+\d{12}$`
	match, _ := regexp.MatchString(pattern, phoneNumber)
	return match
}
