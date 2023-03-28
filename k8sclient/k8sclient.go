package k8sclient

import (
	"os"

	"github.com/randsw/cascadescenariocontroller/logger"
	"go.uber.org/zap"
	"k8s.io/client-go/dynamic"
	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

func ConnectToK8s() *kubernetes.Clientset {
	var kubeconfig string

	config, err := rest.InClusterConfig()
	if err != nil {
		// fallback to kubeconfigkubeconfig := filepath.Join("~", ".kube", "config")
		if envvar := os.Getenv("KUBECONFIG"); len(envvar) > 0 {
			kubeconfig = envvar
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			logger.Zaplog.Error("The kubeconfig cannot be loaded", zap.String("err", err.Error()))
			os.Exit(1)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Zaplog.Error("Failed to create K8s clientset")
		os.Exit(1)
	}

	return clientset
}

func ConnectToK8sDinamic() dynamic.Interface {
	var kubeconfig string

	config, err := rest.InClusterConfig()
	if err != nil {
		// fallback to kubeconfigkubeconfig := filepath.Join("~", ".kube", "config")
		if envvar := os.Getenv("KUBECONFIG"); len(envvar) > 0 {
			kubeconfig = envvar
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			logger.Zaplog.Error("The kubeconfig cannot be loaded", zap.String("err", err.Error()))
			os.Exit(1)
		}
	}

	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		logger.Zaplog.Error("Failed to create K8s clientset")
		os.Exit(1)
	}

	return dynClient
}
