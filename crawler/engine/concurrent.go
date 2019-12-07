package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler    Scheduler
	WorkereCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkereCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	count := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("%d Got item: %v", count, item)
			count++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, e := worker(request)
			if e != nil {
				continue
			}
			out <- result
		}
	}()
}
