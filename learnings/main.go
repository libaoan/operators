package main

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"

	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// restfulClient()
	clientSet()

}

func restfulClient() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/stan/.kube/config")
	if err != nil {
		println(err.Error())
	}

	// client
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		print(err.Error())
	} else {
		pod := v1.Pod{}
		err := restClient.Get().Namespace("default").Resource("pods").Name(
			"nginx").Do(context.TODO()).Into(&pod)
		if err != nil {
			println("err", err.Error())
		} else {
			println("hello", pod.Name, pod.Spec.NodeName, pod.Status.HostIP)
		}
	}
}

func clientSet() {

	config, err := clientcmd.BuildConfigFromFlags("", "/Users/stan/.kube/config")
	if err != nil {
		println(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		println(err.Error())
	}
	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
	if err != nil {
		println(err.Error())
	} else {
		println(pod.Name, pod.Spec.NodeName, pod.Status.HostIP)
	}
}