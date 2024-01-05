package repository

import (
	"errors"
	"fmt"

	"github.com/sgokul961/time-peace-admin-service/pkg/models"
	interfacees "github.com/sgokul961/time-peace-admin-service/pkg/repository/interface"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfacees.AdminRepository {
	return &AdminRepository{db}
}
func (c *AdminRepository) CheckAdminAvailability(email string) bool {
	var count int

	err := c.DB.Raw(`SELECT COUNT(*) FROM users WHERE email=? AND role ='admin'`, email).Scan(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (c *AdminRepository) FindByEmail(email string) (models.LoginCheck, error) {
	var adminDetails models.LoginCheck

	err := c.DB.Raw(`SELECT id,name,email,ph_no,password FROM users WHERE email=? AND role='admin'`, email).Scan(&adminDetails).Error
	fmt.Println("the repo admin details id :", adminDetails)
	if err != nil {
		return models.LoginCheck{}, errors.New("error in fetching admindetails")
	}
	return adminDetails, nil

}
