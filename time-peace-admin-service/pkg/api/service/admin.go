package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sgokul961/time-peace-admin-service/pkg/helper"
	"github.com/sgokul961/time-peace-admin-service/pkg/pb"
	interfacees "github.com/sgokul961/time-peace-admin-service/pkg/repository/interface"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	Repo interfacees.AdminRepository
	pb.UnimplementedAdminServiceServer
}

func NewAdminService(repo interfacees.AdminRepository) *AdminService {
	return &AdminService{Repo: repo}
}
func (a *AdminService) Login(ctx context.Context, admin *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {

	if ok := a.Repo.CheckAdminAvailability(admin.Email); !ok {
		fmt.Println("admindata is:", admin)

		return nil, errors.New("user does not exist")

	}
	fmt.Println("the admin data is :", admin)
	adminCompare, err := a.Repo.FindByEmail(admin.Email)
	fmt.Println("admin email is :", admin.Email)
	if err != nil {
		return nil, errors.New("error fetching userdata")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(adminCompare.Password), []byte(admin.Password)); err != nil {
		return nil, errors.New("password is incorrect")

	}
	adminDetails := &pb.AdminDetails{Id: adminCompare.ID, Name: adminCompare.Name, Email: adminCompare.Email, Phno: adminCompare.PhNo}
	fmt.Println("admin chek 2:", adminDetails)
	tokenString, err := helper.GenerateAdminToken(adminDetails)

	if err != nil {
		return nil, err
	}
	return &pb.AdminLoginResponse{Admindetails: adminDetails, Token: tokenString}, nil
}
