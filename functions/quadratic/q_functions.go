package go_quadratic

type quadraticFunction struct {
	a, b, c float32
}

type QuadraticFunctionCalculations interface {
	delta()
	zeros()
}

func (self quadraticFunction) delta() float32 {
	return self.b*self.b - 4*self.a*self.c
}

func (self quadraticFunction) zeros() (float32, float32) {
	return (-self.b - self.delta()) / 2 * self.a,
		(-self.b - self.delta()) / 2 * self.a
}

func getQuadaticFunction() *quadraticFunction {
	return &quadraticFunction{0, 0, 0}
}
