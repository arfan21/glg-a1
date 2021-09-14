package exception

import (
	"net/http"

	// "github.com/arfan21/hacktiv8-assignment-2/app/model/modelweb"
	"github.com/arfan21/hacktiv8-assignment-2/app/model/modelweb"
	"github.com/arfan21/hacktiv8-assignment-2/helper"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	res := modelweb.Response{
		Status: "ERROR",
		Data:   err,
	}

	helper.WriteJSON(w, http.StatusInternalServerError, res)
}
