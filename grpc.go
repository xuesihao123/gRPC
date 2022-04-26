package main
//PS C:\Users\DELL\Desktop\Drink> protoc   --proto_path=./protos/ --proto_path=./protos/  --go_out=plugins=grpc:./protos ./protos/having.proto
import (
	"gRPC/dao"
	proto "gRPC/grpc"
	models "gRPC/model"
	pb "gRPC/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	dao.InitMysql()
	err = dao.InitRedis()
	if err != nil {
		panic(err)
	}
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.User{})
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.Order{})
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.Oh{})
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.Fh{})
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.Mh{})
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.Menu{})
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.Having{})
	dao.Db.Set("gorm:table_options","ENGINE = InnoDB").AutoMigrate(models.Food{})

	//defer dao.Db
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//注册服务
	pb.RegisterUserServer(s, &proto.UserServer{})
	pb.RegisterOrderServer(s, &proto.OrderServer{})
	pb.RegisterMenuServer(s, &proto.MenuServer{})
	pb.RegisterFoodServer(s, &proto.FoodServer{})
	pb.RegisterHavingServer(s, &proto.HavingServer{})
	pb.RegisterTokenServer(s, &proto.TokenServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
