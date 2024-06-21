package main

import (
	"fmt"
	"time"

	"github.com/wwiechorek/go-worker-pool/Worker"
)

func main() {
	wrk := Worker.NewWorker(Worker.WorkerConfig[int]{
		Workers: 5,
		Execute: WorkerExecute,
	})

	for i := 0; i < 100; i++ {
		wrk.Add(i)
	}
}

func WorkerExecute(s int) {
	fmt.Println("Worker execute", s)
	time.Sleep(time.Second)
}
