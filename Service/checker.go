package main

import (
	"context"
	"fmt"
	"strings"
)

// Checkable interface
type Checkable interface {
	GetID() string
	Health(context.Context) bool
	GetMetrics(context.Context) string
}

// Checker structure
type Checker struct {
	targets []Checkable
}

// check performs a Health check on the passed Checkable interface
func (ch *Checker) check(ctx context.Context, c Checkable) {
	if !c.Health(ctx) {
		fmt.Println(c.GetID(), FailStatus)
	}
}

func (ch *Checker) Add(item Checkable) {
	ch.targets = append(ch.targets, item)
}

func (ch *Checker) String() string {
	var result strings.Builder
	for _, value := range ch.targets {
		result.WriteString(fmt.Sprintf("%s ", value.GetID()))

	}
	return result.String()
}

func (ch *Checker) Check() {
	for _, value := range ch.targets {
		fmt.Println(value.Health(context.Background()))
	}
}

func MakeChecker() *Checker {
	return &Checker{}
}
