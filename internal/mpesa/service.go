package mpesa


type Service struct {
	config string
}

func NewService(config map[string]interface{}) *Service {
	return nil
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
