package template

// Generator interface
type Generator interface {
	Generate() (string, error)
}

// Generate fn
func Generate(g Generator) (string, error) {
	return g.Generate()
}
