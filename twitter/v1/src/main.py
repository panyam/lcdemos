from concurrent import futures
import logging

import grpc

from follows import FollowService
from tweets import TweetService
from timelines import TimelineService

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    follow_service = FollowService()
    tweet_service = TweetService()
    timeline_service = TimelineService()
    follow_service.add_to_server(server)
    tweet_service.add_to_server(server)
    timeline_service.add_to_server(server)
    server.add_insecure_port('[::]:8080')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    serve()

