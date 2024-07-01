package types

import (
	"context"
	"github.com/emaanmohamed/order-management-system/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
