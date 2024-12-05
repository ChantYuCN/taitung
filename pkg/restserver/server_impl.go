package restserver

import (
	"context"
	"log"
	common "taitung/pkg/common"
	tcnt "taitung/pkg/content"
	svc "taitung/pkg/server"
	"time"

	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/naughtygopher/errors"
)

// var (
// 	//LOGFOLDER = "/mnt/Habana-Content-Management-Service/"
// 	LOGFOLDER = "/home/labuser/habanashared/Habana-Content-Management-Service/"
// )

// future work

type Server struct {
	echo *echo.Echo
}

type RESTServer interface {
	Start()
	Stop()
}

func NewRestServer() (*Server, error) {
	e := echo.New()
	wi := new(wrapStruct)
	svc.RegisterHandlers(e, wi)
	rest := &Server{
		echo: e,
	}
	return rest, nil
}

func (r *Server) Start() {
	if err := r.echo.Start("0.0.0.0:9911"); err != nil {
		log.Fatal("Failed to start echo-struct rest api ")
	}
}

func (r *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// context to stop the server goroutine
	defer cancel()
	if err := r.echo.Shutdown(ctx); err != nil {
		r.echo.Logger.Fatal(err)
	}
}

type wrapStruct struct{}

func (w *wrapStruct) GetFilepath(ctx echo.Context) error {
	log.Print("Get Filepath -- Start")
	jsonReq := new(svc.Oam)
	if err := ctx.Bind(jsonReq); err != nil {
		status, msg, _ := errors.HTTPStatusCodeMessage(err)
		log.Print(" status: %v, msg: %v", status, msg)
		return ctx.JSON(status, msg)
	}

	// if condition to prevent nil pointer
	// log.Print(*jsonReq.Sn)
	// log.Print(*jsonReq.Build)

	snFolderPath := tcnt.SearchTheFile(*jsonReq.Sn, common.LOGFOLDER)
	if snFolderPath == "" {
		var resStr = "Folder Not Found: " + *jsonReq.Sn
		jsonRes := svc.Abspath{Abspaths: &resStr}
		return ctx.JSON(503, jsonRes)
	}
	//if cicduuid := os.Getenv(common.HLCICDUUID); cicduuid != "" {
	if *jsonReq.Uuid != "" {
		log.Print("Put into stack")
		common.HashMap[*jsonReq.Uuid] = snFolderPath
	}

	jsonRes := svc.Abspath{Abspaths: &snFolderPath}

	// store into persistent stack
	// get value from jenkins enviroment
	// unique_Id = UUID.randomUUID().toString()

	log.Print("Get Filepath -- Stop")
	return ctx.JSON(200, jsonRes)

}

func (w *wrapStruct) GetHello(ctx echo.Context) error {
	log.Print("Get Hello -- Start")
	log.Print("Chant say hello")
	log.Print(common.HashMap["ASDF"])
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
