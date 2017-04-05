package worker

import "time"

type Worker struct {
	every time.Duration
	run   func()
	quit  chan struct{}
}

func NewWorker(every time.Duration, f func()) (w *Worker) {
	w = &Worker{
		every: every,
		run:   f,
		quit:  make(chan struct{}),
	}
	return
}

func (w *Worker) Start() {
	ticker := time.NewTicker(w.every)
	for {
		select {
		case <-ticker.C:
			w.run()
		case <-w.quit:
			ticker.Stop()
			return
		}
	}
}
func (w *Worker) Close() {
	close(w.quit)
}
