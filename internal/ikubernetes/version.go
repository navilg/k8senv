package ikubernetes

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
		return nil, err
	}

	k8sVersion, err := clientset.ServerVersion()
	if err != nil {
		return nil, err
	}

	k8sSemversionSplit := strings.Split(k8sVersion.String(), ".")

	re := regexp.MustCompile(`[^0-9]`)
	result := re.ReplaceAllString(k8sSemversionSplit[2], "/")

	resultSplit := strings.Split(result, "/")

	k8sSemversionPatch := resultSplit[0]

	k8sSemversion := k8sSemversionSplit[0] + "." + k8sSemversionSplit[1] + "." + k8sSemversionPatch

	return &k8sSemversion, nil

}

func GetVeleroVersion() (*string, error) {
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
		return nil, err
	}

	namespace := "velero"
	deploymentName := "velero"

	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, errors.New("error retrieving velero deployment from velero namespace")
	}

	veleroImageName := deployment.Spec.Template.Spec.Containers[0].Image

	veleroVersion := strings.Split(veleroImageName, ":")[1]

	semVersionRegex := regexp.MustCompile(`^v?(\d+)\.(\d+)\.(\d+)(?:-([0-9A-Za-z-.]+))?(?:\+([0-9A-Za-z-.]+))?$`)

	if !semVersionRegex.MatchString(veleroVersion) {
		return nil, errors.New("velero image used in velero deployment is not semantic version: " + veleroVersion)
	}

	return &veleroVersion, nil

}
