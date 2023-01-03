package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const TestDurataion = 2 * time.Second

type Results struct {
	Total   uint64
	Average time.Duration
}

func (r Results) Print() {
	fmt.Printf("Total:%10d \t Average:%10v \n", r.Total, r.Average)
}

type TestPlan struct {
	WorkerPoolSize int
	Playload       func()
}

func (tp *TestPlan) Run() Results {
	ctx, cancel := context.WithTimeout(context.Background(), TestDurataion)
	defer cancel()
	var total uint64
	wg := sync.WaitGroup{}
	wg.Add(tp.WorkerPoolSize)
	for i := 0; i < tp.WorkerPoolSize; i++ {
		go func() {
		forever:
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					break forever
				default:
					tp.Playload()
					atomic.AddUint64(&total, 1)
				}
			}

		}()
	}
	wg.Wait()

	avg := TestDurataion / time.Duration(total)
	return Results{total, avg}
}
