package tests // import "kego.io/process/validate/selectors/tests"

// for tests
type Image interface {
	Something() bool
}

func (d *Diagram) Something() bool {
	return false
}

func (p *Photo) Something() bool {
	return false
}