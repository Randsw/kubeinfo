package cmd

import (
	"reflect"
	"strings"
	"testing"

	"github.com/randsw/kubeinfo/logger"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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
			var NsNames []string
			fakeClientSet := fake.NewSimpleClientset(test.clusterNS...)
			namespaces, err := GetAllNamespaces(fakeClientSet)
			for _, namespace := range namespaces.Items {
				NsNames = append(NsNames, namespace.Name)
			}
			if !reflect.DeepEqual(test.expectedNS, NsNames) {
				t.Fatalf("Namespace names %s not as expected %s", strings.Join(NsNames, ","), strings.Join(test.expectedNS, ","))
			}
			if len(namespaces.Items) != test.expectedNumber {
				t.Fatalf("Number of namespaces %d are not equal expected value %d", len(namespaces.Items), test.expectedNumber)
			}
			if err != nil {
				t.Fatalf("unexpected error while counting namespaces: %v", err)
			}
		})
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
