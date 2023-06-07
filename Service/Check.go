package main

import (
	"fmt"
)

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

type HealthCheck struct {
	ServiceID string
	Status    string
}

func MakeHealthCheck() *HealthCheck {

	check := HealthCheck{
		ServiceID: "example-service",
		Status:    PassStatus,
	}

	fmt.Println("Updated status:", check.Status)
	return &check
}
