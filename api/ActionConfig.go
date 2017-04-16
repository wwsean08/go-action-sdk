// The basic API for interacting with a google home/google assistant with structures to
// convert the request/response to and from json
package api

type ActionConfig struct {
	ErrorMessage                 string
	ConversationAPIVersionHeader string
	ConversationAPIVersion       string
	ResponseHeaders              map[string]string
}

// NewDefaultActionConfig generates an ActionConfig with the default values and returns it
func NewDefaultActionConfig() ActionConfig {
	headers := make(map[string]string)
	ac := ActionConfig{
		ErrorMessage:                 "Sorry, I am unable to process your request.",
		ConversationAPIVersion:       "v1",
		ConversationAPIVersionHeader: "Google-Assistant-API-Version",
	}
	headers[ac.ConversationAPIVersionHeader] = ac.ConversationAPIVersion
	ac.ResponseHeaders = headers
	return ac
}
