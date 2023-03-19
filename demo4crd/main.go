package main

import (
	clientset "github.com/operators/demo4crd/pkg/generated/clientset/versioned"
	"github.com/operators/demo4crd/pkg/generated/informers/externalversions"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/stan/.kube/config")
	if err != nil {
		log.Fatalln(err)
	}
	client, err := clientset.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	factory := externalversions.NewSharedInformerFactory(client, 0)
	fooInformer := factory.Crd().V1().Foos()
	fooInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {},
	})
}
