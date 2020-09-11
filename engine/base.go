package engine

type Engine interface {

	// Methods
	Set() error
	Generate() error
}
