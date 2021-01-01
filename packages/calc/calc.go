package calc

// Calculator is the struct for execute basic calculate.
type Calculator struct {
	alu ALU
}

// ALU is the interface that wraps the basic calculate method.
type ALU interface {
	Calc(a, b int) int
}

// Calc is calculated two numbers.
func (c *Calculator) Calc(a, b int) int {
	return a + b
}

// NewCalculator returns a Calculator that calculate inputed numbers.
func NewCalculator() *Calculator {
	return &Calculator{}
}
