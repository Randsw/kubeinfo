package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	responsestruct "github.com/randsw/kubeinfo/KubeApiResponseStruct"
	k8sCRUD "github.com/randsw/kubeinfo/cmd"
	k8sClient "github.com/randsw/kubeinfo/k8sclient"
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

func errorResponse(w http.ResponseWriter, err error, reqresource string) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Error while requesting kubernetes API about " + reqresource + ": " + err.Error()
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
		w.WriteHeader(500)
		_, err := w.Write([]byte("Error while marshaling kubeinfo response"))
		if err != nil {
			logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
		}
		return
	}
	w.WriteHeader(503)
	_, err = w.Write(jsonResp)
	if err != nil {
		logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
	}
}

func getKubeInfo(w http.ResponseWriter, r *http.Request) {
	// Variable declaration
	var err error
	var nodes *responsestruct.NodeRespose
	var namespaces *responsestruct.NamespaceRespose
	var pods *responsestruct.PodsResponse
	var ingresses *responsestruct.IngressResponse
	clusterInfo := &responsestruct.ResourceResponce{}
	//Connect to k8s api server
	k8sAPIClientset := k8sClient.ConnectToK8s()
	// Get Node information
	nodes, err = k8sCRUD.ListNodes(k8sAPIClientset)
	if err != nil {
		errorResponse(w, err, "nodes")
		return
	}
	clusterInfo.Nodes = *nodes
	// Get namespace information
	namespaces, err = k8sCRUD.ListNamespaces(k8sAPIClientset)
	if err != nil {
		errorResponse(w, err, "namespaces")
		return
	}
	clusterInfo.Namespaces = *namespaces
	// Get Pod information
	pods, err = k8sCRUD.CountPods(k8sAPIClientset)
	if err != nil {
		errorResponse(w, err, "pods")
		return
	}
	clusterInfo.Pods = *pods
	// Get Ingress information
	ingresses, err = k8sCRUD.CountIngress(k8sAPIClientset)
	if err != nil {
		errorResponse(w, err, "ingresses")
		return
	}
	clusterInfo.Ingresses = *ingresses
	// Get Flux kustomization information
	dynamcClient := k8sClient.ConnectToK8sDinamic()
	fluxKustomization, err := k8sCRUD.CountFluxKustomization(k8sAPIClientset, dynamcClient)
	if err != nil {
		errorResponse(w, err, "FluxKustomizations")
		return
	}
	clusterInfo.FluxKustomizations = *fluxKustomization
	// Get Flux Helmrelease information
	fluxHelmrelease, err := k8sCRUD.CountHelmrelease(k8sAPIClientset, dynamcClient)
	if err != nil {
		errorResponse(w, err, "FluxHelmreleases")
		return
	}
	clusterInfo.FluxHelmreleases = *fluxHelmrelease

	// Marshal map to JSON and send back
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(clusterInfo)
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Error while marshaling cluster resources data: " + err.Error()
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(500)
			logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
			_, err := w.Write([]byte("Error while marshaling kubeinfo response"))
			if err != nil {
				logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
			}
			return
		}
		w.WriteHeader(503)
		_, err = w.Write(jsonResp)
		if err != nil {
			logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
		}
		return
	}
}

func main() {
	//Loger Initialization
	logger.InitLogger()
	//Initialize HTTP server
	mux := mux.NewRouter()
	mux.Use(prom.PrometheusMiddleware)
	mux.HandleFunc("/healthz", getHealth)
	mux.HandleFunc("/metrics", metrics)
	// Main handler
	mux.HandleFunc("/", getKubeInfo)
	servingAddress := ":8080"
	logger.Info("Start serving http request...", zap.String("address", servingAddress))
	err := http.ListenAndServe(servingAddress, mux)
	if err != nil {
		logger.Error("Fail to start http server", zap.String("err", err.Error()))
	}
}
