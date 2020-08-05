package service

import (
	"flag"
	"fmt"
	endpoint1 "github.com/go-kit/kit/endpoint"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	lightsteptracergo "github.com/lightstep/lightstep-tracer-go"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	prometheus1 "github.com/prometheus/client_golang/prometheus"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	endpoint "github.com/viktormelnychuk/monoapi/monoapi/pkg/endpoint"
	http "github.com/viktormelnychuk/monoapi/monoapi/pkg/http"
	"github.com/viktormelnychuk/monoapi/monoapi/pkg/logger"
	service "github.com/viktormelnychuk/monoapi/monoapi/pkg/service"
	"go.uber.org/zap"
	"net"
	http1 "net/http"
	"os"
	"os/signal"
	appdash "sourcegraph.com/sourcegraph/appdash"
	opentracing "sourcegraph.com/sourcegraph/appdash/opentracing"
	"syscall"
)

var tracer opentracinggo.Tracer

var appLogger *zap.SugaredLogger

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("monoapi", flag.ExitOnError)
var debugAddr = fs.String("debug.addr", ":8080", "Debug and metrics listen address")
var httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")
var lightstepToken = fs.String("lightstep-token", "", "Enable LightStep tracing via a LightStep access token")
var appdashAddr = fs.String("appdash-addr", "", "Enable Appdash tracing via an Appdash server host:port")

func Run() {
	fs.Parse(os.Args[1:])

	// Create a single logger, which we'll use and give to other components.
	appLogger = logger.New()
	//  Determine which tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency
	if *lightstepToken != "" {
		appLogger.Info("tracer", "LightStep")
		tracer = lightsteptracergo.NewTracer(lightsteptracergo.Options{AccessToken: *lightstepToken})
		defer lightsteptracergo.FlushLightStepTracer(tracer)
	} else if *appdashAddr != "" {
		appLogger.Info("tracer", "Appdash", "addr", *appdashAddr)
		collector := appdash.NewRemoteCollector(*appdashAddr)
		tracer = opentracing.NewTracer(collector)
		defer collector.Close()
	} else {
		appLogger.Info("tracer", "none")
		tracer = opentracinggo.GlobalTracer()
	}

	svc := service.New(getServiceMiddleware(appLogger))
	eps := endpoint.New(svc, getEndpointMiddleware(appLogger))
	g := createService(eps)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	appLogger.Info("exit", g.Run())

}
func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultHttpOptions(appLogger, tracer)
	// Add your http options here

	httpHandler := http.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		appLogger.Info("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		appLogger.Info("transport", "HTTP", "addr", *httpAddr)
		return http1.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})

}
func getServiceMiddleware(appLogger *zap.SugaredLogger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = addDefaultServiceMiddleware(appLogger, mw)
	// Append your middleware here

	return
}
func getEndpointMiddleware(appLogger *zap.SugaredLogger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	duration := prometheus.NewSummaryFrom(prometheus1.SummaryOpts{
		Help:      "Request duration in seconds.",
		Name:      "request_duration_seconds",
		Namespace: "example",
		Subsystem: "monoapi",
	}, []string{"method", "success"})
	addDefaultEndpointMiddleware(appLogger, duration, mw)
	// Add you endpoint middleware here

	return
}
func initMetricsEndpoint(g *group.Group) {
	http1.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", *debugAddr)
	if err != nil {
		appLogger.Info("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		appLogger.Info("transport", "debug/HTTP", "addr", *debugAddr)
		return http1.Serve(debugListener, http1.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
}
func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}
