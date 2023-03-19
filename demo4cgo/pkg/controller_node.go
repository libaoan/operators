package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	v12 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"log"
	"time"
)

type controllerNode struct {
	client     kubernetes.Interface
	nodeLister v1.NodeLister
	queue      workqueue.RateLimitingInterface
}

func (c controllerNode) updateNode(oldObj interface{}, newObj interface{}) {
	fmt.Println("old node", oldObj)
	fmt.Println("new node", newObj)
}

func NewControllerNode(client kubernetes.Interface, nodeInformer v12.NodeInformer) controllerNode {
	c := controllerNode{
		client:     client,
		nodeLister: nodeInformer.Lister(),
		queue:      workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "node_manager"),
	}

	nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		UpdateFunc: c.updateNode,
	})

	return c
}

func (c *controllerNode) Run(stopCh <-chan struct{}) {
	go wait.Until(c.worker, time.Minute, stopCh)
	<-stopCh
}

func (c *controllerNode) worker() {
	log.Println("count...")
	node, err := c.client.CoreV1().Nodes().Get(context.TODO(), "myk8s-worker", v13.GetOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(node.APIVersion)
	log.Println(node.Kind)
	log.Println(node.Namespace)
	log.Println(node.Name)
	log.Println(node.UID)
	log.Println(node.OwnerReferences)
	log.Println(node.Annotations)
	log.Println(node.Spec.String())
	log.Println(node.Status.String())
	log.Println(node.CreationTimestamp.String())
	//log.Println(node.DeletionTimestamp.String())
	log.Println(node.Labels)
	log.Println(node.Finalizers)
	log.Println(node.String())
	s, _ := node.Spec.Marshal()
	var b bytes.Buffer
	err = json.Indent(&b, s, "", "	")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(b.String())
}
