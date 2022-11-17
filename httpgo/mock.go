package httpgo

import (
	"fmt"
	"net/http"
)

type Mock struct {
	Method        string
	Url           string
	ReqBody       string
	ResBody       string
	Headers       string
	Error         error
	ResStatusCode int
}

func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := Response{
		status:     fmt.Sprintf("%d, %s", m.ResStatusCode, http.StatusText(m.ResStatusCode)), // interesante
		statusCode: m.ResStatusCode,
		body:       []byte(m.ResBody),
	}

	return &response, nil
}
