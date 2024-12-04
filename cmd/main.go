package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	svc "taitung/pkg/server"

	pb "taitung/api/proto"
	grpcs "taitung/pkg/grpcserver"

	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	//	"io"
	//	"bytes"
	// "io/ioutil"
	//      "log"
)

func handleSignal(wg *sync.WaitGroup) {
	defer wg.Done()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	sig := <-sigCh
	log.Printf("received signal: %v, stopping goroutine server", sig)
	// // Signal all goroutines to stop
	// cancel()

	// // Give goroutines time to exit
	// time.Sleep(1 * time.Second)
	os.Exit(9)
}

type wrapStruct struct{}

func (w *wrapStruct) GetHello(ctx echo.Context) error {
	log.Print("Get Hello -- Start")
	log.Print("Chant say hello")
	log.Print("Get Hello -- Stop")
	return ctx.String(http.StatusOK, "Chant say hello\n")
}

func (w *wrapStruct) PostUpload(ctx echo.Context) error {
	log.Print("Post Upload -- Start")

	file, err := ctx.FormFile("file")
	if err != nil {
		log.Print(err)
		return ctx.String(http.StatusBadRequest, "failed Upload file")
	}

	src, err := file.Open()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "failed Upload file")
	}
	defer src.Close()

	x := os.Getenv("HLWORKSPACE")
	//strPwd, _ := os.Getwd()
	dst, err := os.Create(x + file.Filename)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "failed Upload file")
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return ctx.String(http.StatusBadRequest, "failed Upload file")
	}

	log.Print("Post Upload -- stop")
	return ctx.String(http.StatusOK, "Successfully Upload file")

}

func main() {
	log.Print("Start -- rest api server, port 9911")
	e := echo.New()
	wi := new(wrapStruct)
	svc.RegisterHandlers(e, wi)
	go func(e *echo.Echo) {
		if err := e.Start("0.0.0.0:9911"); err != nil {
			log.Fatal("Failed to start echo-struct rest api ")
		}
	}(e)
	log.Print("Stop -- rest api server")

	log.Print("Start -- grpc server, port 9912")

	// fmt.Println(common.SCALINGSERVERPORT)
	lis, err := net.Listen("tcp", ":9912")
	if err != nil {
		log.Fatalf("Failed to listen on port 9912: %v", err)
	}
	grpcSvc := grpc.NewServer()
	pb.RegisterLoadFileModuleInterfaceServer(grpcSvc, &grpcs.ModStr{})

	reflection.Register(grpcSvc)
	go func(grpcSvc *grpc.Server) {
		if err := grpcSvc.Serve(lis); err != nil {
			log.Fatal("Faild to launch a gprc server")
		}
	}(grpcSvc)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go handleSignal(wg)
	wg.Wait()

	log.Print("Stop -- grpc server")

}
