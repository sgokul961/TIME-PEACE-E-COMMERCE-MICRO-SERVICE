package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sgokul961/timepeace-user-service/pkg/helper"
	"github.com/sgokul961/timepeace-user-service/pkg/pb"
	interfaces "github.com/sgokul961/timepeace-user-service/pkg/repository/interface"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo interfaces.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{Repo: repo}
}
func (c *UserService) SignUp(ctx context.Context, user *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	fmt.Println("signupservice")
	isValidEmail := helper.IsValidEmail(user.Email)
	if !isValidEmail {
		return nil, errors.New("please enter a valid phone number")
	}
	isValidPhoneNumber := helper.IsValidPhoneNumber(user.Phno)

	if !isValidPhoneNumber {
		return nil, errors.New("please enter a valid phone number")
	}
	if ok := c.Repo.CheckUserAvailability(user.Email); ok {
		return nil, errors.New("already existing email")
	}
	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return nil, errors.New("error in password hashing")
	}
	user.Password = string(hashPassWord)

	userDetails, err := c.Repo.Save(user)
	if err != nil {
		return nil, errors.New("error in saving user data")
	}
	tokenString, err := helper.GenerateUserToken(userDetails)

	if err != nil {
		return nil, err
	}
	return &pb.SignUpResponse{Userdetails: userDetails, Token: tokenString}, nil
}
func (c *UserService) Login(ctx context.Context, user *pb.LoginRequest) (*pb.LoginResponse, error) {
	if ok := c.Repo.CheckUserAvailability(user.Email); !ok {
		return nil, errors.New("no such user exist")

	}
	if ok, _ := c.Repo.IsBlocked(user.Email); ok {
		return nil, errors.New("user is blocked")
	}

	userCompare, err := c.Repo.FindByEmail(user.Email)

	if err != nil {
		return nil, errors.New("error in fetching userdata")
	}

	if err != nil {
		return nil, errors.New("error in fetching userdata")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userCompare.Password), []byte(user.Password)); err != nil {
		return nil, errors.New("password is incorrect")

	}
	userDetails := &pb.UserDetails{Id: userCompare.ID, Name: userCompare.Name, Email: userCompare.Email, Phno: userCompare.PhNo}

	tokenString, err := helper.GenerateUserToken(userDetails)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Userdetails: userDetails, Token: tokenString}, nil
}
