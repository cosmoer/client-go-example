package main

import "client-go-example/pkg/dynamic"

func main() {
	dynamic.GetPodListByDynamic("/root/.kube/config")
	//clientset.GetPodListByClientSet("/root/.kube/config")
	//restclient.GetPodListByRestClient("/root/.kube/config")
}
