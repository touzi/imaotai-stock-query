package utils

import "fmt"

const (
	API      = "http://www.pushplus.plus/send"
	TOKEN    = "37baed1ffb67403c8a0ae49c836b0504"
	TOPIC    = "maotai"
	TEMPLATE = "markdown"
)

type Push struct{}

func (p *Push) Send(content string) (err error) {
	httpClient := NewHttpClient()
	err, res := httpClient.Get(API, map[string]interface{}{
		"token":    TOKEN,
		"topic":    TOPIC,
		"template": TEMPLATE,
		"content":  content,
	})
	if err != nil {
		return
	}
	fmt.Println(string(res))
	return
}

func New() *Push {
	return &Push{}
}
