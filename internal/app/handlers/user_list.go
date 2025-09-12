package handlers

import (
	"net/http"

	"github.com/yoru0/goapi.git/internal/pkg/common/api"
	"github.com/yoru0/goapi.git/internal/pkg/data/dao"
	"github.com/yoru0/goapi.git/internal/pkg/models"
)

type UserListResponseData struct {
	api.ResponseData
	Users []*models.User `json:"users"`
}

func UserList(w http.ResponseWriter, r *http.Request) {
	users, err := dao.NewUserDAO().GetAll()
	if err != nil {
		response := api.NewAPIResponseWithError("INTERNAL_ERROR", "Failed to retrieve users")
		api.SendResponseJSONWithStatusCode(w, response, http.StatusInternalServerError)
		return
	}

	data := UserListResponseData{
		Users: users,
	}
	response := api.NewAPIResponse()
	response.Data = data
	api.SendResponseJSON(w, response)
}
