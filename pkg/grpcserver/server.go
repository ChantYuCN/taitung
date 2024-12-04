package grpcserver

import (
	"context"
	"log"
	pb "taitung/api/proto"
	// "taitung/cmd/common"
	// "taitung/handler"
	//"google.golang.org/grpc/metadata"
)

const (
	RESPONSESUCCESS = "Success"
)

type ModStr struct {
	pb.UnimplementedLoadFileModuleInterfaceServer
}

func (m *ModStr) StoreLogAbsPath(ctx context.Context, req *pb.FilePathBufRequest) (*pb.FilePathBufResponse, error) {
	log.Println("LaunchEdgeNodeInContainers...")
	pbInstance := new(pb.FilePathBufResponse)
	pbInstance.Filepath = "hello path"
	return pbInstance, nil
}
