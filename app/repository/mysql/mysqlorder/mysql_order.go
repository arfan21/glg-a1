package mysqlorder

import (
	"context"

	"github.com/arfan21/hacktiv8-assignment-2/app/model/modelorder"
	"github.com/arfan21/hacktiv8-assignment-2/config/mysql"
	"github.com/arfan21/hacktiv8-assignment-2/helper"
)

type Repository interface {
	Create(ctx context.Context, order modelorder.Order) modelorder.Order
	GetAll(ctx context.Context) []modelorder.Order
	Update(ctx context.Context, order modelorder.Order) modelorder.Order
	Delete(ctx context.Context, orderId int)
}

type repository struct {
	client mysql.Client
}

func New(client mysql.Client) Repository {
	return &repository{client}
}

func (r *repository) Create(ctx context.Context, order modelorder.Order) modelorder.Order {
	err := r.client.Conn().WithContext(ctx).Create(&order).Error
	helper.PanicIfNeeded(err)

	return order
}

func (r *repository) GetAll(ctx context.Context) []modelorder.Order {
	orders := []modelorder.Order{}
	err := r.client.Conn().WithContext(ctx).Preload("Items").Find(&orders).Error
	helper.PanicIfNeeded(err)

	return orders
}

func (r *repository) Update(ctx context.Context, order modelorder.Order) modelorder.Order {
	err := r.client.Conn().WithContext(ctx).Updates(&order).Error
	helper.PanicIfNeeded(err)

	return order
}

func (r *repository) Delete(ctx context.Context, orderId int) {
	err := r.client.Conn().WithContext(ctx).Delete(&modelorder.Order{}, orderId).Error
	helper.PanicIfNeeded(err)
}
