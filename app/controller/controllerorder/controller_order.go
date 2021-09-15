package controllerorder

import (
	"net/http"
	"strconv"

	"github.com/arfan21/hacktiv8-assignment-2/app/model/modelorder"
	"github.com/arfan21/hacktiv8-assignment-2/app/model/modelweb"
	"github.com/arfan21/hacktiv8-assignment-2/app/service/serviceorder"
	"github.com/arfan21/hacktiv8-assignment-2/helper"
	"github.com/julienschmidt/httprouter"
)

type Controller interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type controller struct {
	srv serviceorder.Service
}

func New(srv serviceorder.Service) Controller {
	return &controller{srv: srv}
}

func (c *controller) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	orderRequest := modelorder.OrderRequest{}
	helper.ReadBody(r, &orderRequest)

	orderResponse := c.srv.Create(r.Context(), orderRequest)

	res := modelweb.Response{
		Status: "CREATED",
		Data:   orderResponse,
	}

	helper.WriteJSON(w, http.StatusCreated, res)
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ordersResponse := c.srv.GetAll(r.Context())
	res := modelweb.Response{
		Status: "SUCCESS",
		Data:   ordersResponse,
	}

	helper.WriteJSON(w, http.StatusOK, res)
}

func (c *controller) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	orderRequest := modelorder.OrderRequest{}
	helper.ReadBody(r, &orderRequest)
	ordersResponse := c.srv.Update(r.Context(), orderRequest)
	res := modelweb.Response{
		Status: "SUCCESS",
		Data:   ordersResponse,
	}

	helper.WriteJSON(w, http.StatusOK, res)
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	orderIdString := params.ByName("orderId")
	orderId, err := strconv.Atoi(orderIdString)
	helper.PanicIfNeeded(err)

	c.srv.Delete(r.Context(), orderId)

	res := modelweb.Response{
		Status: "SUCCESS",
	}

	helper.WriteJSON(w, http.StatusOK, res)
}
