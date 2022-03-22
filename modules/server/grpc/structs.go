package grpc

import (
	"context"

	"github.com/minish144/rentateam-test/gen/pb"
	"github.com/minish144/rentateam-test/modules/server/grpc/controllers/posts"
)

type ApiServiceServer struct {
	postsCtrl *posts.GRPCControllerModule
	pb.UnimplementedApiServiceServer
}

func (s *ApiServiceServer) CreatePost(ctx context.Context, in *pb.Posts_CreateRequest) (*pb.Posts_CreateResponse, error) {
	return s.postsCtrl.Create(ctx, in)
}

func (s *ApiServiceServer) ListPosts(ctx context.Context, in *pb.Posts_ListRequest) (*pb.Posts_ListResponse, error) {
	return s.postsCtrl.List(ctx, in)
}
