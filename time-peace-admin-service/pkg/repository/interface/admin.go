package interfacees

import "github.com/sgokul961/time-peace-admin-service/pkg/models"

type AdminRepository interface {
	CheckAdminAvailability(email string) bool
	FindByEmail(email string) (models.LoginCheck, error)
}
