
import grpc
import tweets_pb2, tweets_pb2_grpc

class TweetService(tweets_pb2_grpc.TweetServiceServicer):
    def add_to_server(self, server):
        tweets_pb2_grpc.add_TweetServiceServicer_to_server(self, server)

    def CreateTweet(self, request, context):
        """Creates a new tweet
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteTweet(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTweet(self, request, context):
        """Returns tweet given its ID
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')
