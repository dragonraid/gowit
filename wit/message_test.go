package wit_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/dragonraid/gowit/wit"
)

// MessageTestStruct struct
type MessageTestStruct struct {
	Input       *wit.Message
	Response    string
	Expected    []error
	Description string
}

var message = []MessageTestStruct{
	{
		Input: config[0].Input.Message("Message test response"),
		Response: `{
                              "_text": "Message test response", 
                              "entities": {
                                "datetime": [
                                  {
                                    "confidence": 0.9,
                                    "value": "test"
                                  }
                                ],
                                "intent": "date"
                              },
                              "msg_id": "1a2b3c"
                           }`,
		Expected:    []error{nil, nil},
		Description: "Correct response",
	},
	{
		Input: config[0].Input.Message("Message test response"),
		Response: `{
                              "_text": "Message test response", 
                              "entities": {
                                  9:false,
                              },
                              "msg_id": "1a2b3c"
                            }`,
		Expected:    []error{nil, fmt.Errorf("invalid character '9' looking for beginning of object key string")},
		Description: "Incorrect response",
	},
}

// MessageHandler function mocks wit.ai/messages endpoint
func (m *MessageTestStruct) MessageHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, m.Response)
	//w.Write(payload)
}

// TestMessageHandler function uses MessageHandler to perform specific usecase tests
func TestMessageHandler(t *testing.T) {
	for _, msg := range message {
		t.Run(msg.Description, func(t *testing.T) {
			responseRecorder, err := WitTestRequest("GET", msg.Input.Params, msg.MessageHandler)
			if err != msg.Expected[0] {
				t.Errorf("An error occured: \"%s\"", err)
			}
			var responseBody wit.MessageResponse
			err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody)
			if err == nil {
			} else if err.Error() != msg.Expected[1].Error() {
				t.Errorf("An error occured: \"%s\"", err)
			}
		})
	}

}
