package users

import (
	"github.com/gin-gonic/gin"
)

type UserValidator struct {
	UserFields struct {
		ID int `json:"id" binding:"numeric,omitempty"`
		Name string `json:"name" binding:"required_without=ID,alphanum,min=1,max=255"`
		Surname string `json:"surname" binding:"required_without=ID,alphanum,min=1,max=255"`
		Email string `json:"email" binding:"required_without=ID,email"`
	} `json:"user"`
	user UserModel `json:"-"`
}

func (self *UserValidator) Bind(ctx *gin.Context) error {
	err := ctx.ShouldBindJSON(self)
	if err != nil {
		return err
	}
	if self.UserFields.ID != 0 {
		self.user.ID = self.UserFields.ID
	}
	self.user.Name = self.UserFields.Name
	self.user.Surname = self.UserFields.Surname
	self.user.Email = self.UserFields.Email
	return nil
}

func CreateUserValidator() UserValidator {
	return UserValidator{}
}

func UpdateUserValidator(user UserModel) UserValidator {
	validator := UserValidator{}
	validator.UserFields.ID = user.ID
	validator.UserFields.Name = user.Name
	validator.UserFields.Surname = user.Surname
	validator.UserFields.Email = user.Email
	return validator
}
