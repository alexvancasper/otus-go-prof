package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil {
		newIn := make(chan interface{})
		close(newIn)
		return newIn
	}
	for _, stage := range stages {
		in = stage(checkDone(in, done))
	}
	return in
}

func checkDone(in In, done In) Out {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}

// https://go.dev/blog/pipelines
