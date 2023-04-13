package main

import (
	"net/http"
	"os"

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
	servingPort := ":8080"
	if envvar := os.Getenv("API_PORT"); len(envvar) > 0 {
		servingPort = envvar
	}
	servingAddress := ""
	if envvar := os.Getenv("API_ADDRESS"); len(envvar) > 0 {
		servingAddress = envvar
	}
	servingAt := servingAddress + ":" + servingPort
	logger.Info("Start serving http request...", zap.String("address", servingAt))
	err := http.ListenAndServe(servingAt, mux)
	if err != nil {
		logger.Error("Fail to start http server", zap.String("err", err.Error()))
	}
}
