package main

import (
	"context"
	"time"
)

type GoogleMetrClient struct {
	ServiceID string
	timeOut   time.Duration
}

func (g *GoogleMetrClient) GetID() string {
	return g.ServiceID
}

func (g *GoogleMetrClient) Health(context.Context) bool {
	return true
}

func (g *GoogleMetrClient) GetMetrics(context.Context) string {
	return ""
}

func NewGoogleMetrClient(serviceID string, to time.Duration) *GoMetrClient {
	return &GoMetrClient{
		ServiceID: serviceID,
		TimeOut:   to,
	}
}
