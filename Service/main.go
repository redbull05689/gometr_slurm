package main

import "fmt"

func main() {
	ch := MakeChecker()
	sid1 := NewGoMetrClient("1")
	sid2 := NewGoMetrClient("2")
	ch.Add(sid1)
	ch.Add(sid2)

	fmt.Println(ch)
	ch.Check()
}
