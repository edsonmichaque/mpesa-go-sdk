package mpesa


type Service struct {
	config *Configuration
}

func NewService(options map[string]interface{}) *Service {
	service := new(Service)
	service.config = NewConfiguration(options)

	return service
}

func (s *Service) HandleReceive(params map[string]string) *Response {
	return new(Response)
}

func (s *Service) HandleSend(params map[string]string) *Response {
	return new(Response)
}

func (s *Service) HandleRevert(params map[string]string) *Response {
	return new(Response)
}

func (s *Service) HandleQuery(params map[string]string) *Response {
	return new(Response)
}
