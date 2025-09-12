package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yoru0/goapi.git/internal/pkg/common/api"
	"github.com/yoru0/goapi.git/internal/pkg/data/dao"
)

type UserDeleteRequestParam struct {
	ID string `json:"id"`
}

type UserDeleteResponseParam struct {
	api.ResponseData
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	var param UserDeleteRequestParam
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		msg := "Request body format is invalid"
		response := api.NewAPIResponseWithError("INVALID_PARAM", msg)
		api.SendResponseJSONWithStatusCode(w, response, http.StatusBadRequest)
		return
	}

	err = dao.NewUserDAO().Delete(param.ID)
	if err != nil {
		response := api.NewAPIResponseWithError("INTERNAL_ERROR", "Failed to delete user")
		api.SendResponseJSONWithStatusCode(w, response, http.StatusInternalServerError)
		return
	}

	data := UserDeleteResponseParam{
		Success: true,
		Message: "User deleted successfully",
	}
	response := api.NewAPIResponse()
	response.Data = data
	api.SendResponseJSON(w, response)
}
