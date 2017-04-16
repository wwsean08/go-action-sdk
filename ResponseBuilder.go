package go_action_sdk

import "github.com/wwsean08/go-action-sdk/api"

type ResponseBuilder interface {
	// Generates a response to the users query, no further questions are asked
	tellResponse(message string) api.RootResponse
	// Generates a response that will ask the user some sort of input.  The reason that the conversation token is a
	// pointer is simply to make life easier and skip nil checking.
	askResponse(message string, conversationToken *string, noInputPrompt []string) api.RootResponse
}

type defaultResponse struct {
	rootResponse api.RootResponse
}

// Create a response builder object used to respond to the request
func NewResponseBuilder() ResponseBuilder {
	rootResponse := api.RootResponse{}
	return defaultResponse{rootResponse: rootResponse}
}

func (r defaultResponse) tellResponse(message string) api.RootResponse {
	rootr := r.rootResponse
	rootr.ExpectUserResponse = false
	fResponse := api.FinalResponse{}
	sResponse := api.SpeechResponse{TextToSpeech: message, SSML: nil}
	fResponse.SpeechResponse_ = sResponse
	rootr.FinalResponse_ = fResponse

	return rootr
}

func (r defaultResponse) askResponse(message string, conversationToken *string, noInputPrompt []string) api.RootResponse {
	rootr := r.rootResponse
	rootr.ExpectUserResponse = true
	rootr.ConversationToken = conversationToken

	return rootr
}
