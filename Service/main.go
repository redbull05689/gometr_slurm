package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := MakeChecker()
	sid1 := NewGoMetrClient("1", 2)
	sid2 := NewGoMetrClient("2", 2)
	sid3 := NewGoMetrClient("3", 2)
	ch.Add(sid1)
	ch.Add(sid2)
	ch.Add(sid3)

	fmt.Println(ch)
	ch.Check()
	ctx := context.Background()
	ch.Run(ctx)
	time.Sleep(50 * time.Second)
	ch.Stop()
	time.Sleep(2 * time.Second)
}
