package serviceorder

import (
	"context"

	"github.com/arfan21/hacktiv8-assignment-2/app/model/modelorder"
	"github.com/arfan21/hacktiv8-assignment-2/app/repository/mysql/mysqlorder"
)

type Service interface {
	Create(ctx context.Context, reqOrder modelorder.OrderRequest) modelorder.OrderResponse
	GetAll(ctx context.Context) []modelorder.OrderResponse
	Update(ctx context.Context, reqOrder modelorder.OrderRequest) modelorder.Order
	Delete(ctx context.Context, orderId int)
}

type service struct {
	repo mysqlorder.Repository
}

func New(repo mysqlorder.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, reqOrder modelorder.OrderRequest) modelorder.OrderResponse {
	newOrder := modelorder.Order{}
	newOrder.SetFromRequest(reqOrder)

	createdOrder := s.repo.Create(ctx, newOrder)

	newOrderResponse := modelorder.OrderResponse{}
	newOrderResponse.Set(createdOrder)

	return newOrderResponse
}

func (s *service) GetAll(ctx context.Context) []modelorder.OrderResponse {
	orders := s.repo.GetAll(ctx)
	ordersResponse := []modelorder.OrderResponse{}

	for _, order := range orders {
		orderResponse := modelorder.OrderResponse{}
		orderResponse.Set(order)
		ordersResponse = append(ordersResponse, orderResponse)
	}

	return ordersResponse
}

func (s *service) Update(ctx context.Context, reqOrder modelorder.OrderRequest) modelorder.Order {
	order := modelorder.Order{}
	order.SetFromRequest(reqOrder)
	updatedOrder := s.repo.Update(ctx, order)
	return updatedOrder
}

func (s *service) Delete(ctx context.Context, orderId int) {
	s.repo.Delete(ctx, orderId)
}
