package mpesa


type Service struct {
	config *Configuration
}

func NewService(options map[string]interface{}) *Service {
	service := new(Service)
	service.config = NewConfiguration(options)

	return service
}

func (s *Service) HandleReceive(intent map[string]string) (*Response, *Error) {
	return s.handleRequest(C2bPayment, intent)
}

func (s *Service) HandleSend(intent map[string]string) (*Response, *Error) {
	opcode := s.detectOperation(intent)
	return s.handleRequest(opcode, intent)
}

func (s *Service) HandleRevert(intent map[string]string) (*Response, *Error) {
	return s.handleRequest(Reversal, intent)
}

func (s *Service) HandleQuery(intent map[string]string) (*Response, *Error) {

	return s.handleRequest(QueryTransactionStatus, intent)
}

func (s *Service) handleRequest(opcode Opcode, intent map[string]string) (*Response, *Error) {
	
	return new(Response), new(Error)
}

func (s *Service) detectOperation(intent map[string]string) Opcode {
	return B2cPayment
}
