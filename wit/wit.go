package wit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Exported variables
var (
	ErrInvalidToken   = fmt.Errorf("Please use valid token")
	ErrInvalidVerbose = fmt.Errorf("WIT_API_VERBOSE must be set to true or false")
)

// ErrInvalidVersion error
type ErrInvalidVersion struct {
	Version string
}

func (e *ErrInvalidVersion) Error() string {
	return "Version %s does not exist " + e.Version
}

// Wit struct is used to configure client
type Wit struct {
	Token   string
	URL     string
	Version string
	Verbose string
}

// New function creates default instance of Wit struct
func New() (*Wit, error) {
	DefaultWit := &Wit{
		defaults("WIT_API_TOKEN", ""),
		defaults("WIT_API_URL", "https://api.wit.ai"),
		defaults("WIT_API_VERSION", "20180527"),
		defaults("WIT_API_VERBOSE", "false"),
	}
	err := DefaultWit.SanityCheck()
	return DefaultWit, err
}

// Message function creates new message struct
func (w *Wit) Message(msg string) *Message {
	urlMessage := &url.URL{Path: msg}
	params := "/message?verbose=" + w.Verbose + "&q=" + urlMessage.String()
	return &Message{w, params}
}

// SanityCheck function checks if Wit struct attributes are valid
func (w *Wit) SanityCheck() error {
	fn := [](func() error){w.validateToken, w.validateVersion, w.validateVerbose}
	for _, f := range fn {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateRequest function creates http.Client and populates headers with default values
func (w *Wit) CreateRequest(method, params string) (*http.Client, *http.Request, error) {
	err := w.SanityCheck()
	if err != nil {
		return nil, nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, w.URL+params, nil)
	req.Header.Add("Authorization", "Bearer "+w.Token)
	req.Header.Add("Accept", "application/vnd.wit."+w.Version+"+json")
	if err != nil {
		return nil, nil, err
	}
	return client, req, nil
}

// DoRequest function makes actual http call and return response body and return code
func (w *Wit) DoRequest(req *http.Request, client *http.Client) ([]byte, int, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return body, resp.StatusCode, err
}

// MakeRequest function is wrapper around DoRequest and CreateRequest functions
func (w *Wit) MakeRequest(method, params string) ([]byte, int, error) {
	client, req, err := w.CreateRequest(method, params)
	if err != nil {
		return nil, 0, err
	}
	body, code, err := w.DoRequest(req, client)
	if err != nil {
		return nil, 0, err
	}
	return body, code, nil
}

func (w *Wit) validateToken() error {
	if w.Token == "" {
		return ErrInvalidToken
	}
	return nil
}

func (w *Wit) validateVersion() error {
	t, err := time.Parse("20060102", w.Version)
	if err != nil {
		return err
	}
	if t.String() > time.Now().String() {
		return &ErrInvalidVersion{Version: w.Version}
	}
	return nil
}

func (w *Wit) validateVerbose() error {
	if w.Verbose != "false" && w.Verbose != "true" {
		return ErrInvalidVerbose
	}
	return nil
}

func defaults(parameter, fallback string) string {
	value := os.Getenv(parameter)
	if value == "" {
		value = fallback
	}
	return value
}

// TODO Also create function to populate optional arguments.
