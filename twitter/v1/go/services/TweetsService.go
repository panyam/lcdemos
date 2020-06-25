package services

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
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

type TweetService struct {
	// DB things here
}

// Creates a new tweet
func (ts *TweetService) CreateTweet(context.Context, *gen.CreateTweetRequest) (*gen.Tweet, error) {
	return nil, nil
}
func (ts *TweetService) DeleteTweet(context.Context, *gen.DeleteTweetRequest) (*empty.Empty, error) {
	return nil, nil
}

// Returns tweet given its ID
func (ts *TweetService) GetTweet(context.Context, *gen.GetTweetRequest) (*gen.Tweet, error) {
	return nil, nil
}

// BatchGet's tweets given IDs
func (ts *TweetService) BatchGetTweets(context.Context, *gen.BatchGetTweetsRequest) (*gen.BatchGetTweetsResponse, error) {
	return nil, nil
}

func NewTweetService() *TweetService {
	t := &TweetService{}
	return t
}
