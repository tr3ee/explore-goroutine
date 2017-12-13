package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 100

func HeavyWork(id int) {
	rand.Seed(int64(id))
	interval := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(interval)
	fmt.Printf("HeavyWork %-3d cost %v\n", id, interval)
}

func main() {
    wrapper := func() chan int {
        c := make(chan int)
        go func() {
            HeavyWork(0)
            c <- 1
        }()
        return c
    }
    select {
    case <-wrapper():
    case <-time.After(1 * time.Second):
        fmt.Println("time limit exceed")
    }
    // time.Sleep(3 * time.Second)
}