package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/glog"

	v1 "k8s.io/api/admissionregistration/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
			if obj.(*v1.ValidatingWebhookConfiguration).Name == "rancher.cattle.io" {
				fmt.Println("Found rancher.cattle.io")
				err := clientset.AdmissionregistrationV1().ValidatingWebhookConfigurations().Delete(context.TODO(), "rancher.cattle.io", meta_v1.DeleteOptions{})
				if err != nil {
					glog.Errorln(err)
				}
				fmt.Println("Deleted ValidatingWebhookConfigurations rancher.cattle.io")
			}
		},
	})

	mutatingWebhookInformer := kubeInformerFactory.Admissionregistration().V1().MutatingWebhookConfigurations()
	mutatingWebhookInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("MutatingWebhookConfigurations added: %s \n", obj)
			if obj.(*v1.MutatingWebhookConfiguration).Name == "rancher.cattle.io" {
				fmt.Println("Found rancher.cattle.io")
				err := clientset.AdmissionregistrationV1().MutatingWebhookConfigurations().Delete(context.TODO(), "rancher.cattle.io", meta_v1.DeleteOptions{})
				if err != nil {
					glog.Errorln(err)
				}
				fmt.Println("Deleted MutatingWebhookConfigurations rancher.cattle.io")
			}
		},
	})

	stop := make(chan struct{})
	defer close(stop)
	kubeInformerFactory.Start(stop)
	for {
		time.Sleep(time.Second)
	}
}
