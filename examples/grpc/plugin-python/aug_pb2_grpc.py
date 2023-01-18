# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import aug_pb2 as aug__pb2


class TemplateAugmentorStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Id = channel.unary_unary(
                '/proto.TemplateAugmentor/Id',
                request_serializer=aug__pb2.Empty.SerializeToString,
                response_deserializer=aug__pb2.IdResponse.FromString,
                )
        self.Augment = channel.unary_unary(
                '/proto.TemplateAugmentor/Augment',
                request_serializer=aug__pb2.Empty.SerializeToString,
                response_deserializer=aug__pb2.Empty.FromString,
                )


class TemplateAugmentorServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Id(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Augment(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TemplateAugmentorServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Id': grpc.unary_unary_rpc_method_handler(
                    servicer.Id,
                    request_deserializer=aug__pb2.Empty.FromString,
                    response_serializer=aug__pb2.IdResponse.SerializeToString,
            ),
            'Augment': grpc.unary_unary_rpc_method_handler(
                    servicer.Augment,
                    request_deserializer=aug__pb2.Empty.FromString,
                    response_serializer=aug__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.TemplateAugmentor', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TemplateAugmentor(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Id(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.TemplateAugmentor/Id',
            aug__pb2.Empty.SerializeToString,
            aug__pb2.IdResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Augment(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.TemplateAugmentor/Augment',
            aug__pb2.Empty.SerializeToString,
            aug__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
