package api

import "github.com/adetiamarhadi/golang-hex-arch/internal/ports"

// Adapter ...
type Adapter struct {
	arith ports.ArithmeticPort
}

// NewAdapter ...
func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

// GetAdd ...
func (ad Adapter) GetAdd(a, b int32) (int32, error) {
	val, err := ad.arith.Add(a, b)
	if err != nil {
		return 0, err
	}

	return val, nil
}

// GetSubtract ...
func (ad Adapter) GetSubtract(a, b int32) (int32, error) {
	val, err := ad.arith.Subtract(a, b)
	if err != nil {
		return 0, err
	}

	return val, nil
}

// GetMultiply ...
func (ad Adapter) GetMultiply(a, b int32) (int32, error) {
	val, err := ad.arith.Multiply(a, b)
	if err != nil {
		return 0, err
	}

	return val, nil
}

// GetDivide ...
func (ad Adapter) GetDivide(a, b int32) (int32, error) {
	val, err := ad.arith.Divide(a, b)
	if err != nil {
		return 0, err
	}

	return val, nil
}
