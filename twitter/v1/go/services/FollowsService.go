package services

import (
	"context"
	// "encoding/json"
	// "fmt"
	// "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	// "log"
	// "net"
	// "sync"
	// "time"
	"leetcoach.com/demos/twitter/gen"
)

type FollowService struct {
	// DB things here
}

// Returns tweet given its ID
func (ts *FollowService) GetFollowers(context.Context, *gen.GetFollowRequest) (*gen.GetFollowResponse, error) {
	return nil, nil
}

func (ts *FollowService) GetFollowees(context.Context, *gen.GetFollowRequest) (*gen.GetFollowResponse, error) {
	return nil, nil
}

func (ts *FollowService) CreateFollow(context.Context, *gen.CreateFollowRequest) (*gen.Follow, error) {
	return nil, nil
}

func (ts *FollowService) DeleteFollow(context.Context, *gen.DeleteFollowRequest) (*empty.Empty, error) {
	return nil, nil
}

func NewFollowService() *FollowService {
	t := &FollowService{}
	return t
}
