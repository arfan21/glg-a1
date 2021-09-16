package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/arfan21/hacktiv8-golang-swagger/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Orders API
// @version 1.0
// @description this is sample api for orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email test@test.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	router := mux.NewRouter()

	// Get Orders
	// router.HandleFunc("/orders", GetOrders).Methods("GET")
	// router.HandleFunc("/orders", PostOrder).Methods("POST")

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	log.Fatal(http.ListenAndServe(":8000", router))
}

type Order struct {
	OrderID      int       `json:"order_id" example:"1"`
	CustomerName string    `json:"customerName" example:"test Customer"`
	OrderedAt    time.Time `json:"orderedAt" example:"2021-09-16T12:51:56.558Z"`
	Items        []Item    `json:"items"`
}

type Item struct {
	ItemID      int    `json:"item_id" example:"1"`
	Description string `json:"description" example:"test Description"`
	Quantity    int    `json:"quantity" example:"1"`
}

type OrderResponsePost struct {
	ID int `json:"id" example:"1"`
}

// GetOrders godoc
// @summary get all orders
// @description get all of orders
// @tags orders
// @accept json
// @produce json
// @success 200 {array} Order
// @router /orders [get]
func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	jsobByte, _ := json.Marshal([]Order{})
	w.Write(jsobByte)
}

// PostOrder godoc
// @summary create order
// @description create order
// @tags orders
// @accept json
// @produce json
// @param order body Order true "Create Order"
// @success 200 {object} OrderResponsePost
// @router /orders [post]
func PostOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	jsobByte, _ := json.Marshal([]Order{})
	w.Write(jsobByte)
}
