package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bagusandrian/mini-api/src/common"
	"github.com/bagusandrian/mini-api/src/common/monitor"
	cRouter "github.com/bagusandrian/mini-api/src/common/router"
	config "github.com/bagusandrian/mini-api/src/config"
	"github.com/bagusandrian/mini-api/src/db"
	"github.com/bagusandrian/mini-api/src/tax"
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
	sMdle := tax.NewModule(conf)
	tax.RegisterRoutes(router, sMdle)

	router.GET("/ping", ping)
	tracer.Init(&tracer.Config{Port: conf.Server.TracerPort, Enabled: true})
	log.Fatal(grace.Serve(":"+conf.Server.Port, router.WrapperHandler()))
}

func ping(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.JSONResponse) {
	// this is just for check that service is running
	resp = &common.JSONResponse{
		Data:       "Welcome to Mini-API",
		StatusCode: http.StatusOK,
	}
	return
}
