package ports

// GRPCPort ...
type GRPCPort interface {
	Run()
	GetAdd()
	GetSubtract()
	GetMultiply()
	GetDivide()
}
