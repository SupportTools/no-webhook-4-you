package main

import (
	"fmt"
	"time"

	"github.com/golang/glog"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		glog.Errorln(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Errorln(err)
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(clientset, time.Second*30)

	validatingWebhookInformer := kubeInformerFactory.Admissionregistration().V1().ValidatingWebhookConfigurations()
	validatingWebhookInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("ValidatingWebhookConfiguration added: %s \n", obj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Printf("ValidatingWebhookConfiguration deleted: %s \n", obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Printf("ValidatingWebhookConfiguration changed: %s \n", newObj)
		},
	})

	mutatingWebhookInformer := kubeInformerFactory.Admissionregistration().V1().MutatingWebhookConfigurations()
	mutatingWebhookInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("MutatingWebhookConfigurations added: %s \n", obj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Printf("MutatingWebhookConfigurations deleted: %s \n", obj)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Printf("MutatingWebhookConfigurations changed: %s \n", newObj)
		},
	})

	stop := make(chan struct{})
	defer close(stop)
	kubeInformerFactory.Start(stop)
	for {
		time.Sleep(time.Second)
	}
}
