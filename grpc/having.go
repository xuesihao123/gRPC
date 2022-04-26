package grpc

import (
	errorD "gRPC/errors"
	models "gRPC/model"
	pb "gRPC/protos"
	"golang.org/x/net/context"
	"time"
)

type HavingServer struct{}

func (s *HavingServer) CreateHaving(ctx context.Context, in *pb.CreateHavingRequest) (*pb.CreateHavingResponse, error) {
	HavingResp := pb.CreateHavingResponse{}
	HavingResp.Err = errorD.Success
	HavingResp.Msg = "创建成功"
	FHs := make([]models.Fh,0)
	for _, data := range in.Foods {
		fh := models.Fh{
			FoodId:     int(data.FoodId),
			FoodWeight: float64(data.FoodWeight),
		}
		FHs = append(FHs,fh)
	}
	having := models.Having{
		HavingPrice:      float64(in.HavingData.HavingPrice),
		HavingName:       in.HavingData.HavingName,
		HavingCW:         float64(in.HavingData.HavingCw),
		HavingOil:        float64(in.HavingData.HavingOil),
		HavingEnergy:     float64(in.HavingData.HavingEnergy),
		HavingProtein:    float64(in.HavingData.HavingProtein),
		HavingUpdateTime: time.Now().Format("2006-01-02 15:04:05"),
		HavingContent:    in.HavingData.HavingContent,
	}

	err := models.CreateHaving(&having,FHs)
	if err != nil {
		HavingResp.Err = errorD.Fail
		HavingResp.Msg = "创建失败"
	}
	return &HavingResp, nil
}
func (s *HavingServer) DeleteHaving(ctx context.Context, in *pb.DeleteHavingRequest) (*pb.DeleteHavingResponse, error) {
	HavingResp := pb.DeleteHavingResponse{}
	HavingResp.Err = errorD.Success
	HavingResp.Msg = "成功"
	err := models.DeleteHaving(int(in.HavingId))
	if err != nil {
		HavingResp.Err = errorD.Fail
		HavingResp.Msg = "失败"
	}
	return &HavingResp, nil
}
func (s *HavingServer) UpdateHaving(ctx context.Context, in *pb.UpdateHavingRequest) (*pb.UpdateHavingResponse, error) {
	HavingResp := pb.UpdateHavingResponse{}
	HavingResp.Err = errorD.Success
	HavingResp.Msg = "成功"
	FHs := make([]models.Fh,0)
	for _, data := range in.Having {
		temp := models.Fh{
			FoodId:     int(data.FoodId),
			HavingId:   int(in.HavingId),
			FoodWeight: float64(data.FoodWeight),
		}
		FHs = append(FHs,temp)
	}
	if in.Status == Insert{
		err := models.CreateFh(int(in.HavingId),FHs)
		if err != nil {
			HavingResp.Err = errorD.Fail
			HavingResp.Msg = "失败"
		}
	}else if in.Status == Delete {
		err := models.DeleteFh(int(in.HavingId),FHs)
		if err != nil {
			HavingResp.Err = errorD.Fail
			HavingResp.Msg = "失败"
		}
	}
	return &HavingResp, nil
}
func (s *HavingServer) GetHaving(ctx context.Context, in *pb.GetHavingRequest) (*pb.GetHavingResponse, error) {
	HavingResp := pb.GetHavingResponse{}
	HavingResp.Err = errorD.Success
	HavingResp.Msg = "成功"
	having := models.Having{
		HavingPrice:      float64(in.Having.HavingPrice),
		HavingName:       in.Having.HavingName,
	}
	havings, err := models.GetHaving(int(in.Limit),int(in.Page),having)
	HavingResp.HavingData = s.LoadHaving(havings)
	if err != nil {
		HavingResp.Err = errorD.Fail
		HavingResp.Msg = "失败"
	}
	return &HavingResp, nil
}

func (s *HavingServer) GetOrderHaving(ctx context.Context, in *pb.GetOrderHavingRequest) (*pb.GetOrderHavingResponse, error) {
	HavingResp := pb.GetOrderHavingResponse{}
	HavingResp.Err = errorD.Success
	HavingResp.Msg = "成功"
	havings, err := models.GetHavingOrder(int(in.OrderId))
	HavingResp.HavingData = s.LoadHaving(havings)
	if err != nil {
		HavingResp.Err = errorD.Fail
		HavingResp.Msg = "失败"
	}
	return &HavingResp, nil
}

func (s *HavingServer)LoadHaving(havings *[]models.Having)[]*pb.Having  {
	HavingResp := make([]*pb.Having,len(*havings))
	for _, having := range *havings {
		temp := pb.Having{
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
		HavingResp = append(HavingResp,&temp)
	}
	return HavingResp
}