package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	svc "taitung/pkg/server"

	"github.com/labstack/echo/v4"
	//	"io"
	//	"bytes"
	// "io/ioutil"
	//      "log"
)

type wrapStruct struct{}

func (w *wrapStruct) GetHello(ctx echo.Context) error {
	log.Print("Post Upload -- Start")
	log.Print("Chant say hello")
	log.Print("Post Upload -- Stop")
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
	e := echo.New()
	wi := new(wrapStruct)
	svc.RegisterHandlers(e, wi)
	err := e.Start("0.0.0.0:9911")
	if err != nil {
		fmt.Println("echo start rest api failed")
	}

}
