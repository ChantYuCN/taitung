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

type fr struct {
	lx map[string]interface{}
}

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

	//strPwd, _ := os.Getwd()
	dst, err := os.Create(file.Filename + "_done")
	if err != nil {
		return ctx.String(http.StatusBadRequest, "failed Upload file")
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return ctx.String(http.StatusBadRequest, "failed Upload file")
	}
	// //oa := new(svc.Oneuy)

	// //if err := ctx.Bind(oa); err != nil {
	// //	return ctx.JSON(http.StatusBadRequest, err.Error())
	// //}
	// //fmt.Println(oa.TemplateFiles)
	// file, err := ctx.FormFile("data")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// /*
	// 		bytebuf := bytes.NewBuffer(nil)
	// 		if _, err := io.Copy(bytebuf, file); err != nil {
	// 	                fmt.Println(err)
	// 	                return err
	// 		}
	// */
	// fmt.Println(0)
	// src, err := file.Open()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// defer src.Close()

	// //data := wrapStruct
	// /*
	// 	buf := make([]byte, 100)
	// 	n, err := src.Read(buf)
	// 	fmt.Println(n, err, buf[:n])
	// */
	// fmt.Println(1)
	// bodyresp, err := ioutil.ReadAll(src)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err

	// }
	// var fi fr
	// fmt.Println(2)
	// err = yaml.Unmarshal(bodyresp, &fi.lx)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// fmt.Println(fi.lx)

	// lp := new(interface{})
	// if err := ctx.Bind(lp); err != nil {
	// 	return err
	// }
	// fmt.Println(lp)
	log.Print("Post Upload -- stop")
	return ctx.String(http.StatusOK, "Successfully Upload file")

}

type yType struct {
	Name    string
	Version string
	Chart   string
}

/*
   fmt.Println("Start to make yaml")
   yf, err := ioutil.ReadFile("cool.ext")
   if err != nil {
           log.Fatal(err)
   }
*/

func main() {
	e := echo.New()
	wi := new(wrapStruct)
	svc.RegisterHandlers(e, wi)
	err := e.Start("0.0.0.0:9911")
	if err != nil {
		fmt.Println("echo start rest api failed")
	}

}
