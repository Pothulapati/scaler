package k8s

import (
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// GetPods returns the number of pods under the deployment
func GetPods(label string) (int, []v12.Pod, error) {

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return 0,nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return 0,nil, err
	}

	pods, err := clientset.CoreV1().Pods("emojivoto").List(v1.ListOptions{LabelSelector:label})
	if err != nil {
		return 0,nil, err
	}

	return len(pods.Items), pods.Items, nil
}

// Updates number of pods
func UpdatePods(new int32) (error) {

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	deployment, err := clientset.AppsV1().Deployments("emojivoto").Get("web", v1.GetOptions{})
	if err != nil {
		return err
	}

	deployment.Spec.Replicas = &new

	_, err = clientset.AppsV1().Deployments("emojivoto").Update(deployment)
	if err != nil {
		return err
	}

	return nil
}
