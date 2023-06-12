package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
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
	sync.Mutex
	cf context.CancelFunc
}

// check performs a Health check on the passed Checkable interface
func (ch *Checker) check(ctx context.Context, c Checkable) {
	if !c.Health(ctx) {
		fmt.Println(c.GetID(), FailStatus)
	}
}

func (ch *Checker) Add(item Checkable) {
	ch.Lock()
	defer ch.Unlock()
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
	return &Checker{
		Mutex:   sync.Mutex{},
		targets: make([]Checkable, 0),
	}
}

func (ch *Checker) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	go ch.run(ctx)
	log.Println("Checker started")
	ch.cf = cancel
	return
}

func (ch *Checker) Stop() {
	if ch.cf != nil {
		ch.cf()
	}
	log.Println("Checker stopped")
}

func (ch *Checker) run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ch.checkAll(ctx)
		}
	}
}

func (ch *Checker) checkAll(ctx context.Context) {
	ch.Lock()
	defer ch.Unlock()
	for _, item := range ch.targets {
		select {
		case <-ctx.Done():
			return
		default:
			go item.Health(ctx)
		}
	}
}
