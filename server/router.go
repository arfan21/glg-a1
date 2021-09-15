package server

import (
	"github.com/arfan21/hacktiv8-assignment-2/app/controller/controllerorder"
	"github.com/arfan21/hacktiv8-assignment-2/app/repository/mysql/mysqlorder"
	"github.com/arfan21/hacktiv8-assignment-2/app/service/serviceorder"
	"github.com/arfan21/hacktiv8-assignment-2/config/mysql"
	"github.com/arfan21/hacktiv8-assignment-2/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(mysqlClient mysql.Client) *httprouter.Router {
	r := httprouter.New()
	r.PanicHandler = exception.ErrorHandler

	repoOrder := mysqlorder.New(mysqlClient)
	srv := serviceorder.New(repoOrder)
	ctrlOrder := controllerorder.New(srv)
	r.POST("/", ctrlOrder.Create)
	r.GET("/orders", ctrlOrder.GetAll)
	r.PUT("/orders", ctrlOrder.Update)
	r.DELETE("/orders/:orderId", ctrlOrder.Delete)

	return r
}
