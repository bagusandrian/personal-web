package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bagusandrian/web-personal/src/common"
	"github.com/bagusandrian/web-personal/src/common/monitor"
	cRouter "github.com/bagusandrian/web-personal/src/common/router"
	config "github.com/bagusandrian/web-personal/src/config"
	"github.com/bagusandrian/web-personal/src/db"
	// "github.com/bagusandrian/web-personal/src/tax"
	"github.com/google/gops/agent"
	"github.com/julienschmidt/httprouter"
	grace "gopkg.in/tokopedia/grace.v1"
	logging "gopkg.in/tokopedia/logging.v1"
	"gopkg.in/tokopedia/logging.v1/tracer"
)

var conf *config.Config
var err error

func init() {
	conf = config.ReadConfig()
	log.Printf("%+v\n", conf)
	db.Init(conf)
	monitor.Init(conf)
}

func main() {
	flag.Parse()
	logging.LogInit()
	log.SetFlags(log.LstdFlags | log.Llongfile)

	debug := logging.Debug.Println

	debug("app started") // message will not appear unless run with -debug switch

	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	go logging.StatsLog()

	router := cRouter.New()
	// sMdle := tax.NewModule(conf)
	// tax.RegisterRoutes(router, sMdle)

	router.GET("/ping", ping)
	router.GETHTML("/pinghtml", pingHTML)
	router.GETFILE("/css/:filename", getCSS)
	router.GETFILE("/js/:filename", getJS)
	router.GETFILE("/fonts/:filename", getFont)
	router.GETFILE("/img/:filename", getImage)
	// router.GETFILE("/img/:path/:filename", getImageA)
	tracer.Init(&tracer.Config{Port: conf.Server.TracerPort, Enabled: true})
	log.Fatal(grace.Serve(":"+conf.Server.Port, router.WrapperHandler()))
}

func ping(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.JSONResponse) {
	// this is just for check that service is running
	resp = &common.JSONResponse{
		Data:       "Welcome to web-personal",
		StatusCode: http.StatusOK,
	}
	return
}

func pingHTML(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.HTMLResponse) {
	resp = &common.HTMLResponse{
		Data:       "Testing Wae",
		StatusCode: http.StatusOK,
	}
	return
}

func getCSS(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.FileResponse) {
	// this is just for check that service is running
	file := params.ByName("filename")
	resp = &common.FileResponse{
		Path: "css/",
		File: file,
		Type: "text/css",
	}
	return
}
func getJS(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.FileResponse) {
	// this is just for check that service is running
	file := params.ByName("filename")
	resp = &common.FileResponse{
		Path: "js/",
		File: file,
		Type: "text/javascript",
	}
	return
}
func getFont(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.FileResponse) {
	// this is just for check that service is running
	file := params.ByName("filename")
	resp = &common.FileResponse{
		Path: "fonts/",
		File: file,
		Type: "text/plain",
	}
	return
}
func getImage(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.FileResponse) {
	// this is just for check that service is running
	file := params.ByName("filename")
	resp = &common.FileResponse{
		Path: "img/",
		File: file,
		Type: "image/jpg",
	}
	return
}
