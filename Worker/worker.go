package Worker

import "fmt"

type Worker[T any] struct {
	jobs chan T
}

type WorkerConfig[T any] struct {
	Workers int
	Execute func(T)
}

func NewWorker[T any](conf WorkerConfig[T]) Worker[T] {
	var wkr = Worker[T]{}
	wkr.jobs = make(chan T, conf.Workers)

	for w := 1; w <= conf.Workers; w++ {
		fmt.Println("Worker", w, "started")
		go worker(w, wkr.jobs, conf.Execute)
	}

	return wkr
}

func worker[T any](id int, jobs <-chan T, run func(T)) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job")
		run(j)
	}
}

func (worker *Worker[T]) Add(parameter T) {
	worker.jobs <- parameter
}
