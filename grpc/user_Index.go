package grpc
//
//import (
//	"Drink/dao"
//	models "Drink/model"
//	"log"
//	"net"
//
//	pb "Drink/protos"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/reflection"
//)
//
//// server is used to implement helloworld.GreeterServer.
//type server struct{}
//
//// SayHello implements helloworld.GreeterServer
//func (s *server) UserIndex(ctx context.Context, in *pb.UserIndexRequest) (*pb.UserIndexResponse, error) {
//	return &pb.UserIndexResponse{Msg: "xuesihao"}, nil
//}
//
//// SayHello implements helloworld.GreeterServer
//func (s *server) UserView(ctx context.Context, in *pb.UserViewRequest) (*pb.UserViewResponse,error) {
//	return &pb.UserViewResponse{Msg: "xuesihao"}, nil
//}
//
//// SayHello implements helloworld.GreeterServer
//func (s *server) UserDelete(ctx context.Context, in *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
//	return &pb.UserDeleteResponse{Msg: "xuesihao"}, nil
//}
//
//// SayHello implements helloworld.GreeterServer
//func (s *server) UserRegister(ctx context.Context, in *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
//	User := &models.User{
//		UserName:     in.UserName,
//		UserPassword: in.UserPassword,
//		UserEmail:    in.UserEmail,
//		UserStatus:   1,
//	}
//	err := models.CreateU(User)
//	if err != nil {
//		return nil,err
//	}
//	return &pb.UserRegisterResponse{Msg: "123"}, nil
//}
//
//// SayHello implements helloworld.GreeterServer
//func (s *server) UserUpdate(ctx context.Context, in *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
//	return &pb.UserUpdateResponse{Msg: "xuesihao"}, nil
//}
//
//func main() {
//	lis, err := net.Listen("tcp", ":50051")
//	dao.InitMysql()
//	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.User{})
//
//	defer dao.Db.Close()
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//	pb.RegisterUserServer(s, &server{})
//	// Register reflection service on gRPC server.
//	reflection.Register(s)
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
//
