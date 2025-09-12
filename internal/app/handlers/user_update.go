package handlers

import (
	"net/http"
	"regexp"

	"github.com/yoru0/goapi.git/internal/pkg/common/api"
	"github.com/yoru0/goapi.git/internal/pkg/data/dao"
	"github.com/yoru0/goapi.git/internal/pkg/models"
)

type UserUpdateRequestParam struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserUpdateResponseParam struct {
	api.ResponseData
	User *models.User `json:"user"`
}

func (param *UserUpdateRequestParam) Validate() (msg, field string) {
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

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	var param UserUpdateRequestParam
	if !api.DecodeBodyJSON(w, r, &param) {
		return
	}

	user, err := dao.NewUserDAO().Update(param.ID, &models.User{
		Name:  param.Name,
		Email: param.Email,
	})
	if err != nil {
		response := api.NewAPIResponseWithError("INTERNAL_ERROR", "Failed to update user")
		api.SendResponseJSONWithStatusCode(w, response, http.StatusInternalServerError)
		return
	}

	data := UserUpdateResponseParam{
		User: user,
	}
	response := api.NewAPIResponse()
	response.Data = data
	api.SendResponseJSON(w, response)
}
