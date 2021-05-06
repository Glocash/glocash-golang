package glocash

import "time"

type Client interface {
	PayEmbed(data map[string]string) string
	PayClassic(data map[string]string) *RespPayment
	PayDirect(data map[string]string) *RespDirect
	Refund(data map[string]string) *RespRefund
	Query(data map[string]string) *RespQueryList
}

// implement client
type Glocash struct {
	Environ      string
	Email        string
	MerchantName string
	Key          string
	JsKey        string
	Domain       string
	Timeout      time.Duration
	Scheme       string
	RequestParam map[string]string
	RespJson     string
	url          string
}
