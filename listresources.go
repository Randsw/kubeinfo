package main

import (
	"context"

	responsestruct "github.com/randsw/kubeinfo/KubeApiResponseStruct"
	"github.com/randsw/kubeinfo/logger"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListNodes(client kubernetes.Interface) (*responsestruct.NodeRespose, error) {
	logger.Info("Get cluster nodes...")
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("Error getting cluster nodes...", zap.String("error", err.Error()))
		return nil, err
	}
	nodesInfo := responsestruct.NodeRespose{}
	nodesInfo.NodeNumber = len(nodes.Items)
	for _, node := range nodes.Items {
		if node.Labels["node-role.kubernetes.io/control-plane"] == "true" {
			nodesInfo.ContolPlaneNumber++
		} else if node.Labels["node-role.kubernetes.io/compute"] == "true" {
			nodesInfo.WorkerNumber++
		}
		nodesInfo.KubernetesVersion = node.Status.NodeInfo.KubeletVersion
	}
	return &nodesInfo, nil
}

func ListNamespaces(client kubernetes.Interface) (*responsestruct.NamespaceRespose, error) {
	logger.Info("Get cluster namespaces...")
	namespaces, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	namespaceInfo := responsestruct.NamespaceRespose{}
	namespaceInfo.NamespaceNumber = len(namespaces.Items)
	return &namespaceInfo, nil
}

func ListIngress(client kubernetes.Interface) (*responsestruct.IngressResponse, error) {
	logger.Info("Get cluster ingress...")
	ingressInfo := responsestruct.IngressResponse{}
	namespaces, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	for _, namespace := range namespaces.Items {
		ingresses, err := client.NetworkingV1().Ingresses(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Info("Error getting cluster ingress...", zap.String("namespace", namespace.Name), zap.String("error", err.Error()))
			return nil, err
		}
		ingressInfo.IngressNumber += len(ingresses.Items)
	}
	return &ingressInfo, nil
}

func ListPods(client kubernetes.Interface) (*responsestruct.PodsResponse, error) {
	logger.Info("Get cluster pods...")
	podsInfo := responsestruct.PodsResponse{}
	namespaces, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	for _, namespace := range namespaces.Items {
		pods, err := client.CoreV1().Pods(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Info("Error getting cluster ingress...", zap.String("namespace", namespace.Name), zap.String("error", err.Error()))
			return nil, err
		}
		for _, pod := range pods.Items {
			podsInfo.PodsNumber++
			switch pod.Status.Phase {
			case "Pending":
				podsInfo.PodsNumberByStatus.PendingNumber++
			case "Running":
				podsInfo.PodsNumberByStatus.RunningNumber++
			case "Succeeded":
				podsInfo.PodsNumberByStatus.SucceededNumber++
			case "Failed":
				podsInfo.PodsNumberByStatus.FailedNumber++
			}
		}
	}
	return &podsInfo, nil
}
