package mpesa

type ResponseStatus struct {
	Code uint
	Description string
}

type Response struct {
	Success bool
	Status ResponseStatus
	Data map[string]string
}