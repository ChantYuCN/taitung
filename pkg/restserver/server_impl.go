package restserver

import (
	"context"
	"log"
	svc "taitung/pkg/server"
	"time"

	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

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
