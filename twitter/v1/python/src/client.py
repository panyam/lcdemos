
from __future__ import print_function
import logging

import grpc
import follows_pb2, follows_pb2_grpc
import tweets_pb2, tweets_pb2_grpc
import timelines_pb2, timelines_pb2_grpc

Tweet = tweets_pb2.Tweet
CreateTweetRequest = tweets_pb2.CreateTweetRequest
DeleteTweetRequest = tweets_pb2.DeleteTweetRequest
GetTweetRequest = tweets_pb2.GetTweetRequest
ListTweetsRequest = timelines_pb2.ListTweetsRequest
Timeline = timelines_pb2.Timeline

Follow = follows_pb2.Follow 
GetFollowRequest = follows_pb2.GetFollowRequest 
GetFollowResponse = follows_pb2.GetFollowResponse 
CreateFollowRequest = follows_pb2.CreateFollowRequest 
DeleteFollowRequest = follows_pb2.DeleteFollowRequest 

class Client(object):
    def __init__(self, channel = None):
        channel = channel or grpc.insecure_channel('localhost:50051')
        self.channel = channel
        self.Tweets = tweets_pb2_grpc.TweetServiceStub(channel)
        self.Follows = follows_pb2_grpc.FollowServiceStub(channel)
        self.Timelines = timelines_pb2_grpc.TimelineServiceStub(channel)

        # response = stub.SayHello(helloworld_pb2.HelloRequest(name='you'))

