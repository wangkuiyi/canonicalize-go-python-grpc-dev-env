import sqlflow_pb2
import sqlflow_pb2_grpc

import grpc


def SQLFlowRun():
    with grpc.insecure_channel('localhost:10000') as channel:
        stub = sqlflow_pb2_grpc.SQLFlowStub(channel)
        job = stub.Run(sqlflow_pb2.SQL(sql="SELECT *"))
        return job.id
