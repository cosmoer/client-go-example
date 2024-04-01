package clientset

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// GetPodListByClientSet 通过 clientset 对象获取 Pod 列表
func GetPodListByClientSet(kubeconfig string) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	pods, err := clientset.
		CoreV1().
		Pods("kube-system").
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, pod := range pods.Items {
		log.Printf("namespace: %s, name: %s\n", pod.Namespace, pod.Name)
	}
}
