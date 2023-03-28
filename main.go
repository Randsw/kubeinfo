package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	responsestruct "github.com/randsw/kubeinfo/KubeApiResponseStruct"
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

func getKubeInfo(w http.ResponseWriter, r *http.Request) {
	// Variable declaration
	var err error
	nodes := &responsestruct.NodeRespose{}
	namespaces := &responsestruct.NamespaceRespose{}
	pods := &responsestruct.PodsResponse{}
	ingresses := &responsestruct.IngressResponse{}
	clusterInfo := &responsestruct.ResourceResponce{}
	//Connect to k8s api server
	k8sAPIClientset := k8sClient.ConnectToK8s()
	// Get Node information
	nodes, err = ListNodes((k8sAPIClientset))
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Error while requesting kubernetes API about nodes: " + err.Error()
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
		}
		w.Write(jsonResp)
		return
	}
	clusterInfo.Nodes = *nodes
	// Get namespace information
	namespaces, err = ListNamespaces((k8sAPIClientset))
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Error while requesting kubernetes API about namespaces: " + err.Error()
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
		}
		w.Write(jsonResp)
		return
	}
	clusterInfo.Namespaces = *namespaces
	// Get Pod information
	pods, err = ListPods((k8sAPIClientset))
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Error while requesting kubernetes API about pods: " + err.Error()
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
		}
		w.Write(jsonResp)
		return
	}
	clusterInfo.Pods = *pods
	// Get Ingress information
	ingresses, err = ListIngress(k8sAPIClientset)
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Error while requesting kubernetes API about pods: " + err.Error()
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
		}
		w.Write(jsonResp)
		return
	}
	clusterInfo.Ingresses = *ingresses
	// Get Flux kustomization information
	dynamcClient := k8sClient.ConnectToK8sDinamic()
	fluxKustomization, err := ListFluxKustomization(*k8sAPIClientset, dynamcClient)
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Error while requesting kubernetes API about FluxKustomizations: " + err.Error()
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
		}
		w.Write(jsonResp)
		return
	}
	clusterInfo.FluxKustomizations = *fluxKustomization
	// Get Flux Helmrelease information
	fluxHelmrelease, err := ListHelmrelease(*k8sAPIClientset, dynamcClient)
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Error while requesting kubernetes API about FluxHelmreleases: " + err.Error()
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
		}
		w.Write(jsonResp)
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
			logger.Error("Error while marshaling kubeinfo response", zap.String("err", err.Error()))
		}
		w.Write(jsonResp)
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
