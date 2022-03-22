package posts

import (
	"context"

	"github.com/minish144/rentateam-test/gen/pb"
)

func (c *GRPCControllerModule) Create(ctx context.Context, in *pb.Posts_CreateRequest) (*pb.Posts_CreateResponse, error) {
	return &pb.Posts_CreateResponse{Post: &pb.Post{Header: "test", Body: "test"}}, nil
}

func (c *GRPCControllerModule) List(ctx context.Context, in *pb.Posts_ListRequest) (*pb.Posts_ListResponse, error) {
	return &pb.Posts_ListResponse{Posts: []*pb.Post{{Header: "test", Body: "test"}}}, nil
}
