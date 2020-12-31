package calc

type calculator struct{}

// Calculator is the interface that wraps the basic calculate method.
type Calculator interface {
	Calc(a, b int) int
}

// Calc is calculated two numbers.
func (c *calculator) Calc(a, b int) int {
	return a + b
}

// NewCalculator returns a Calculator that calculate inputed numbers.
func NewCalculator() Calculator {
	return &calculator{}
}
