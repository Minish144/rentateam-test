package posts

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/minish144/rentateam-test/gen/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c *GRPCControllerModule) Create(ctx context.Context, in *pb.Posts_CreateRequest) (*pb.Posts_CreateResponse, error) {
	// validating request body
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// creating new orm structure to insert data to db
	postID := uuid.New().String()

	newPost := &pb.PostORM{
		ID:     postID,
		Header: in.Header,
		Body:   in.Body,
	}

	var tags []*pb.PostTagORM
	for _, tag := range in.Tags {
		tags = append(tags, &pb.PostTagORM{
			PostID: &postID,
			Tag:    tag,
		})
	}

	// creating new blog post and adding its tags in transaction
	err := c.db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(newPost).Error; err != nil {
			return err
		}
		if err := tx.Create(tags).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"entity": "posts",
			"method": "create",
			"error":  err.Error(),
		}).Errorln("database error occured")
		return nil, status.Error(codes.Internal, err.Error())
	}

	// converting db response to pb structure
	convertedPost, err := newPost.ToPB(ctx)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"entity": "posts",
			"method": "create",
			"error":  err.Error(),
		}).Errorln("could not convert orm structure to pb")
		return nil, status.Error(codes.Internal, err.Error())
	}

	var convertedTags []*pb.PostTag
	for _, tag := range tags {
		if tagConerted, err := tag.ToPB(ctx); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		} else {
			convertedTags = append(convertedTags, &tagConerted)
		}
	}
	convertedPost.Tags = convertedTags

	return &pb.Posts_CreateResponse{Post: &convertedPost}, nil
}

func (c *GRPCControllerModule) List(ctx context.Context, in *pb.Posts_ListRequest) (*pb.Posts_ListResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var postsList []*pb.PostORM
	result := c.db.DB.Preload(clause.Associations).Find(&postsList)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.log.WithFields(logrus.Fields{
			"entity": "posts",
			"method": "create",
			"error":  result.Error.Error(),
		}).Errorln("database error occured")
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	var convertedList []*pb.Post
	for _, ath := range postsList {
		converted, err := ath.ToPB(ctx)
		if err != nil {
			c.log.WithFields(logrus.Fields{
				"entity": "posts",
				"method": "create",
				"error":  err.Error(),
			}).Errorln("could not convert orm structure to pb")
			return nil, status.Error(codes.Internal, err.Error())
		}
		convertedList = append(convertedList, &converted)
	}

	return &pb.Posts_ListResponse{Posts: convertedList, Count: uint32(len(convertedList))}, nil
}
