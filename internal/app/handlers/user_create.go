package handlers

import (
	"net/http"
	"regexp"

	"github.com/yoru0/goapi.git/internal/pkg/common/api"
	"github.com/yoru0/goapi.git/internal/pkg/data/dao"
	"github.com/yoru0/goapi.git/internal/pkg/models"
)

type UserCreateRequestParam struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCreateResponseParam struct {
	api.ResponseData
	Success bool        `json:"success"`
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

func (param *UserCreateRequestParam) Validate() (msg, field string) {
	if param.Name == "" {
		return "Name is required", "name"
	}
	if param.Email == "" {
		return "Email is required", "email"
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`).MatchString(param.Email) {
		return "Email format is invalid", "email"
	}
	return "", ""
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var param UserCreateRequestParam
	if !api.DecodeBodyJSON(w, r, &param) {
		return
	}

	user := &models.User{
		Name:  param.Name,
		Email: param.Email,
	}

	createdUser, err := dao.NewUserDAO().Create(user)
	if err != nil {
		response := api.NewAPIResponseWithError("INTERNAL_ERROR", "Failed to create user")
		api.SendResponseJSONWithStatusCode(w, response, http.StatusInternalServerError)
		return
	}

	data := UserCreateResponseParam{
		Success: true,
		Message: "User created successfully",
		User:    *createdUser,
	}
	response := api.NewAPIResponse()
	response.Data = data
	api.SendResponseJSON(w, response)
}
