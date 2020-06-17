
from gen import timelines_pb2, timelines_pb2_grpc

class TimelineService(timelines_pb2_grpc.TimelineServiceServicer):
    def add_to_server(self, server):
        timelines_pb2_grpc.add_TimelineServiceServicer_to_server(self, server)

    def ListTweets(self, request, context):
        """Returns a user's home timeline with recent N tweets from 
        all users the user follows
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTimeline(self, request, context):
        """Returns most recent N tweets for a user
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')
