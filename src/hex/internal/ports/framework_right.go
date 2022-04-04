package ports

// DbPort ...
type DbPort interface {
	CloseDbConnection()
	AddToHistory(value int32, operation string) error
}
