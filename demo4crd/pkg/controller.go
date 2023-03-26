package pkg

import (
	clientset "github.com/operators/demo4crd/pkg/generated/clientset/versioned"
	v12 "github.com/operators/demo4crd/pkg/generated/informers/externalversions/samplecrd/v1"
	"github.com/operators/demo4crd/pkg/generated/listers/samplecrd/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"log"
)

type controllerFoo struct {
	client    clientset.Interface
	fooLister v1.FooLister
	queue     workqueue.RateLimitingInterface
}

func (c *controllerFoo) addFoo(obj interface{}) {
	log.Println("added", obj)
}

func (c *controllerFoo) updateFoo(oldObj interface{}, newObj interface{}) {
	log.Println("updated", oldObj, newObj)
}

func (c *controllerFoo) deleteFoo(obj interface{}) {
	log.Println("deleted", obj)
}

func (c *controllerFoo) Run(stopCh <-chan struct{}) {

	<-stopCh
}

func NewControllerFoo(client clientset.Interface, fooInformer v12.FooInformer) *controllerFoo {
	c := &controllerFoo{
		client:    client,
		fooLister: fooInformer.Lister(),
		queue:     workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "foo_manager"),
	}
	fooInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addFoo,
		UpdateFunc: c.updateFoo,
		DeleteFunc: c.deleteFoo,
	})

	return c
}
