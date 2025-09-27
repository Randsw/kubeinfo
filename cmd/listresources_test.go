package cmd

import (
	"reflect"
	"strings"
	"testing"

	helmv2 "github.com/fluxcd/helm-controller/api/v2"
	kustomizev1 "github.com/fluxcd/kustomize-controller/api/v1beta2"
	responsestruct "github.com/randsw/kubeinfo/KubeApiResponseStruct"
	"github.com/randsw/kubeinfo/logger"
	v1 "k8s.io/api/core/v1"
	network "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	dynamicfake "k8s.io/client-go/dynamic/fake"

	//"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/client-go/kubernetes/fake"
)

func TestListNodes(t *testing.T) {
	testCases := []struct {
		name           string
		clusterNS      []runtime.Object
		expectedNumber int
		expectedNS     []string
	}{
		{
			clusterNS: []runtime.Object{
				&v1.NamespaceList{
					TypeMeta: metav1.TypeMeta{},
					ListMeta: metav1.ListMeta{},
					Items: []v1.Namespace{
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name: "default",
							},
							Spec:   v1.NamespaceSpec{},
							Status: v1.NamespaceStatus{},
						},
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name: "kube-system",
							},
							Spec:   v1.NamespaceSpec{},
							Status: v1.NamespaceStatus{},
						},
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name: "monitoring",
							},
							Spec:   v1.NamespaceSpec{},
							Status: v1.NamespaceStatus{},
						},
					},
				},
			},
			name:           "standart",
			expectedNumber: 3,
			expectedNS:     []string{"default", "kube-system, monitoring"},
		},
	}
	logger.InitLogger()
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset(test.clusterNS...)
			nsnumber, err := ListNamespaces(fakeClientSet)
			if nsnumber.NamespaceNumber != test.expectedNumber {
				t.Fatalf("Number of namespaces %d are not equal expected value %d", nsnumber.NamespaceNumber, test.expectedNumber)
			}
			if err != nil {
				t.Fatalf("unexpected error while counting namespaces: %v", err)
			}
		})
	}
}

func TestGetAllNamespaces(t *testing.T) {
	testCases := []struct {
		name           string
		clusterNS      []runtime.Object
		expectedNumber int
		expectedNS     []string
	}{
		{
			clusterNS: []runtime.Object{
				&v1.NamespaceList{
					TypeMeta: metav1.TypeMeta{},
					ListMeta: metav1.ListMeta{},
					Items: []v1.Namespace{
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name: "default",
							},
							Spec:   v1.NamespaceSpec{},
							Status: v1.NamespaceStatus{},
						},
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name: "kube-system",
							},
							Spec:   v1.NamespaceSpec{},
							Status: v1.NamespaceStatus{},
						},
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name: "monitoring",
							},
							Spec:   v1.NamespaceSpec{},
							Status: v1.NamespaceStatus{},
						},
					},
				},
			},
			name:           "standart",
			expectedNumber: 3,
			expectedNS:     []string{"default", "kube-system", "monitoring"},
		},
	}
	logger.InitLogger()
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset(test.clusterNS...)
			namespaces, err := GetAllNamespaces(fakeClientSet)
			// for _, namespace := range namespaces.Items {
			// 	NsNames = append(NsNames, namespace.Name)
			// }
			if !reflect.DeepEqual(test.expectedNS, namespaces) {
				t.Fatalf("Namespace names %s not as expected %s", strings.Join(namespaces, ","), strings.Join(test.expectedNS, ","))
			}
			if len(namespaces) != test.expectedNumber {
				t.Fatalf("Number of namespaces %d are not equal expected value %d", len(namespaces), test.expectedNumber)
			}
			if err != nil {
				t.Fatalf("unexpected error while counting namespaces: %v", err)
			}
		})
	}
}

func TestListIngress(t *testing.T) {
	testCases := []struct {
		name           string
		clusterIng     []runtime.Object
		expectedNumber int
		expectedNS     []string
	}{
		{
			clusterIng: []runtime.Object{
				&network.IngressList{
					TypeMeta: metav1.TypeMeta{},
					ListMeta: metav1.ListMeta{},
					Items: []network.Ingress{
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-ing",
								Namespace: "default",
							},
							Spec:   network.IngressSpec{},
							Status: network.IngressStatus{},
						},
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "kube-system-ing",
								Namespace: "kube-system",
							},
							Spec:   network.IngressSpec{},
							Status: network.IngressStatus{},
						},
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "monitoring-ing",
								Namespace: "monitoring",
							},
							Spec:   network.IngressSpec{},
							Status: network.IngressStatus{},
						},
					},
				},
			},
			name:           "standart",
			expectedNumber: 3,
			expectedNS:     []string{"default", "kube-system", "monitoring"},
		},
	}
	logger.InitLogger()
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			var ingressNumber *int
			fakeClientSet := fake.NewSimpleClientset(test.clusterIng...)
			ingressNumber, err := ListIngress(fakeClientSet, test.expectedNS)

			if *ingressNumber != test.expectedNumber {
				t.Fatalf("Number of ingresses %d are not equal expected value %d", *ingressNumber, test.expectedNumber)
			}
			if err != nil {
				t.Fatalf("unexpected error while counting namespaces: %v", err)
			}
		})
	}
}

func TestListPods(t *testing.T) {
	testCases := []struct {
		name            string
		clusterPods     []runtime.Object
		expectedNumber  int
		expectedSuccess int
		expectedRunning int
		expectedFailed  int
		expectedPending int
		namespace       string
	}{
		{
			clusterPods: []runtime.Object{
				&v1.PodList{
					TypeMeta: metav1.TypeMeta{},
					ListMeta: metav1.ListMeta{},
					Items: []v1.Pod{
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-1",
								Namespace: "default",
							},
							Spec: v1.PodSpec{},
							Status: v1.PodStatus{
								Phase: "Running",
							},
						},
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-2",
								Namespace: "default",
							},
							Spec: v1.PodSpec{},
							Status: v1.PodStatus{
								Phase: "Running",
							},
						},
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-3",
								Namespace: "default",
							},
							Spec: v1.PodSpec{},
							Status: v1.PodStatus{
								Phase: "Pending",
							},
						},
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-4",
								Namespace: "default",
							},
							Spec: v1.PodSpec{},
							Status: v1.PodStatus{
								Phase: "Succeeded",
							},
						},
						{
							TypeMeta: metav1.TypeMeta{},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-5",
								Namespace: "default",
							},
							Spec: v1.PodSpec{},
							Status: v1.PodStatus{
								Phase: "Failed",
							},
						},
					},
				},
			},
			name:            "standart",
			expectedNumber:  5,
			expectedSuccess: 1,
			expectedRunning: 2,
			expectedFailed:  1,
			expectedPending: 1,
			namespace:       "default",
		},
	}
	logger.InitLogger()
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			podsInfo := &responsestruct.PodsResponse{}
			fakeClientSet := fake.NewSimpleClientset(test.clusterPods...)
			err := ListPods(fakeClientSet, test.namespace, podsInfo)
			if podsInfo.PodsNumber != test.expectedNumber {
				t.Fatalf("Number of pods in namespace %s - %d are not equal expected value %d", test.namespace, podsInfo.PodsNumber, test.expectedNumber)
			}
			if podsInfo.PodsNumberByStatus.RunningNumber != test.expectedRunning {
				t.Fatalf("Number of running pods in namespace %s - %d are not equal expected value %d", test.namespace, podsInfo.PodsNumberByStatus.RunningNumber, test.expectedRunning)
			}
			if podsInfo.PodsNumberByStatus.SucceededNumber != test.expectedSuccess {
				t.Fatalf("Number of succeeded pods in namespace %s - %d are not equal expected value %d", test.namespace, podsInfo.PodsNumberByStatus.SucceededNumber, test.expectedSuccess)
			}
			if podsInfo.PodsNumberByStatus.FailedNumber != test.expectedFailed {
				t.Fatalf("Number of failed pods in namespace %s - %d are not equal expected value %d", test.namespace, podsInfo.PodsNumberByStatus.FailedNumber, test.expectedFailed)
			}
			if podsInfo.PodsNumberByStatus.PendingNumber != test.expectedPending {
				t.Fatalf("Number of pending pods in namespace %s - %d are not equal expected value %d", test.namespace, podsInfo.PodsNumberByStatus.PendingNumber, test.expectedPending)
			}
			if err != nil {
				t.Fatalf("unexpected error while counting namespaces: %v", err)
			}
		})
	}
}

func TestListFluxKustomization(t *testing.T) {
	testCases := []struct {
		name                      string
		clusterFluxKustomizations []runtime.Object
		expectedNumber            int
		expectedReady             int
		expectedNotReady          int
		expectedUnknown           int
		namespace                 string
	}{
		{
			clusterFluxKustomizations: []runtime.Object{
				&kustomizev1.KustomizationList{
					TypeMeta: metav1.TypeMeta{
						APIVersion: "kustomize.toolkit.fluxcd.io/v1beta2",
						Kind:       "KustomizationList",
					},
					ListMeta: metav1.ListMeta{},
					Items: []kustomizev1.Kustomization{
						{
							TypeMeta: metav1.TypeMeta{
								APIVersion: "kustomize.toolkit.fluxcd.io/v1beta2",
								Kind:       "Kustomization",
							},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-1",
								Namespace: "default",
							},
							Spec: kustomizev1.KustomizationSpec{},
							Status: kustomizev1.KustomizationStatus{
								Conditions: []metav1.Condition{
									{
										Status: "True",
									},
									{},
								},
							},
						},
						{
							TypeMeta: metav1.TypeMeta{
								APIVersion: "kustomize.toolkit.fluxcd.io/v1beta2",
								Kind:       "Kustomization",
							},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-2",
								Namespace: "default",
							},
							Spec: kustomizev1.KustomizationSpec{},
							Status: kustomizev1.KustomizationStatus{
								Conditions: []metav1.Condition{
									{
										Status: "True",
									},
									{},
								},
							},
						},
						{
							TypeMeta: metav1.TypeMeta{
								APIVersion: "kustomize.toolkit.fluxcd.io/v1beta2",
								Kind:       "Kustomization",
							},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-3",
								Namespace: "default",
							},
							Spec: kustomizev1.KustomizationSpec{},
							Status: kustomizev1.KustomizationStatus{
								Conditions: []metav1.Condition{
									{
										Status: "False",
									},
									{},
								},
							},
						},
						{
							TypeMeta: metav1.TypeMeta{
								APIVersion: "kustomize.toolkit.fluxcd.io/v1beta2",
								Kind:       "Kustomization",
							},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-4",
								Namespace: "default",
							},
							Spec: kustomizev1.KustomizationSpec{},
							Status: kustomizev1.KustomizationStatus{
								Conditions: []metav1.Condition{
									{
										Status: "True",
									},
									{},
								},
							},
						},
						{
							TypeMeta: metav1.TypeMeta{
								APIVersion: "kustomize.toolkit.fluxcd.io/v1beta2",
								Kind:       "Kustomization",
							},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-5",
								Namespace: "default",
							},
							Spec: kustomizev1.KustomizationSpec{},
							Status: kustomizev1.KustomizationStatus{
								Conditions: []metav1.Condition{
									{
										Status: "Unknown",
									},
									{},
								},
							},
						},
					},
				},
			},
			name:             "standart",
			expectedNumber:   5,
			expectedReady:    3,
			expectedNotReady: 1,
			expectedUnknown:  1,
			namespace:        "default",
		},
	}
	logger.InitLogger()
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			var fluxKustomizations = schema.GroupVersionResource{Group: "kustomize.toolkit.fluxcd.io", Version: "v1beta2", Resource: "kustomizations"}
			testScheme := runtime.NewScheme()
			err := kustomizev1.AddToScheme(testScheme)
			if err != nil {
				t.Fatalf("unexpected error while adding to Scheme: %v", err)
			}
			fluxKustomizationsInfo := &responsestruct.FluxKustomizationsResponse{}
			fakeClientSet := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(testScheme,
				map[schema.GroupVersionResource]string{
					fluxKustomizations: "KustomizationList"}, test.clusterFluxKustomizations...)
			err = ListFluxKustomization(fakeClientSet, test.namespace, fluxKustomizationsInfo)
			if fluxKustomizationsInfo.FluxKustomizationsNumber != test.expectedNumber {
				t.Fatalf("Number of FluxKustomizations in namespace %s - %d are not equal expected value %d", test.namespace,
					fluxKustomizationsInfo.FluxKustomizationsNumber,
					test.expectedNumber)
			}
			if fluxKustomizationsInfo.FluxKustomizationsNumberbyStatus.Ready != test.expectedReady {
				t.Fatalf("Number of running FluxKustomizations in namespace %s - %d are not equal expected value %d", test.namespace,
					fluxKustomizationsInfo.FluxKustomizationsNumberbyStatus.Ready,
					test.expectedReady)
			}
			if fluxKustomizationsInfo.FluxKustomizationsNumberbyStatus.NotReady != test.expectedNotReady {
				t.Fatalf("Number of succeeded FluxKustomizations in namespace %s - %d are not equal expected value %d", test.namespace,
					fluxKustomizationsInfo.FluxKustomizationsNumberbyStatus.NotReady,
					test.expectedNotReady)
			}
			if fluxKustomizationsInfo.FluxKustomizationsNumberbyStatus.Unknown != test.expectedUnknown {
				t.Fatalf("Number of failed FluxKustomizations in namespace %s - %d are not equal expected value %d", test.namespace,
					fluxKustomizationsInfo.FluxKustomizationsNumberbyStatus.Unknown,
					test.expectedUnknown)
			}
			if err != nil {
				t.Fatalf("unexpected error while counting namespaces: %v", err)
			}
		})
	}
}

func TestListHelmreleases(t *testing.T) {
	testCases := []struct {
		name                   string
		clusterFluxHelmrelease []runtime.Object
		expectedNumber         int
		expectedReady          int
		expectedNotReady       int
		expectedUnknown        int
		namespace              string
	}{
		{
			clusterFluxHelmrelease: []runtime.Object{
				&helmv2.HelmReleaseList{
					TypeMeta: metav1.TypeMeta{
						APIVersion: "helm.toolkit.fluxcd.io/v2",
						Kind:       "HelmReleaseList",
					},
					ListMeta: metav1.ListMeta{},
					Items: []helmv2.HelmRelease{
						{
							TypeMeta: metav1.TypeMeta{
								APIVersion: "helm.toolkit.fluxcd.io/v2",
								Kind:       "HelmRelease",
							},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-1",
								Namespace: "default",
							},
							Spec: helmv2.HelmReleaseSpec{},
							Status: helmv2.HelmReleaseStatus{
								Conditions: []metav1.Condition{
									{
										Status: "True",
									},
									{},
								},
							},
						},
						{
							TypeMeta: metav1.TypeMeta{
								APIVersion: "helm.toolkit.fluxcd.io/v2",
								Kind:       "HelmRelease",
							},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-2",
								Namespace: "default",
							},
							Spec: helmv2.HelmReleaseSpec{},
							Status: helmv2.HelmReleaseStatus{
								Conditions: []metav1.Condition{
									{
										Status: "True",
									},
									{},
								},
							},
						},
						{
							TypeMeta: metav1.TypeMeta{
								APIVersion: "helm.toolkit.fluxcd.io/v2",
								Kind:       "HelmRelease",
							},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-3",
								Namespace: "default",
							},
							Spec: helmv2.HelmReleaseSpec{},
							Status: helmv2.HelmReleaseStatus{
								Conditions: []metav1.Condition{
									{
										Status: "False",
									},
									{},
								},
							},
						},
						{
							TypeMeta: metav1.TypeMeta{
								APIVersion: "helm.toolkit.fluxcd.io/v2",
								Kind:       "HelmRelease",
							},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-4",
								Namespace: "default",
							},
							Spec: helmv2.HelmReleaseSpec{},
							Status: helmv2.HelmReleaseStatus{
								Conditions: []metav1.Condition{
									{
										Status: "False",
									},
									{},
								},
							},
						},
						{
							TypeMeta: metav1.TypeMeta{
								APIVersion: "helm.toolkit.fluxcd.io/v2",
								Kind:       "HelmRelease",
							},
							ObjectMeta: metav1.ObjectMeta{
								Name:      "default-5",
								Namespace: "default",
							},
							Spec: helmv2.HelmReleaseSpec{},
							Status: helmv2.HelmReleaseStatus{
								Conditions: []metav1.Condition{
									{
										Status: "Unknown",
									},
									{},
								},
							},
						},
					},
				},
			},
			name:             "standart",
			expectedNumber:   5,
			expectedReady:    2,
			expectedNotReady: 2,
			expectedUnknown:  1,
			namespace:        "default",
		},
	}
	logger.InitLogger()
	for _, test := range testCases {
		//	t.Run(test.name, func(t *testing.T) {
		var fluxHelmreleases = schema.GroupVersionResource{Group: "helm.toolkit.fluxcd.io", Version: "v2", Resource: "helmreleases"}
		testScheme := runtime.NewScheme()
		err := helmv2.AddToScheme(testScheme)
		if err != nil {
			t.Fatalf("unexpected error while adding to Scheme: %v", err)
		}
		fluxHelmreleasesInfo := &responsestruct.FluxHelmreleasesResponse{}
		fakeClientSet := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(testScheme,
			map[schema.GroupVersionResource]string{
				fluxHelmreleases: "HelmReleaseList"}, test.clusterFluxHelmrelease...)
		err = ListHelmreleases(fakeClientSet, test.namespace, fluxHelmreleasesInfo)
		if fluxHelmreleasesInfo.FluxHelmreleasesNumber != test.expectedNumber {
			t.Fatalf("Number of FluxHelmreleases in namespace %s - %d are not equal expected value %d", test.namespace,
				fluxHelmreleasesInfo.FluxHelmreleasesNumber,
				test.expectedNumber)
		}
		if fluxHelmreleasesInfo.FluxHelmreleasesNumberbyStatus.Ready != test.expectedReady {
			t.Fatalf("Number of running FluxHelmreleases in namespace %s - %d are not equal expected value %d", test.namespace,
				fluxHelmreleasesInfo.FluxHelmreleasesNumberbyStatus.Ready,
				test.expectedReady)
		}
		if fluxHelmreleasesInfo.FluxHelmreleasesNumberbyStatus.NotReady != test.expectedNotReady {
			t.Fatalf("Number of succeeded FluxHelmreleases in namespace %s - %d are not equal expected value %d", test.namespace,
				fluxHelmreleasesInfo.FluxHelmreleasesNumberbyStatus.NotReady,
				test.expectedNotReady)
		}
		if fluxHelmreleasesInfo.FluxHelmreleasesNumberbyStatus.Unknown != test.expectedUnknown {
			t.Fatalf("Number of failed FluxHelmreleases in namespace %s - %d are not equal expected value %d", test.namespace,
				fluxHelmreleasesInfo.FluxHelmreleasesNumberbyStatus.Unknown,
				test.expectedUnknown)
		}
		if err != nil {
			t.Fatalf("unexpected error while counting namespaces: %v", err)
		}
		//	})
	}
}

// func TestDynamicClient(t *testing.T) {
// 	// Setup an Object as mock on the client
// 	// Write it like its YAML manifest
// 	mdb := &unstructured.Unstructured{}
// 	mdb.SetUnstructuredContent(map[string]interface{}{
// 		"apiVersion": "mongodbcommunity.mongodb.com/v1",
// 		"kind":       "MongoDBCommunity",
// 		"metadata": map[string]interface{}{
// 			"name":      "mongodb-test",
// 			"namespace": "default",
// 		},
// 		"spec": map[string]interface{}{
// 			"members": 3,
// 		},
// 	})

// 	dynamicClient := dynamicfake.NewSimpleDynamicClient(runtime.NewScheme(), mdb)

// 	// Run any logic that depend on the dynamic client
// 	NotifyMongoDBs(context.Background(), dynamicClient)

// 	AssertActions(t, dynamicClient.Actions(), []ExpectedAction{
// 		{
// 			Verb:      "list",
// 			Namespace: "default",
// 			Resource:  "mongodbcommunity",
// 		},
// 	})

// }
