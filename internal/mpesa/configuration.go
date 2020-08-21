package mpesa

type Configuration struct {
	Timeout uint
	ApiKey string
	PublicKey string
	AccessToken string
	VerifySSL bool
	Debug bool
	Endpoint string
	Origin string
	UserAgent string
	
	Environment Environment
}

func NewConfiguration(options map[string]interface{}) *Configuration {
	config := new(Configuration)
	config.initDefault()

	for k, v := range(options) {
		switch k {
		case "apiKey": config.ApiKey = v.(string)
		case "publicKey": config.PublicKey = v.(string)
		case "timeout": config.Timeout = v.(uint)
		case "userAgent": config.UserAgent = v.(string)
		case "accessToken": config.AccessToken = v.(string)
		case "origin" : config.Origin = v.(string)
		case "environment": config.Environment = v.(Environment)
		case "endpoint": config.Endpoint = v.(string)
		}
	}

	if config.Endpoint == "" {
		if config.Environment == Production {
			config.Endpoint = "https://api.sandbox.vm.co.mz"
		} else {
			config.Endpoint = "https://api.mpesa.vm.co.mz"
		}
	}

	return config
}

func (c *Configuration) initDefault() {
	c.Timeout = 0
	c.Origin = "*"
	c.VerifySSL = false
	c.Debug = true
	c.Environment = Sandbox
	c.UserAgent = "M-Pesa/Go"
}
