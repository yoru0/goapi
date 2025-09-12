package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yoru0/goapi.git/internal/pkg/common/api"
	"github.com/yoru0/goapi.git/internal/pkg/data/dao"
	"github.com/yoru0/goapi.git/internal/pkg/models"
)

type UserGetRequestParam struct {
	ID string `json:"id"`
}

type UserGetResponseParam struct {
	api.ResponseData
	User *models.User `json:"user"`
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	var param UserGetRequestParam
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		msg := "Request body format is invalid"
		response := api.NewAPIResponseWithError("INVALID_PARAM", msg)
		api.SendResponseJSONWithStatusCode(w, response, http.StatusBadRequest)
		return
	}

	user, err := dao.NewUserDAO().GetByID(param.ID)
	if err != nil {
		response := api.NewAPIResponseWithError("USER_NOT_FOUND", err.Error())
		api.SendResponseJSONWithStatusCode(w, response, http.StatusNotFound)
		return
	}

	data := &UserGetResponseParam{
		User: user,
	}
	response := api.NewAPIResponse()
	response.Data = data
	api.SendResponseJSON(w, response)
}
