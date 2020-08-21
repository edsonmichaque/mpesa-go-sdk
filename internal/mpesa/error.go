package mpesa

type ErrorType string

const (
	MissingProperties = iota
	ValidationErrors
) 
type Error struct {
	Type ErrorType
}