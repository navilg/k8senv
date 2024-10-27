package ikubernetes

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetK8sVersion() (*string, error) {

	var kubeconfig *rest.Config
	homeDir, _ := os.UserHomeDir()
	var err error = nil

	if kubeconfigEnv := os.Getenv("KUBECONFIG"); kubeconfigEnv != "" {
		kubeconfig, err = clientcmd.BuildConfigFromFlags("", kubeconfigEnv)
		if err != nil {
			return nil, err
		}
	} else if _, err := os.Stat(filepath.Join(homeDir, ".kube", "config")); err == nil {
		kubeconfig, err = clientcmd.BuildConfigFromFlags("", filepath.Join(homeDir, ".kube", "config"))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("KUBECONFIG file not found")
	}

	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		log.Fatalf("Error setting up K8s client")
	}

	k8sVersion, err := clientset.ServerVersion()
	if err != nil {
		log.Fatalf("Error getting K8s server version: %v", err)
	}

	k8sSemversion := k8sVersion.String()

	return &k8sSemversion, nil

}
