
from __future__ import print_function
import logging

import grpc
import follows_pb2, follows_pb2_grpc
import tweets_pb2, tweets_pb2_grpc
import timelines_pb2, timelines_pb2_grpc

class Client(object):
    def __init__(self, channel = None):
        self.channel = channel or grpc.insecure_channel('localhost:50051')
        self.Tweets = tweets_pb2_grpc.TweetServiceStub(channel)
        self.Follows = tweets_pb2_grpc.FollowServiceStub(channel)
        self.Timelines = timelines_pb2_grpc.TimelineServiceStub(channel)

        # response = stub.SayHello(helloworld_pb2.HelloRequest(name='you'))

