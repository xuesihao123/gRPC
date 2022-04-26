package grpc

import (
	errorD "gRPC/errors"
	models "gRPC/model"
	pb "gRPC/protos"
	"golang.org/x/net/context"
)

//明天的工作  24
//1.编写grpc代码，上线
//2.购买服务器
//3.编写Controller层

// 后2天  25 26
//4.修改small前端代码

type FoodServer struct{}

func (s *FoodServer) CreateFood(ctx context.Context, in *pb.CreateFoodRequest) (*pb.CreateFoodResponse, error) {
	FoodResp := pb.CreateFoodResponse{}
	FoodResp.Err = errorD.Success
	FoodResp.Msg = "创建成功"
	food := make([]models.Food,0)
	for _, data := range in.FoodData {
		temp := models.Food{
			FoodCw:      float64(data.FoodCw),
			FoodOil:     float64(data.FoodOil),
			FoodEnergy:  float64(data.FoodEnergy),
			FoodProtein: float64(data.FoodProtein),
			FoodName:    data.FoodName,
			FoodType:    data.FoodType,
		}
		food = append(food,temp)
	}

	err := models.CreateFood(&food)
	if err != nil {
		FoodResp.Err = errorD.Fail
		FoodResp.Msg = "创建失败"
	}
	return &FoodResp, nil
}

func (s *FoodServer) DeleteFood(ctx context.Context, in *pb.DeleteFoodRequest) (*pb.DeleteFoodResponse, error) {
	FoodResp := pb.DeleteFoodResponse{}
	FoodResp.Err = errorD.Success
	FoodResp.Msg = "创建成功"
	err := models.DeleteFood(int(in.FoodId))
	if err != nil {
		FoodResp.Err = errorD.Fail
		FoodResp.Msg = "创建失败"
	}
	return &FoodResp, nil
}

func (s *FoodServer) UpdateFood(ctx context.Context, in *pb.UpdateFoodRequest) (*pb.UpdateFoodResponse, error) {
	FoodResp := pb.UpdateFoodResponse{}
	FoodResp.Err = errorD.Success
	FoodResp.Msg = "创建成功"
	food := models.Food{
		FoodId:      int(in.Food.FoodId),
		FoodCw:      float64(in.Food.FoodCw),
		FoodOil:     float64(in.Food.FoodOil),
		FoodEnergy:  float64(in.Food.FoodEnergy),
		FoodProtein: float64(in.Food.FoodProtein),
		FoodName:    in.Food.FoodName,
		FoodType:    in.Food.FoodType,
	}
	err := models.UpdateFood(&food)
	if err != nil {
		FoodResp.Err = errorD.Fail
		FoodResp.Msg = "创建失败"
	}
	return &FoodResp, nil
}

func (s *FoodServer) GetFood(ctx context.Context, in *pb.GetFoodRequest) (*pb.GetFoodResponse, error) {
	FoodResp := pb.GetFoodResponse{}
	FoodResp.Err = errorD.Success
	FoodResp.Msg = "创建成功"
	if in.HavingId != 0{
		Foods ,err := models.GetHavingFood(int(in.HavingId))
		if err != nil {
			FoodResp.Err = errorD.Fail
			FoodResp.Msg = "创建失败"
			return &FoodResp, nil
		}
		FoodResp.Food = s.LoadFood(Foods)
		return &FoodResp, nil
	}
	food := models.Food{
		FoodId:      int(in.Food.FoodId),
		FoodCw:      float64(in.Food.FoodCw),
		FoodOil:     float64(in.Food.FoodOil),
		FoodEnergy:  float64(in.Food.FoodEnergy),
		FoodProtein: float64(in.Food.FoodProtein),
		FoodName:    in.Food.FoodName,
		FoodType:    in.Food.FoodType,
	}
	Foods,err := models.GetFood(food,int(in.Limit),int(in.Page))
	if err != nil {
		FoodResp.Err = errorD.Fail
		FoodResp.Msg = "创建失败"
		return &FoodResp, nil
	}
	FoodResp.Food = s.LoadFood(Foods)
	return &FoodResp, nil
}

func (s *FoodServer)LoadFood(foods *[]models.Food)[]*pb.Food  {
	FoodResp := make([]*pb.Food,0)
	for _, food := range *foods {
		temp := pb.Food{
			FoodId:      int32(food.FoodId),
			FoodCw:      float32(food.FoodCw),
			FoodOil:     float32(food.FoodOil),
			FoodEnergy:  float32(food.FoodEnergy),
			FoodProtein: float32(food.FoodProtein),
			FoodName:    food.FoodName,
			FoodType:    food.FoodType,
		}
		FoodResp = append(FoodResp,&temp)
	}
	return FoodResp
}