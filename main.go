package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	k8sClient "github.com/randsw/cascadescenariocontroller/k8sclient"
	"github.com/randsw/kubeinfo/logger"
	prom "github.com/randsw/kubeinfo/prometheus-exporter"
	"go.uber.org/zap"
)

func getHealth(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Ok. KubeInfo. Version 1.0.0"))
}

func metrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

func getKubeInfo(w http.ResponseWriter, r *http.Request) {
	//Connect to k8s api server
	k8sAPIClientset := k8sClient.ConnectToK8s()
	// #TODO Get Node information
	// #TODO Get namespace information
	// #TODO Get Pod informatio
	// #TODO Get Ingress information
	// #TODO Get Flux kustomization information
	// #TODO Get Flux Helmrelease information

	// # TODO Marshal map to JSON and send back
}

func main() {
	//Loger Initialization
	logger.InitLogger()
	//Initialize HTTP server
	mux := mux.NewRouter()
	mux.Use(prom.PrometheusMiddleware)
	mux.HandleFunc("/healthz", getHealth)
	mux.HandleFunc("/metrics", metrics)
	//#TODO main handler
	servingAddress := ":8080"
	logger.Info("Start serving http request...", zap.String("address", servingAddress))
	err := http.ListenAndServe(servingAddress, mux)
	if err != nil {
		logger.Error("Fail to start http server", zap.String("err", err.Error()))
	}
}
