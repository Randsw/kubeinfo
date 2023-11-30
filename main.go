package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/randsw/kubeinfo/handlers"
	"github.com/randsw/kubeinfo/logger"
	prom "github.com/randsw/kubeinfo/prometheus-exporter"
	_ "github.com/randsw/kubeinfo/docs" // docs is generated by Swag CLI, you have to import it.
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"
)

// @title Kubeinfo API
// @version 1.0
// @description Kubeinfo used to get information about k8s cluster components
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host kubeinfo-kubeinfo-backend.kubeinfo.svc
// @BasePath /
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
	// Swagger handler
	mux.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
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
