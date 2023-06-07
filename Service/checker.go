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
	Targets []Checkable
}

// check performs a Health check on the passed Checkable interface
func (ch *Checker) check(ctx context.Context, c Checkable) {
	if !c.Health(ctx) {
		fmt.Println(c.GetID(), FailStatus)
	}
}

func (ch *Checker) Add(item Checkable) {
	ch.Targets = append(ch.Targets, item)
}

func (ch *Checker) String() string {
	var result strings.Builder
	for _, value := range ch.Targets {
		result.WriteString(fmt.Sprintf("%s ", value.GetID()))

	}
	return result.String()
}

func (ch *Checker) Check() {
	for _, value := range ch.Targets {
		fmt.Println(value.Health(context.Background()))
	}
}

func MakeChecker() *Checker {
	return &Checker{}
}
