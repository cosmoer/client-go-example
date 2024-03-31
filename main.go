package main

import "client-go-example/pkg/restclient"

func main() {
	restclient.GetPodListByRestClient("/root/.kube/config")
}
