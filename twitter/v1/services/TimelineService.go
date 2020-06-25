package services

import (
	"context"
	// "encoding/json"
	// "fmt"
	// "github.com/golang/protobuf/proto"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
	// "log"
	// "net"
	// "sync"
	// "time"
	"leetcoach.com/demos/twitter/gen"
)

type TimelineService struct {
	// DB things here
}

func (ts *TimelineService) ListTweets(context.Context, *gen.ListTweetsRequest) (*gen.Timeline, error) {
	return nil, nil
}

// Returns most recent N tweets for a user
func (ts *TimelineService) GetTimeline(context.Context, *gen.ListTweetsRequest) (*gen.Timeline, error) {
	return nil, nil
}

func NewTimelineService() *TimelineService {
	t := &TimelineService{}
	return t
}
