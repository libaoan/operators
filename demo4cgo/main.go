package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/stan/.kube/config")
	if err != nil {
		log.Fatal("get config error", err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("get client error", err)
	}

	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "container",
					Image: "nginx",
				},
			},
		},
	}
	pod.Name = "nginx2"
	pod.Namespace = "default"

	podTest, err := client.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(podTest)

	podTest2, err := client.CoreV1().Pods("default").Get(context.TODO(), "nginx2", metav1.GetOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(podTest2)

	//factory := informers.NewSharedInformerFactory(client, 0)
	//nodeInformer := factory.Core().V1().Nodes()
	//nodeController := pkg.NewControllerNode(client, nodeInformer)
	//var stopCh = make(chan struct{})
	//nodeController.Run(stopCh)
}
