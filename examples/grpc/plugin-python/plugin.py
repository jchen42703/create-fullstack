from concurrent import futures
import sys
import time

import grpc

import aug_pb2
import aug_pb2_grpc

from grpc_health.v1.health import HealthServicer
from grpc_health.v1 import health_pb2, health_pb2_grpc


class TemplateAugmentor(aug_pb2_grpc.TemplateAugmentorServicer):
    """Implementation of TemplateAugmentorServicer."""

    def Id(self, request, context):
        result = {"id": "ExampleAugmentor2"}
        value = "Written from TemplateAugmentor.Id\n"
        with open("./create-fullstack.log", "w") as f:
            f.write(value + str(result))
        return aug_pb2.IdResponse(**result)

    def Augment(self, request, context):
        value = "\n\nWritten from TemplateAugmentor.Augment"
        with open("./create-fullstack.log", "w") as f:
            f.write(value)
        return aug_pb2.Empty()


def serve():
    # We need to build a health service to work with go-plugin
    health = HealthServicer()
    health.set("plugin",
               health_pb2.HealthCheckResponse.ServingStatus.Value('SERVING'))

    # Start the server.
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    aug_pb2_grpc.add_TemplateAugmentorServicer_to_server(
        TemplateAugmentor(), server)
    health_pb2_grpc.add_HealthServicer_to_server(health, server)
    server.add_insecure_port('127.0.0.1:1234')
    server.start()

    # Output information
    print("1|1|tcp|127.0.0.1:1234|grpc")
    sys.stdout.flush()

    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
