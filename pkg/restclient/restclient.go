package restclient

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetPodListByRestClient  获取 pod 列表, 入参 为 kubeconfig 文件路径，如果为空，则使用 in-cluster config,
// 直接使用 RestClient 去连接 k8s apiserver 并获取 pod 列表, 不使用 clientset。 输出日志使用 client-go 提供的 log 库
func GetPodListByRestClient(kubeconfig string) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	result := &corev1.PodList{}
	err = restClient.Get().
		Namespace("kube-system").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	if err != nil {
		panic(err)
	}
	for _, pod := range result.Items {
		log.Printf("namespace: %s, name: %s\n", pod.Namespace, pod.Name)
	}
}
