package router

import (
	"context"
	"net/http"
	"time"

	"github.com/bagusandrian/mini-api/src/common"
	"github.com/bagusandrian/mini-api/src/common/monitor"
	"github.com/felixge/httpsnoop"
	"github.com/julienschmidt/httprouter"
	"github.com/opentracing/opentracing-go"
	tlog "github.com/opentracing/opentracing-go/log"
)

type MyRouter struct {
	Httprouter *httprouter.Router
	Options    *Options
}

type Options struct {
	Prefix  string
	Timeout int
}

var HttpRouter *httprouter.Router

func New() *MyRouter {
	HttpRouter = httprouter.New()
	myrouter := &MyRouter{Options: &Options{Prefix: "/mini-api", Timeout: 5}}
	myrouter.Httprouter = HttpRouter
	return myrouter
}

type Handle func(http.ResponseWriter, *http.Request, httprouter.Params) *common.JSONResponse

func (mr *MyRouter) HEAD(path string, handle Handle) {
	mr.Httprouter.HEAD(path, mr.handleNow(path, handle))
}

func (mr *MyRouter) GET(path string, handle Handle) {
	fullPath := mr.Options.Prefix + path
	mr.Httprouter.GET(fullPath, mr.handleNow(fullPath, handle))
}

func (mr *MyRouter) POST(path string, handle Handle) {
	fullPath := mr.Options.Prefix + path
	mr.Httprouter.POST(fullPath, mr.handleNow(fullPath, handle))
}

func (mr *MyRouter) PUT(path string, handle Handle) {
	fullPath := mr.Options.Prefix + path
	mr.Httprouter.PUT(fullPath, mr.handleNow(fullPath, handle))
}

func (mr *MyRouter) DELETE(path string, handle Handle) {
	fullPath := mr.Options.Prefix + path
	mr.Httprouter.DELETE(fullPath, mr.handleNow(fullPath, handle))
}

func (mr *MyRouter) OPTIONS(path string, handle Handle) {
	fullPath := mr.Options.Prefix + path
	mr.Httprouter.OPTIONS(fullPath, mr.handleNow(fullPath, handle))
}

func (mr *MyRouter) ServeFiles(path string, root http.FileSystem) {
	mr.Httprouter.ServeFiles(path, root)
}

func (mr *MyRouter) TestHack(fullPath string, handle Handle) httprouter.Handle {
	return mr.handleNow(fullPath, handle)
}

func (mr *MyRouter) handleNow(fullPath string, handle Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		t := time.Now()
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(mr.Options.Timeout))

		defer cancel()

		ctx = context.WithValue(ctx, "HTTPParams", ps)

		span, ctx := opentracing.StartSpanFromContext(ctx, r.RequestURI)
		defer span.Finish()

		done := make(chan bool)

		r.Header.Set("routePath", fullPath)
		r = r.WithContext(ctx)

		go func() {
			select {
			case <-ctx.Done():
				if ctx.Err() == context.DeadlineExceeded {
					w.WriteHeader(504)
					w.Write([]byte("timeout")) //TODO: shoud be cusotm response
					done <- true
					return
				}
			}
		}()

		go func() {
			resp := handle(w, r, ps)
			if resp != nil {
				span.LogFields(tlog.Object("log", resp.Log))
				span.SetTag("httpCode", resp.StatusCode)
				resp.Header.ProcessTime = time.Since(t).Seconds() * 1000
				resp.SendResponse(w)
			}
			done <- true
		}()
		<-done
		return
	}
}

func GetHttpParam(ctx context.Context, name string) string {
	ps := ctx.Value("HTTPParams").(httprouter.Params)
	return ps.ByName(name)
}

func (mr *MyRouter) WrapperHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(HttpRouter, w, r)
		if r.URL.String() != "/metrics" { // Note: Dont want to monitor for /metrics
			monitor.FeedHTTPMetrics(m.Code, m.Duration, r.Header.Get("routePath"), r.Method)
		}
	})
}
