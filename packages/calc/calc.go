package calc

// ALU is the interface that wraps the basic calculate method.
type ALU interface {
	Calc(a, b int) int
}

// Calculator is the struct for execute basic calculate.
type Calculator struct {
	alu CalcALU
}

type CalcALU struct{}

// Calc is calculated two numbers.
func (c *CalcALU) Calc(a, b int) int {
	return a + b
}

// NewCalculator returns a Calculator that calculate inputed numbers.
func NewCalculator() *Calculator {
	return &Calculator{}
}
