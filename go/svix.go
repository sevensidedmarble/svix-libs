package svix

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/svix/svix-libs/go/internal/openapi"
	"github.com/svix/svix-libs/go/internal/version"
)

type (
	SvixOptions struct {
		Debug bool

		// Overrides the base URL (protocol + hostname) used for all requests sent by this Svix client. (Useful for testing)
		DebugURL   *url.URL
		HTTPClient *http.Client
	}
	Svix struct {
		Authentication *Authentication
		Application    *Application
		Endpoint       *Endpoint
		EventType      *EventType
		Message        *Message
		MessageAttempt *MessageAttempt
	}
)

var defaultHTTPClient = &http.Client{
	Timeout: 60 * time.Second,
}

func String(s string) *string {
	return &s
}
func Int32(i int32) *int32 {
	return &i
}

func New(token string, options *SvixOptions) *Svix {
	conf := openapi.NewConfiguration()
	conf.Scheme = "https"
	conf.Host = "api.svix.com"
	conf.HTTPClient = defaultHTTPClient
	if options != nil {
		conf.Debug = options.Debug
		if options.DebugURL != nil {
			conf.Scheme = options.DebugURL.Scheme
			conf.Host = options.DebugURL.Host
		}
		if options.HTTPClient != nil {
			conf.HTTPClient = options.HTTPClient
		}
	}
	conf.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	conf.UserAgent = fmt.Sprintf("svix-libs/%s/go", version.Version)
	apiClient := openapi.NewAPIClient(conf)
	return &Svix{
		Authentication: &Authentication{
			api: apiClient,
		},
		Application: &Application{
			api: apiClient,
		},
		Endpoint: &Endpoint{
			api: apiClient,
		},
		EventType: &EventType{
			api: apiClient,
		},
		Message: &Message{
			api: apiClient,
		},
		MessageAttempt: &MessageAttempt{
			api: apiClient,
		},
	}
}
