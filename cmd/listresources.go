package cmd

import (
	"context"

	kustomizev1 "github.com/fluxcd/kustomize-controller/api/v1beta2"
	responsestruct "github.com/randsw/kubeinfo/KubeApiResponseStruct"
	"github.com/randsw/kubeinfo/logger"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func getAllNamespaces(client kubernetes.Interface) (*v1.NamespaceList, error) {
	namespaces, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaces, nil
}

func ListNodes(client kubernetes.Interface) (*responsestruct.NodeRespose, error) {
	logger.Info("Get cluster nodes...")
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("Error getting cluster nodes...", zap.String("error", err.Error()))
		return nil, err
	}
	nodesInfo := &responsestruct.NodeRespose{}
	nodesInfo.NodeNumber = len(nodes.Items)
	for _, node := range nodes.Items {
		if _, ok := node.Labels["node-role.kubernetes.io/control-plane"]; ok {
			nodesInfo.ContolPlaneNumber++
		} else {
			nodesInfo.WorkerNumber++
		}
		nodesInfo.KubernetesVersion = node.Status.NodeInfo.KubeletVersion
		nodesInfo.OsImage = node.Status.NodeInfo.OSImage
	}
	return nodesInfo, nil
}

func ListNamespaces(client kubernetes.Interface) (*responsestruct.NamespaceRespose, error) {
	logger.Info("Get cluster namespaces...")
	namespaces, err := getAllNamespaces(client)
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	namespaceInfo := &responsestruct.NamespaceRespose{}
	namespaceInfo.NamespaceNumber = len(namespaces.Items)
	return namespaceInfo, nil
}

func ListIngress(client kubernetes.Interface) (*responsestruct.IngressResponse, error) {
	logger.Info("Get cluster ingresses...")
	ingressInfo := &responsestruct.IngressResponse{}
	namespaces, err := getAllNamespaces(client)
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
	return ingressInfo, nil
}

func ListPods(client kubernetes.Interface) (*responsestruct.PodsResponse, error) {
	logger.Info("Get cluster pods...")
	podsInfo := &responsestruct.PodsResponse{}
	namespaces, err := getAllNamespaces(client)
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	for _, namespace := range namespaces.Items {
		pods, err := client.CoreV1().Pods(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Info("Error getting cluster pods...", zap.String("namespace", namespace.Name), zap.String("error", err.Error()))
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
	return podsInfo, nil
}

func ListFluxKustomization(client kubernetes.Interface, clientDynamic dynamic.Interface) (*responsestruct.FluxKustomizationsResponse, error) {
	logger.Info("Get cluster flux kustomizations...")
	fluxKustomizationsInfo := &responsestruct.FluxKustomizationsResponse{}
	namespaces, err := getAllNamespaces(client)
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	var fluxKustomizations = schema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta2", Resource: "kustomizations"}
	for _, namespace := range namespaces.Items {
		fluxKustomizationsList, err := clientDynamic.Resource(fluxKustomizations).Namespace(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Info("Error getting cluster flux kustomization...", zap.String("error", err.Error()))
			return nil, err
		}
		if len(fluxKustomizationsList.Items) > 0 {
			fluxKustomizationsInfo.FluxKustomizationsNumber += len(fluxKustomizationsList.Items)
			for _, item := range fluxKustomizationsList.Items {
				unstructured := item.UnstructuredContent()
				var structured kustomizev1.Kustomization
				err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, &structured)
				if err != nil {
					logger.Info("Error getting cluster flux kustomization...", zap.String("error", err.Error()))
					return nil, err
				}
				switch structured.Status.Conditions[0].Status {
				case "True":
					fluxKustomizationsInfo.FluxKustomizationsNumberbyStatus.Ready++
				case "False":
					fluxKustomizationsInfo.FluxKustomizationsNumberbyStatus.NotReady++
				case "Unknown":
					fluxKustomizationsInfo.FluxKustomizationsNumberbyStatus.Unknown++
				}
			}
		}
	}
	return fluxKustomizationsInfo, nil
}

func ListHelmrelease(client kubernetes.Interface, clientDynamic dynamic.Interface) (*responsestruct.FluxHelmreleasesResponse, error) {
	logger.Info("Get cluster flux helmreleases...")
	fluxHelmreleasesInfo := &responsestruct.FluxHelmreleasesResponse{}
	namespaces, err := getAllNamespaces(client)
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	var fluxHelmreleases = schema.GroupVersionResource{Group: "helm.toolkit.fluxcd.io", Version: "v2beta1", Resource: "helmreleases"}
	for _, namespace := range namespaces.Items {
		fluxHelmreleasesList, err := clientDynamic.Resource(fluxHelmreleases).Namespace(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Info("Error getting cluster flux helmreleases...", zap.String("error", err.Error()))
			return nil, err
		}
		if len(fluxHelmreleasesList.Items) > 0 {
			fluxHelmreleasesInfo.FluxHelmreleasesNumber += len(fluxHelmreleasesList.Items)
			for _, item := range fluxHelmreleasesList.Items {
				unstructured := item.UnstructuredContent()
				var structured kustomizev1.Kustomization
				err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, &structured)
				if err != nil {
					logger.Info("Error getting cluster flux helmreleases...", zap.String("error", err.Error()))
					return nil, err
				}
				switch structured.Status.Conditions[0].Status {
				case "True":
					fluxHelmreleasesInfo.FluxHelmreleasesNumberbyStatus.Ready++
				case "False":
					fluxHelmreleasesInfo.FluxHelmreleasesNumberbyStatus.NotReady++
				case "Unknown":
					fluxHelmreleasesInfo.FluxHelmreleasesNumberbyStatus.Unknown++
				}
			}
		}
	}
	return fluxHelmreleasesInfo, nil
}
