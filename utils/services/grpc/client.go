package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sample_go_grpc_testing/utils/config"
)

func GetGrpcClient(grpcService config.GrpcService) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%d", grpcService.Host, grpcService.Port)

	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
