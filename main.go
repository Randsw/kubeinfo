package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/randsw/kubeinfo/handlers"
	"github.com/randsw/kubeinfo/logger"
	prom "github.com/randsw/kubeinfo/prometheus-exporter"
	"go.uber.org/zap"
)

func main() {
	//Loger Initialization
	logger.InitLogger()
	defer logger.CloseLogger()
	//Initialize HTTP server
	mux := mux.NewRouter()
	mux.Use(prom.PrometheusMiddleware)
	// Metrics and health handler
	mux.HandleFunc("/healthz", handlers.GetHealth)
	mux.HandleFunc("/metrics", handlers.Metrics)
	// Resource handlers
	mux.HandleFunc("/nodes", handlers.GetNodes)
	mux.HandleFunc("/namespaces", handlers.GetNamespaces)
	mux.HandleFunc("/pods", handlers.GetPods)
	mux.HandleFunc("/ingresses", handlers.GetIngresses)
	mux.HandleFunc("/fluxkustomizations", handlers.GetFluxKustomizations)
	mux.HandleFunc("/fluxhelmreleases", handlers.GetFluxHelmreleases)
	// Main handler
	mux.HandleFunc("/", handlers.GetKubeInfo)
	// Define serving port
	servingPort := "8080"
	if envvar := os.Getenv("API_PORT"); len(envvar) > 0 {
		servingPort = envvar
	}
	// Define serving address
	servingAddress := ""
	if envvar := os.Getenv("API_ADDRESS"); len(envvar) > 0 {
		servingAddress = envvar
	}
	// Construct serving endpoint
	servingAt := servingAddress + ":" + servingPort
	srv := &http.Server{
		Addr: servingAt,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux, // Pass our instance of gorilla/mux in.
	}
	logger.Info("Start serving http request...", zap.String("address", servingAt))
	//Start app
	err := srv.ListenAndServe()
	if err != nil {
		logger.Error("Fail to start http server", zap.String("err", err.Error()))
	}
}
