package grpc

import (
	"fmt"
	"gRPC/dao"
	pb "gRPC/protos"
	"gRPC/service"
	"golang.org/x/net/context"
	"strconv"
)

type TokenServer struct{}

func (s *TokenServer) ParseToken(ctx context.Context, in *pb.ParseTokenRequest) (*pb.ParseTokenResponse, error) {
	data ,_ := service.ParseToken(in.Token)
	S,_ := dao.GetKey(ctx,strconv.Itoa(data.ID))
	status := "false"
	if S == "Index"{
		status = "Index"
	}
	fmt.Println(S)
	return &pb.ParseTokenResponse{
		UserName:   data.Username,
		UserStatus: int32(data.UserStatus),
		UserEmail:  data.UserEmail,
		UserId:     int32(data.ID),
		UserPid:    data.UserPid,
		Err:        200,
		Msg:        "成功",
		Flag:       status,
	}, nil
}