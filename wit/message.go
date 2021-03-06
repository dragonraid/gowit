package wit

import (
	"encoding/json"
	"net/url"
)

// Message struct
type Message struct {
	WitConf *Wit
	Params  string
}

// MessageResponse struct to map response body of wit.ai/message endpoint
type MessageResponse struct {
	Text     string                 `json:"_text"`
	Entities map[string]interface{} `json:"entities"`
	MsgID    string                 `json:"msg_id"`
}

// Do function makes actual call to wit.ai/message endpoint and
// retrieves its response
func (m *Message) Do() (MessageResponse, int, error) {
	var responseBody MessageResponse
	body, statusCode, err := m.WitConf.MakeRequest("GET", m.Params)
	if err != nil {
		return responseBody, 0, err
	}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return responseBody, 0, err
	}
	return responseBody, statusCode, nil
}

// SetParameters function sets additional parameters to url message
func (m *Message) SetParameters(parameters string) {
	encodedParameters := &url.URL{Path: parameters}
	m.Params += "&" + encodedParameters.String()
}
