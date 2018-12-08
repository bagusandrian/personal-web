package monitor

import (
	"fmt"
	"github.com/bagusandrian/mini-api/src/config"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"time"
)

var (
	httpLatencyHistogram      *prometheus.HistogramVec
	httpResponsesTotalCounter *prometheus.CounterVec
	cfg                       *config.Config
)

func Init(conf *config.Config) {
	cfg = conf
	registerHistogram()
	registerCounter()
}

func registerHistogram() {
	httpLatencyHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "mini_api_http_duration_seconds",
		Help: "the latency of http calls",
	}, []string{"handler", "method", "httpcode", "env"})

	err := prometheus.Register(httpLatencyHistogram)
	if err != nil {
		log.Printf("[Monitor] Unable to Register httpLatencyHistogram. Err: %+v\n", err)
	}
}

func registerCounter() {
	httpResponsesTotalCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mini_api_http_responses_total",
			Help: "The count of http responses issued",
		},
		[]string{"handler", "method", "httpcode", "env"},
	)
	err := prometheus.Register(httpResponsesTotalCounter)
	if err != nil {
		log.Printf("[Monitor] Unable to Register httpResponsesTotalCounter. Err: %+v\n", err)
	}
}

// FeedHTTPMetrics to monitor latency, http code counts
func FeedHTTPMetrics(status int, duration time.Duration, path string, method string) {
	httpLatencyHistogram.With(prometheus.Labels{"handler": "all", "method": method, "httpcode": fmt.Sprintf("%d", status), "env": cfg.Environment}).Observe(duration.Seconds() * 1000)
	httpLatencyHistogram.With(prometheus.Labels{"handler": path, "method": method, "httpcode": fmt.Sprintf("%d", status), "env": cfg.Environment}).Observe(duration.Seconds() * 1000)

	httpResponsesTotalCounter.With(prometheus.Labels{"handler": "all", "method": method, "httpcode": fmt.Sprintf("%d", status), "env": cfg.Environment}).Inc()

}
