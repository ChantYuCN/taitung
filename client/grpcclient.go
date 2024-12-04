package main

import (
	"context"
	"log"
	"os"
	pb "taitung/api/proto"
	common "taitung/pkg/common"

	"google.golang.org/grpc"
)

func main() {
	log.Print("Client start")
	conn, err := grpc.Dial("127.0.0.1"+":9912", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	client := pb.NewLoadFileModuleInterfaceClient(conn)
	resp, err := client.StoreLogAbsPath(context.Background(), &pb.FilePathBufRequest{Uuid: os.Getenv(common.HLCICDUUID)})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp.Filepath)
}
