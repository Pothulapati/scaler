package prometheus

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"os"
	"time"
)



// GetMetrics returns the value of the deployment name
func GetMetrics(addr string, pod string) (float64, error) {
	prom, err := api.NewClient(api.Config{Address:addr})
	if err != nil {
		return -1, err
	}

	v1api := v1.NewAPI(prom)

	result, warnings, err := v1api.Query(context.Background(), fmt.Sprintf("avg(container_memory_rss{ pod_name=\"%s\", image=~\"tarun.*\"}/1024/1024)", pod), time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Printf("\nValue of Query %s:%v",fmt.Sprintf("avg(container_memory_rss{ pod_name=\"%s\", image=~\"tarun.*\"}/1024/1024)", pod), result)

	if len(result.(model.Vector)) >= 1 {
		return float64(result.(model.Vector)[0].Value), nil
	}

	return 0, errors.New(fmt.Sprintf("Zero metrics found for ", pod))
}
