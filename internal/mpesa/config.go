package mpesa

type Configuration struct {
	Timeout uint
	ApiKey string
	PublicKey string
	AccessToken string
	VerifySSL bool
	Debug bool
	Endpoint string
	
	envoronment string
}