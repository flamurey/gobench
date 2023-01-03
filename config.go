package main

import (
	"flag"
	"runtime"
)

type Config struct {
	WorkersCount int
	ThreadsCount int
}

func ParseConfig() Config {
	single := flag.Bool("single", false, "Whether run it just 1 worker and 1 thread")
	workers := flag.Int("workers", runtime.NumCPU(), "A size of worker pool")
	threads := flag.Int("threads", runtime.NumCPU(), "A size of thread pool, aka GOMAXPROCS")
	flag.Parse()
	if *single {
		return Config{1, 1}
	}
	return Config{*workers, *threads}
}
