package arithmetic

// Adapter ...
type Adapter struct{}

// NewAdapter ...
func NewAdapter() *Adapter {
	return &Adapter{}
}

// Add ...
func (ad Adapter) Add(a int32, b int32) (int32, error) {
	return a + b, nil
}

// Subtract ...
func (ad Adapter) Subtract(a int32, b int32) (int32, error) {
	return a - b, nil
}

// Multiply ...
func (ad Adapter) Multiply(a int32, b int32) (int32, error) {
	return a * b, nil
}

// Divide ...
func (ad Adapter) Divide(a int32, b int32) (int32, error) {
	return a / b, nil
}