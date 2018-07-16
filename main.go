package main

import (
	"github.com/kubeflow/tf-operator/pkg/client/clientset/versioned"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	tfjobClient := versioned.NewForConfigOrDie(restConfig)

	tfjobList, err := tfjobClient.KubeflowV1alpha2().TFJobs(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("failed due to %v", err)
	}

}
