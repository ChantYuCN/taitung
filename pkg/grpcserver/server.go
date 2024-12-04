package grpcserver

import (
	"context"
	"log"
	pb "taitung/api/proto"
	common "taitung/pkg/common"
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
	log.Print("UUID: ", req.Uuid)
	log.Print("file path: ", common.HashMap[req.Uuid])
	pbInstance.Filepath = common.HashMap[req.Uuid]
	return pbInstance, nil
}
