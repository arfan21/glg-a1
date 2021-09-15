package exception

import (
	"log"
	"net/http"
	"strings"

	"github.com/arfan21/hacktiv8-assignment-2/app/model/modelweb"
	"github.com/arfan21/hacktiv8-assignment-2/helper"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println(err)
	if clientError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func clientError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	w.Header().Set("Content-Type", "application/json")

	res := modelweb.Response{
		Status: "ERROR",
	}

	if newErr, ok := err.(error); ok && strings.Contains(newErr.Error(), "invalid character") {
		res.Data = newErr.Error()
		helper.WriteJSON(w, http.StatusBadRequest, res)
		return true
	}

	return false
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")

	res := modelweb.Response{
		Status: "ERROR",
	}

	newErr, ok := err.(error)
	if ok {
		res.Data = newErr.Error()
		helper.WriteJSON(w, http.StatusInternalServerError, res)
	} else {
		res.Data = err
		helper.WriteJSON(w, http.StatusInternalServerError, res)
	}

}
