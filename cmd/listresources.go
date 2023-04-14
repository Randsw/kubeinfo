package cmd

import (
	"context"

	kustomizev1 "github.com/fluxcd/kustomize-controller/api/v1beta2"
	responsestruct "github.com/randsw/kubeinfo/KubeApiResponseStruct"
	"github.com/randsw/kubeinfo/logger"
	"go.uber.org/zap"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func GetAllNamespaces(client kubernetes.Interface) ([]string, error) {
	var clusterNs []string
	namespaces, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, namespace := range namespaces.Items {
		clusterNs = append(clusterNs, namespace.Name)
	}
	return clusterNs, nil
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
	namespaces, err := GetAllNamespaces(client)
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	namespaceInfo := &responsestruct.NamespaceRespose{}
	namespaceInfo.NamespaceNumber = len(namespaces)
	return namespaceInfo, nil
}

func ListIngress(client kubernetes.Interface, namespaces []string) (*int, error) {
	var ingressNumber int
	for _, namespace := range namespaces {
		ingresses, err := client.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			logger.Info("Error getting cluster ingress...", zap.String("namespace", namespace), zap.String("error", err.Error()))
			return nil, err
		}
		ingressNumber += len(ingresses.Items)
	}
	return &ingressNumber, nil
}

func CountIngress(client kubernetes.Interface) (*responsestruct.IngressResponse, error) {
	var ingressNumber *int
	logger.Info("Get cluster ingresses...")
	ingressInfo := &responsestruct.IngressResponse{}
	namespaces, err := GetAllNamespaces(client)
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	ingressNumber, err = ListIngress(client, namespaces)
	if err != nil {
		return nil, err
	}
	ingressInfo.IngressNumber = *ingressNumber
	return ingressInfo, nil
}

func ListPods(client kubernetes.Interface, namespace string, podsInfo *responsestruct.PodsResponse) error {
	pods, err := client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("Error getting cluster pods...", zap.String("namespace", namespace), zap.String("error", err.Error()))
		return err
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
	return nil
}

func CountPods(client kubernetes.Interface) (*responsestruct.PodsResponse, error) {
	logger.Info("Get cluster pods...")
	podsInfo := &responsestruct.PodsResponse{}
	namespaces, err := GetAllNamespaces(client)
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	for _, namespace := range namespaces {
		err := ListPods(client, namespace, podsInfo)
		if err != nil {
			logger.Info("Error getting cluster pods...", zap.String("namespace", namespace), zap.String("error", err.Error()))
			return nil, err
		}
	}
	return podsInfo, nil
}

func ListFluxKustomization(clientDynamic dynamic.Interface, namespace string, fluxKustomizationsInfo *responsestruct.FluxKustomizationsResponse) error {
	var fluxKustomizations = schema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta2", Resource: "kustomizations"}
	fluxKustomizationsList, err := clientDynamic.Resource(fluxKustomizations).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("Error getting cluster flux kustomization...", zap.String("error", err.Error()))
		return err
	}
	if len(fluxKustomizationsList.Items) > 0 {
		fluxKustomizationsInfo.FluxKustomizationsNumber += len(fluxKustomizationsList.Items)
		for _, item := range fluxKustomizationsList.Items {
			unstructured := item.UnstructuredContent()
			var structured kustomizev1.Kustomization
			err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, &structured)
			if err != nil {
				logger.Info("Error getting cluster flux kustomization...", zap.String("error", err.Error()))
				return err
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
	return nil
}

func CountFluxKustomization(client kubernetes.Interface, clientDynamic dynamic.Interface) (*responsestruct.FluxKustomizationsResponse, error) {
	logger.Info("Get cluster flux kustomizations...")
	fluxKustomizationsInfo := &responsestruct.FluxKustomizationsResponse{}
	namespaces, err := GetAllNamespaces(client)
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	for _, namespace := range namespaces {
		err := ListFluxKustomization(clientDynamic, namespace, fluxKustomizationsInfo)
		if err != nil {
			logger.Info("Error getting cluster flux kustomization...", zap.String("error", err.Error()))
			return nil, err
		}
	}
	return fluxKustomizationsInfo, nil
}

func ListHelmreleases(clientDynamic dynamic.Interface, namespace string, fluxHelmreleasesInfo *responsestruct.FluxHelmreleasesResponse) error {
	var fluxHelmreleases = schema.GroupVersionResource{Group: "helm.toolkit.fluxcd.io", Version: "v2beta1", Resource: "helmreleases"}
	fluxHelmreleasesList, err := clientDynamic.Resource(fluxHelmreleases).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("Error getting cluster flux helmreleases...", zap.String("error", err.Error()))
		return err
	}
	if len(fluxHelmreleasesList.Items) > 0 {
		fluxHelmreleasesInfo.FluxHelmreleasesNumber += len(fluxHelmreleasesList.Items)
		for _, item := range fluxHelmreleasesList.Items {
			unstructured := item.UnstructuredContent()
			var structured kustomizev1.Kustomization
			err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured, &structured)
			if err != nil {
				logger.Info("Error getting cluster flux helmreleases...", zap.String("error", err.Error()))
				return err
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
	return nil
}

func CountHelmrelease(client kubernetes.Interface, clientDynamic dynamic.Interface) (*responsestruct.FluxHelmreleasesResponse, error) {
	logger.Info("Get cluster flux helmreleases...")
	fluxHelmreleasesInfo := &responsestruct.FluxHelmreleasesResponse{}
	namespaces, err := GetAllNamespaces(client)
	if err != nil {
		logger.Info("Error getting cluster namespaces...", zap.String("error", err.Error()))
		return nil, err
	}
	for _, namespace := range namespaces {
		err := ListHelmreleases(clientDynamic, namespace, fluxHelmreleasesInfo)
		if err != nil {
			logger.Info("Error getting cluster flux helmreleases...", zap.String("error", err.Error()))
			return nil, err
		}
	}
	return fluxHelmreleasesInfo, nil
}
