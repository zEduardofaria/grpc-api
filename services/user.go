package services

import "github.com/zEduardofaria/grpc-api"

type UserService struct {
	pb.mustEmbedUnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// Insert - Database
	fmt.Println(req.Name)

	return &pb.User{
		Id: "123",
		Name: req.GetName(),
		Email: req.GetEmail(),
	}, nil
}