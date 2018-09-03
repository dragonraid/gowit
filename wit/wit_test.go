package wit_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/dragonraid/gowit/wit"
)

var config = []struct {
	Input       *wit.Wit
	Expected    error
	Description string
}{
	{
		Input: &wit.Wit{
			Token:   "WitAPItoken",
			URL:     "https://api.wit.ai",
			Version: "20180711",
			Verbose: "false",
		},
		Expected:    nil,
		Description: "Correct configuration",
	},
	{
		Input: &wit.Wit{
			Token:   "",
			URL:     "https://api.wit.ai",
			Version: "20180711",
			Verbose: "false",
		},
		Expected:    wit.ErrInvalidToken,
		Description: "Invalid API token",
	},
	{
		Input: &wit.Wit{
			Token:   "WitAPItoken",
			URL:     "https://api.wit.ai",
			Version: "2018071111",
			Verbose: "false",
		},
		Expected: &time.ParseError{
			Layout:     "20060102",
			Value:      "2018071111",
			LayoutElem: "",
			ValueElem:  "11",
			Message:    ": extra text: " + "11"},
		Description: "Invalid version format",
	},
	{
		Input: &wit.Wit{
			Token:   "WitAPItoken",
			URL:     "https://api.wit.ai",
			Version: "30180111",
			Verbose: "false",
		},
		Expected:    &wit.ErrInvalidVersion{Version: "30180111"},
		Description: "Invalid version",
	},
	{
		Input: &wit.Wit{
			Token:   "WitAPItoken",
			URL:     "https://api.wit.ai",
			Version: "20180711",
			Verbose: "true",
		},
		Expected:    nil,
		Description: "Verbose true",
	},
	{
		Input: &wit.Wit{
			Token:   "WitAPItoken",
			URL:     "https://api.wit.ai",
			Version: "20180711",
			Verbose: "",
		},
		Expected:    wit.ErrInvalidVerbose,
		Description: "Invalid Verbose parameter",
	},
}

// RequestHandler type
type RequestHandler func(w http.ResponseWriter, r *http.Request)

// TestWit f
func TestWit(t *testing.T) {

	for _, conf := range config {
		t.Run(conf.Description, func(t *testing.T) {
			if err := conf.Input.SanityCheck(); reflect.TypeOf(err) == reflect.TypeOf(conf.Expected) {
				if err == nil {
				} else if err.Error() != conf.Expected.Error() {
					t.Errorf("Expected \"%s\", but got \"%s\"\n",
						conf.Expected.Error(), err.Error())
				}
			} else {
				t.Errorf("Expected \"%T\", but got \"%T\"\n",
					conf.Expected, err)
			}
		})
	}
}

// WitTestRequest f
func WitTestRequest(method, params string, fn RequestHandler) (*httptest.ResponseRecorder, error) {
	_, request, err := config[0].Input.CreateRequest(method, params)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(fn)
	handler.ServeHTTP(responseRecorder, request)
	return responseRecorder, err
}
