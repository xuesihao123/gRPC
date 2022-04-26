package grpc

import (
	errorD "gRPC/errors"
	models "gRPC/model"
	pb "gRPC/protos"
	"golang.org/x/net/context"
)

// server is used to implement helloworld.GreeterServer.
type UserServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *UserServer) UserIndex(ctx context.Context, in *pb.UserIndexRequest) (*pb.UserIndexResponse, error) {
	if in.GetUserName() != ""{
		UserReq := &models.User{
			UserName:     in.UserName,
			UserPassword: in.UserPassword,
		}
		err , _ := models.GetNameAndPassword(UserReq)
		if err != nil {
			Resp := &pb.UserIndexResponse{
				Err: errorD.Fail,
				Msg: "失败",
			}
			return Resp , err
		}
		Resp := &pb.UserIndexResponse{
			Err: errorD.Success,
			Msg: "成功",
		}
		return Resp,nil
	}

	if in.GetUserEmail() != ""{
		UserReq := &models.User{
			UserEmail:    in.UserEmail,
			UserPassword: in.UserPassword,
		}
		err , _ := models.GetEmailAndPassword(UserReq)
		if err != nil {
			Resp := &pb.UserIndexResponse{
				Err: errorD.Fail,
				Msg: "失败",
			}
			return Resp , err
		}
		Resp := &pb.UserIndexResponse{
			Err: errorD.Success,
			Msg: "成功",
		}
		return Resp,nil
	}

	return &pb.UserIndexResponse{Msg: "内部错误",Err: errorD.InnError}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *UserServer) UserView(ctx context.Context, in *pb.UserViewRequest) (*pb.UserViewResponse, error) {
	return &pb.UserViewResponse{Msg: "xuesihao"}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *UserServer) UserDelete(ctx context.Context, in *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
	if in.GetUserId() == 0{
		return nil,nil
	}
	err := models.DeleteU(int(in.UserId))
	if err != nil {
		return nil,err
	}
	UserResp := pb.UserDeleteResponse{}
	UserResp.Err = errorD.Success
	UserResp.Msg = "成功"
	return &UserResp, nil
}

// SayHello implements helloworld.GreeterServer
func (s *UserServer) UserRegister(ctx context.Context, in *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	User := &models.User{
		UserName:     in.UserName,
		UserPassword: in.UserPassword,
		UserEmail:    in.UserEmail,
		UserStatus:   1,
	}
	err := models.CreateU(User)
	if err != nil {
		return nil, err
	}
	UserResp := pb.UserRegisterResponse{}
	UserResp.Err = errorD.Success
	UserResp.Msg = "成功"
	return &UserResp, nil
}

// SayHello implements helloworld.GreeterServer
func (s *UserServer) UserUpdate(ctx context.Context, in *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	User := &models.User{
		UserId:     int(in.UserId),
		UserPassword: in.UserPassword,
	}
	if err := models.UpdateU(User); err != nil {
		return nil,nil
	}
	UserResp := pb.UserUpdateResponse{}
	UserResp.Err = errorD.Success
	UserResp.Msg = "成功"
	return &UserResp, nil
}

// SayHello implements helloworld.GreeterServer
func (s *UserServer) UserGet(ctx context.Context, in *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	UserReq := &models.User{
		UserName:     in.UserName,
		UserEmail:    in.UserEmail,
		UserStatus:   int(in.UserStatus),
	}
	UserResp := pb.UserGetResponse{}
	if users,err := models.GetUser(UserReq,int(in.Limit),int(in.Page));err != nil{
		UserResp.Err = errorD.Success
		UserResp.Msg = "成功"
		for _, user := range *users {
			temp := pb.UserEntity{
				UserName:  	user.UserName,
				UserStatus: int32(user.UserStatus),
				UserEmail:  user.UserEmail,
				UserId:     int32(user.UserId),
				UserPid:    user.UserPid,
			}
			UserResp.Data = append(UserResp.Data,&temp)
		}
	}
	return &UserResp, nil
}
