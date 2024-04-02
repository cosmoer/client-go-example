package main

import "client-go-example/pkg/discovery"

func main() {
	discovery.GetApiResource("/root/.kube/config")
	//dynamic.GetPodListByDynamic("/root/.kube/config")
	//clientset.GetPodListByClientSet("/root/.kube/config")
	//restclient.GetPodListByRestClient("/root/.kube/config")
}
