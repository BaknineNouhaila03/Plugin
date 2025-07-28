package main

import (
	"fmt"
	"log"

	"go-tenant-cluster/stackelberg"
)

func main() {
	// Example total cluster resources (can come from K8s in real use)
	totalCPU := 110.0
	totalMemory := 300.0


	params := map[string]interface{}{
		
    "web_app_max_response_time": 52.5,
    "web_app_budget": 800.0,
    "web_app_min_cpu": 2.0,
    "web_app_min_memory": 8.0,
    "web_app_desired_replicas": 5.0,
    "web_app_min_replicas": 1.0,

    "data_processing_min_throughput": 1000.0,
    "data_processing_budget": 1200.0,
    "data_processing_min_cpu": 1.0,
    "data_processing_min_memory": 4.0,
    "data_processing_desired_replicas": 4.0,
    "data_processing_min_replicas": 1.0,

    "ml_training_max_training_time": 8.0,
    "ml_training_budget": 1500.0,
    "ml_training_min_cpu": 4.0,
    "ml_training_min_memory": 16.0,
    "ml_training_desired_replicas": 2.0,
    "ml_training_min_replicas": 1.0,

    "alpha1": 0.3,
    "alpha2": 0.2,
    "alpha3": 0.5,

    "cpu_norm": 2.0,
    "memory_norm": 4.0,
    "base_exponent": 0.7,
    "rt_const1": 50.0,
    "rt_const2": 500.0,
    "rt_exponent": 0.3,
    "latency_thresh": 52.5,
    "latency_penalty": 100.0,

    "tenant_b_base_coeff": 120.0,
    "tenant_b_memory_exp1": 0.6,
    "tenant_b_base_exp": 0.8,
    "tenant_b_throughput_coeff": 15.0,
    "tenant_b_throughput_cpu_exp": 0.8,
    "tenant_b_throughput_mem_exp": 0.4,
    "tenant_b_queue_penalty_thresh": 1000.0,
    "tenant_b_queue_penalty_coeff": 2.0,

    "tenant_c_base_coeff": 120.0,
    "tenant_c_memory_exp1": 0.5,
    "tenant_c_log_const": 1.0,
    "tenant_c_training_cpu_exp": 0.7,
    "tenant_c_training_mem_exp": 0.3,
    "tenant_c_time_penalty_thresh": 8.0,
    "tenant_c_time_penalty_coeff": 40.0,

    "initial_p_cpu": 5.0,
    "initial_p_memory": 2.0,
  
	}

	// Test connection first
	client := stackelberg.NewClient("http://localhost:5000")
	if err := client.Health(); err != nil {
		log.Fatalf("Python sidecar health check failed: %v", err)
	}

	// Call the allocation service
result, err := client.Allocate(totalCPU, totalMemory, params)
	if err != nil {
		log.Fatalf("Failed to call sidecar: %v", err)
	}

	fmt.Println("--- Stackelberg Allocation Result ---")
	fmt.Printf("Allocations: %+v\n", result.Allocations)
	fmt.Printf("Prices: %+v\n", result.Prices)
	fmt.Printf("Platform Utility: %f\n", result.PlatformUtility)
	fmt.Printf("Metrics: %+v\n", result.Metrics)
}