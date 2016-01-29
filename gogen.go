package gogen

import (
	"os"
	"sync"

	"github.com/op/go-logging"
)

var (
	// Resources is set resources that were firstly defined
	Resources ResourceContainer

	// Pipes is set of pipelines that should be run when
	// generate is called
	Pipes []Pipeline
)

var (
	genlog = logging.MustGetLogger("gogen")

	logFormat = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
)

// initialize logging
func init() {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	formatter := logging.NewBackendFormatter(backend, logFormat)

	logging.SetBackend(formatter)
}

// Define will store the defined model for the use in
// the generators.
func Define(resource interface{}) {
	switch val := resource.(type) {
	case RemoteResource:
		// append all fetched resources from the remote
		Resources = append(Resources, val.Get()...)
	default:
		// not known resource, append
		Resources = append(Resources, resource)
	}
}

// Pipe will register new pipe that will be run
// in parallel
func Pipe(gens ...Generable) {
	pipe := Pipeline{}
	for _, gen := range gens {
		pipe.Add(gen)
	}

	Pipes = append(Pipes, pipe)
}

// Generate will startup the generating process
// Note: this function will panic instead of returning
// 		error. This behavior is intended so it easier for
//		users to write configs
func Generate() {
	genlog.Info("Starting gogen")

	wg := sync.WaitGroup{}

	for pipeindex, pipe := range Pipes {
		wg.Add(1)
		go func(pipe Pipeline, pipeindex int) {
			for _, gen := range pipe.generators {
				genlog.Info("Starting generator %s in pipe %d", gen.Name(), pipeindex)
				gen.Initialize(&Resources)

				err := gen.Generate()
				if err != nil {
					panic(err)
				}

				err = gen.Output()
				if err != nil {
					panic(err)
				}

				genlog.Info("End of generator %s in pipe %d", gen.Name(), pipeindex)
			}

			wg.Done()
		}(pipe, pipeindex)
	}

	wg.Wait()
}
