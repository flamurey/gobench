package main

import (
	"fmt"
	"runtime"
)

func main() {
	config := ParseConfig()
	fmt.Printf("Running bench with config %+v \n", config)
	runtime.GOMAXPROCS(config.ThreadsCount)

	testPlan := TestPlan{config.WorkersCount, LighPayload}
	results := testPlan.Run()
	fmt.Println("Results of light payload")
	results.Print()

	testPlan = TestPlan{config.WorkersCount, HardPayload}
	fmt.Println("Results of hard payload")
	results = testPlan.Run()
	results.Print()
}
