package main

import (
	"auto-scaler/pkg/k8s"
	"auto-scaler/pkg/prometheus"
	"fmt"
	"math"
)

var (
	desiredMetricValue float64 = 15
)

func main() {

	for {
		// Get Number of Pods
		// use k8s api to get current number of instances
		currentReplicas,pods, err := k8s.GetPods("app=web-svc")
		if err !=nil {
			fmt.Println("Couldn't get k8s metrics: ", err)
			continue
		}

		fmt.Println("Current Number of Instances: ", currentReplicas)

		var currentMetricValue float64
		currentMetricValue = 0
		// Get metric value for each pod
		for _, pod := range pods {
			podValue, err := prometheus.GetMetrics("http://scaler-prometheus.scaler:9090", pod.Name)
			fmt.Println("finding metric for ", pod.Name, "is ", podValue)

			if err !=nil {
				fmt.Println("Couldnt get prometheus metrics:", err)
				continue
			}
			currentMetricValue += podValue
		}

		currentMetricValue = currentMetricValue/float64(currentReplicas)

		fmt.Println("Current Metric Value:",currentMetricValue)


		desiredReplicas := math.Ceil(float64(currentReplicas) * (currentMetricValue / desiredMetricValue))

		fmt.Println("Desired Number of Instances: ", desiredReplicas)

		if int(desiredReplicas) != currentReplicas {

			err = k8s.UpdatePods(int32(desiredReplicas))
			if err!=nil {
				fmt.Println(" Couldn't Update Pods:", err)
				continue
			}
		}

		fmt.Println("Successfully Updated Number of Pods")

		fmt.Println("-----------------------------------")
	}

}