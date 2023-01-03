package main

import (
	"flag"
	"fmt"
	"runtime"
)

type Config struct {
	WorkersCount int
	ThreadsCount int
}

func (c Config) String() string {
	return fmt.Sprintf("{Workers:%v, Threads:%v}", c.WorkersCount, c.ThreadsCount)
}

func ParseConfig() Config {
	workers := flag.Int("workers", runtime.NumCPU(), "A size of worker pool")
	threads := flag.Int("threads", runtime.NumCPU(), "A size of thread pool, aka GOMAXPROCS")
	flag.Parse()
	return Config{*workers, *threads}
}
