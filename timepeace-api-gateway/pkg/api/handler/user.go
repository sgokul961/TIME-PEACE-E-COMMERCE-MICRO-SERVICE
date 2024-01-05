package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sgokul961/timepeace-api-gateway/pkg/client"
	"github.com/sgokul961/timepeace-api-gateway/pkg/pb"
	"github.com/sgokul961/timepeace-api-gateway/pkg/utils/response"
)

type UserHandler struct {
	Client *client.UserClient
}

func NewUserHandler(c *client.UserClient) *UserHandler {
	return &UserHandler{c}
}
func (ur *UserHandler) SignUp(c *gin.Context) {

	var signupDetails *pb.SignUpRequest
	if err := c.ShouldBindJSON(&signupDetails); err != nil {
		errRes := response.Response{Statuscode: http.StatusBadRequest, Message: "constrains not satisfied", Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, errRes)
		return

	}
	if err := validator.New().Struct(signupDetails); err != nil {
		errRes := response.Response{Statuscode: http.StatusBadRequest, Message: "constrains not satisfied", Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userdata, err := ur.Client.Client.SignUp(context.Background(), signupDetails)
	fmt.Println("op:", userdata)
	if err != nil {
		erRes := response.Response{Statuscode: http.StatusInternalServerError, Message: "internal server error", Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, erRes)
		return
	}

	succecRes := response.Responses(http.StatusCreated, "successfully signed in", userdata, nil)
	c.JSON(http.StatusCreated, succecRes)

}
func (ur *UserHandler) Login(c *gin.Context) {
	var loginDetails *pb.LoginRequest
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are  provided in the bad format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := validator.New().Struct(loginDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userData, err := ur.Client.Client.Login(context.Background(), loginDetails)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	sucRes := response.Responses(http.StatusOK, "successfully logged in", userData, nil)
	c.JSON(http.StatusOK, sucRes)
}
