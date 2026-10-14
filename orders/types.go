package main

import (
	"context"
	pb "github.com/anand-gec/menu/api"
)

type OrdersService interface {
	CreateOrder(context.Context) error //for payload output
	validateOrder(context.Context,*pb.CreateOrderRequest)error
}

type OrderStore interface {
	Create(context.Context) error
}
