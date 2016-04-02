package gogen

import (
	"os"
	"sync"

	"github.com/alecthomas/kingpin"
	"github.com/op/go-logging"
)

var (
	// resources is set resources that were firstly defined
	//resources ResourceContainer

	// pipes is set of pipelines that should be run when
	// generate is called
	pipes []Pipeline

	// importers is a collection of available importers in
	// the current session of generator
	//importers unk.ImporterCollection

	// cmdSources is a set of sources that was passed by the
	// command line. All these sources will be recognized by
	// the importers
	cmdSources = kingpin.Flag("source", "List of importable sources").Short('s').Strings()
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

// Pipe will register new pipe that will be run
// in parallel
func Pipe(gens ...Generable) {
	pipe := Pipeline{}
	for _, gen := range gens {
		pipe.Add(gen)
	}

	pipes = append(pipes, pipe)
}

// Generate will startup the generating process
// Note: this function will panic instead of returning
// 		error. This behavior is intended so it easier for
//		users to write configs
func Generate() {
	genlog.Info("Starting gogen")

	// parse arguments from the command line
	kingpin.Parse()

	// define unknown resources
	for _, srcPath := range *cmdSources {
		genlog.Info("Resolving resource [%s]", srcPath)
		// get the source and resolve it
		//src := unk.NewSource(srcPath)
		//err := src.Resolve()
		//if err != nil {
		//	panic("Given source could not be resolved")
		//}
	}

	// go through each pipe
	wg := sync.WaitGroup{}
	for pipeindex, pipe := range pipes {
		wg.Add(1)
		go func(pipe Pipeline, pipeindex int) {
			for _, gen := range pipe.generators {
				genlog.Info("Starting generator %s in pipe %d", gen.Name(), pipeindex)
				//gen.Initialize(&resources, genlog)

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
