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
	ok, quit := make(chan int, 1), make(chan int, 1)
	go func() {
		i := 0
		for {
			select {
			case <-quit:
				ok <- 1
				return
			default:
				HeavyWork(i)
				i++
			}
		}
	}()
	time.Sleep(5 * time.Second)
	quit <- 1
	<-ok
}
