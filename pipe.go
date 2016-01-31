package gogen

// Pipeline is a set of stages that must be run
// in order to get the result
type Pipeline struct {
	generators []Generable
}

// Add will add passed generator into the pipeline
func (p *Pipeline) Add(gen Generable) {
	p.generators = append(p.generators, gen)
}

// Size returns the number of generators that are
// in the pipelne
func (p *Pipeline) Size() int {
	return len(p.generators)
}
