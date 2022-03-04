package main

import (
	"context"
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
			name := obj.(*v1.ValidatingWebhookConfiguration).GetName()
			glog.Infof("ValidatingWebhookConfiguration added: %s \n", name)
			if name == "rancher.cattle.io" {
				glog.Infoln("Found rancher.cattle.io")
				err := clientset.AdmissionregistrationV1().ValidatingWebhookConfigurations().Delete(context.TODO(), "rancher.cattle.io", meta_v1.DeleteOptions{})
				if err != nil {
					glog.Errorln(err)
				}
				glog.Infoln("Deleted ValidatingWebhookConfigurations rancher.cattle.io")
			}
		},
	})

	mutatingWebhookInformer := kubeInformerFactory.Admissionregistration().V1().MutatingWebhookConfigurations()
	mutatingWebhookInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			name := obj.(*v1.MutatingWebhookConfiguration).GetName()
			glog.Infof("MutatingWebhookConfigurations added: %s \n", name)
			if name == "rancher.cattle.io" {
				glog.Infoln("Found rancher.cattle.io")
				err := clientset.AdmissionregistrationV1().MutatingWebhookConfigurations().Delete(context.TODO(), "rancher.cattle.io", meta_v1.DeleteOptions{})
				if err != nil {
					glog.Errorln(err)
				}
				glog.Infoln("Deleted MutatingWebhookConfigurations rancher.cattle.io")
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
