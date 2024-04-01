package dynamic

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

// GetPodListByDynamic 通过 dynamic client 获取 pod 列表
func GetPodListByDynamic(kubeconfig string) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}
	unStructData, err := dynamicClient.
		Resource(gvr).
		Namespace("kube-system").
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unStructData.UnstructuredContent(), podList)
	if err != nil {
		panic(err)
	}
	for _, pod := range podList.Items {
		log.Printf("namespace: %s, name: %s\n", pod.Namespace, pod.Name)
	}
}
