package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

// GoMetrClient structure
type GoMetrClient struct {
	ServiceID string
	TimeOut   time.Duration
}

// GetID returns the service identifier
func (g *GoMetrClient) GetID() string {
	return g.ServiceID
}

// getHealth simulates the health check and returns a HealthCheck structure
func (g *GoMetrClient) getHealth(ctx context.Context) HealthCheck {
	// Simulating health check results
	// You can modify this logic to perform actual health checks
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	return HealthCheck{
		ServiceID: g.ServiceID,
		Status:    PassStatus,
	}
}

// Health calls the getHealth method and returns the health status
func (g *GoMetrClient) Health(ctx context.Context) bool {
	result := make(chan bool, 1)
	go func() {
		result <- g.getHealth(ctx).Status == PassStatus
	}()
	select {
	case <-time.After(g.TimeOut):
		log.Println("Health check timed out, returning false")
		return false
	case result := <-result:
		log.Println("Health check result:", result)
		return result
	}

}

// GetMetrics returns the metrics of the service
func (g *GoMetrClient) GetMetrics(ctx context.Context) string {
	// Implement the metrics retrieval logic here
	return ""
}

// NewGoMetrClient is a constructor function for GoMetrClient
func NewGoMetrClient(serviceID string, to int) *GoMetrClient {
	return &GoMetrClient{
		ServiceID: serviceID,
		TimeOut:   time.Duration(to) * time.Second,
	}
}
