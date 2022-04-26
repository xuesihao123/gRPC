package grpc

import (
	errorD "gRPC/errors"
	models "gRPC/model"
	pb "gRPC/protos"
	"golang.org/x/net/context"
)

type MenuServer struct{}

func (s *MenuServer) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	MenuResp := pb.CreateMenuResponse{}
	MenuResp.Err = errorD.Success
	MenuResp.Msg = "创建成功"
	err := models.CreateMenu(in.HavingIds)
	if err != nil {
		MenuResp.Err = errorD.Fail
		MenuResp.Msg = "创建失败"
	}
	return &MenuResp, nil
}
func (s *MenuServer) DeleteMenu(ctx context.Context, in *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {
	MenuResp := pb.DeleteMenuResponse{}
	MenuResp.Err = errorD.Success
	MenuResp.Msg = "成功"
	err := models.DeleteMh(int(in.MenuId),in.HavingIds)
	if err != nil {
		MenuResp.Err = errorD.Fail
		MenuResp.Msg = "失败"
	}
	return &MenuResp, nil
}
func (s *MenuServer) GetHavingFind(ctx context.Context, in *pb.GetHavingFindRequest) (*pb.GetHavingFindResponse, error) {
	MenuResp := pb.GetHavingFindResponse{}
	MenuResp.Err = errorD.Success
	MenuResp.Msg = "成功"
	err ,Havings := models.FindHaving(int(in.Limit),int(in.Page))
	if err != nil {
		MenuResp.Err = errorD.Fail
		MenuResp.Msg = "失败"
		return &MenuResp,nil
	}
	for _, having := range Havings {
		temp := pb.HavingS{
			HavingId:         int32(having.HavingId),
			HavingPrice:      float32(having.HavingPrice),
			HavingName:       having.HavingName,
			HavingCw:         float32(having.HavingCW),
			HavingOil:        float32(having.HavingOil),
			HavingEnergy:     float32(having.HavingEnergy),
			HavingProtein:    float32(having.HavingProtein),
			HavingUpdateTime: having.HavingUpdateTime,
			HavingContent:    having.HavingContent,
		}
		MenuResp.HavingData = append(MenuResp.HavingData,&temp)
	}
	return &MenuResp,nil
}
func (s *MenuServer) GetMenuDate(ctx context.Context, in *pb.GetMenuDateRequest) (*pb.GetMenuDateResponse, error) {
	MenuResp := pb.GetMenuDateResponse{}
	MenuResp.Err = errorD.Success
	MenuResp.Msg = "成功"
	Havings ,err := models.GetMenuHavingDate(in.CreateDate)
	if err != nil {
		MenuResp.Err = errorD.Fail
		MenuResp.Msg = "失败"
		return &MenuResp, nil
	}
	for _, having := range *Havings {
		temp := pb.HavingS{
			HavingId:         int32(having.HavingId),
			HavingPrice:      float32(having.HavingPrice),
			HavingName:       having.HavingName,
			HavingCw:         float32(having.HavingCW),
			HavingOil:        float32(having.HavingOil),
			HavingEnergy:     float32(having.HavingEnergy),
			HavingProtein:    float32(having.HavingProtein),
			HavingUpdateTime: having.HavingUpdateTime,
			HavingContent:    having.HavingContent,
		}
		MenuResp.HavingData = append(MenuResp.HavingData,&temp)
	}
	return &MenuResp,nil
}
func (s *MenuServer) DeleteMenuToday(ctx context.Context, in *pb.DeleteMenuTodayRequest) (*pb.DeleteMenuTodayResponse, error) {
	MenuResp := pb.DeleteMenuTodayResponse{}
	MenuResp.Err = errorD.Success
	MenuResp.Msg = "成功"
	menu,err := models.GetMenuDate(in.CreateDate)
	if err != nil {
		MenuResp.Err = errorD.Fail
		MenuResp.Msg = "失败"
		return &MenuResp, nil
	}
	MenuId := menu.MenuId
	err = models.DeleteMh(int(MenuId),in.HavingIds)
	if err != nil {
		MenuResp.Err = errorD.Fail
		MenuResp.Msg = "失败"
		return &MenuResp, nil
	}
	return &MenuResp, nil
}
func (s *MenuServer) CreateOftenMenu(ctx context.Context, in *pb.CreateOftenMenuRequest) (*pb.CreateOftenMenuResponse, error) {
	MenuResp := pb.CreateOftenMenuResponse{}
	MenuResp.Err = errorD.Success
	MenuResp.Msg = "创建成功"
	err := models.CreateOftenMenu(in.HavingIds)
	if err != nil {
		MenuResp.Err = errorD.Fail
		MenuResp.Msg = "创建失败"
	}
	return &MenuResp, nil
}
func (s *MenuServer) DeleteOftenMenu(ctx context.Context, in *pb.DeleteOftenMenuRequest) (*pb.DeleteOftenMenuResponse, error) {
	MenuResp := pb.DeleteOftenMenuResponse{}
	MenuResp.Err = errorD.Success
	MenuResp.Msg = "创建成功"
	err := models.DeleteOftenMh(0,in.HavingIds)
	if err != nil {
		MenuResp.Err = errorD.Fail
		MenuResp.Msg = "创建失败"
	}
	return &MenuResp, nil
}