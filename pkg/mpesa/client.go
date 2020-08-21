package mpesa

import (
	mpesa "github.com/edsonmichaque/mpesa-go-sdk/internal/mpesa"
)

type Client struct {
	service mpesa.Service
}

func New(config map[string]interface{}) *Client {
	client := new(Client)
	return client
}

func NewClient(config map[string]interface{}) *Client {
	return New(config)
}

func (c *Client) Receive(params map[string]string)  *mpesa.Response {
	return c.service.HandleReceive(params)
}

func (c *Client) Send(params map[string]string) *mpesa.Response {
	return c.service.HandleSend(params)
}

func (c *Client) Revert(params map[string]string) *mpesa.Response {
	return c.service.HandleRevert(params)
}

func (c *Client) Query(params map[string]string) *mpesa.Response {
	return c.service.HandleQuery(params)
}
