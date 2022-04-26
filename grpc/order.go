package grpc

import (
	"fmt"
	errorD "gRPC/errors"
	models "gRPC/model"
	pb "gRPC/protos"
	"golang.org/x/net/context"
)

const (
	Insert = "Insert"
	Delete = "Delete"
)

type OrderServer struct{}

func (s *OrderServer) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	OrderResp := pb.CreateOrderResponse{}
	OrderResp.Err = errorD.Success
	OrderResp.Msg = "创建成功"
	Havings := make([]models.Oh,0)
	HavingIds := make([]int,len(in.Havings))
	if len(in.Havings) == 0 {
		OrderResp.Err = errorD.Fail
		OrderResp.Msg = "没有选择商品"
	}
	price := float64(0)
	for i := 0; i < len(in.Havings); i++ {
		Havings[0].HavingId = int(in.Havings[i].HavingId)
		price += float64(in.Havings[i].HavingPrice)
		HavingIds[i] = int(in.Havings[i].HavingId)
	}
	fmt.Printf("%v",HavingIds)
	order := &models.OrderJson{
		UserId:          int(in.UserId),
		OrderTag:        in.OrderTag,
		OrderPrice:      price,
		OrderHavingId:   HavingIds,
	}
	err := models.CreateOrder(order)
	if err != nil {
		OrderResp.Err = errorD.Fail
		OrderResp.Msg = "创建失败"
	}
	return &OrderResp, nil
}
func (s *OrderServer) UpdateOrder(ctx context.Context, in *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	OrderResp := pb.UpdateOrderResponse{}
	OrderResp.Err = errorD.Success
	OrderResp.Msg = "成功"
	if in.OrderStatus == Insert{
		err := models.CreateOh(int(in.OrderId),in.Having)
		if err != nil {
			OrderResp.Err = errorD.Fail
			OrderResp.Msg = "失败"
		}
	}else if in.OrderStatus == Delete {
		err := models.DeleteOh(int(in.OrderId),in.Having)
		if err != nil {
			OrderResp.Err = errorD.Fail
			OrderResp.Msg = "失败"
		}
	}
	return &OrderResp, nil
}
func (s *OrderServer) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	OrderResp := pb.GetOrderResponse{}
	OrderResp.Err = errorD.Success
	OrderResp.Msg = "成功"
	order := &models.Order{
		OrderId:         int(in.OrderId),
		UserId:          int(in.UserId),
		OrderStatus:     in.OrderStatus,
	}
	Orders ,err := models.GetOrder(int(in.Limit),int(in.Page),order)
	if err != nil {
		OrderResp.Err = errorD.Fail
		OrderResp.Msg = "失败"
	}
	for _, o := range Orders {
		temp := pb.Order{
			OrderId:         int32(o.OrderId),
			UserId:          int32(o.UserId),
			OrderTag:        o.OrderTag,
			OrderPrice:      float32(o.OrderPrice),
			OrderCreateTime: o.OrderCreateTime,
			OrderUpdateTime: o.OrderUpdateTime,
			OrderStatus:     o.OrderStatus,
			OrderRemark:     o.OrderRemark,
		}
		OrderResp.Data = append(OrderResp.Data,&temp)
	}
	return &OrderResp, nil
}
