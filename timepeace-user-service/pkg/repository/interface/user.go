package interfaces

import (
	"github.com/sgokul961/timepeace-user-service/pkg/models"
	"github.com/sgokul961/timepeace-user-service/pkg/pb"
)

type UserRepository interface {
	CheckUserAvailability(email string) bool
	FindByEmail(email string) (models.UserLoginCheck, error)
	Save(user *pb.SignUpRequest) (*pb.UserDetails, error)
	IsBlocked(email string) (bool, error)
}
