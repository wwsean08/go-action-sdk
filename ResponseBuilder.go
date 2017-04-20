package go_action_sdk

import "github.com/wwsean08/go-action-sdk/api"

type ResponseBuilder interface {
	// Generates a new response object which ends the discussion, not asking the user for any sort of prompt.
	TellResponse(message string, conversationToken *string) api.RootResponse
	// Generates a response object which asks the user to respond.  You can also provide a list of strings that
	// can be said to not provide a response.  A conversation token is key in making sure you do not lose the context of
	// the conversation, make sure this is either the same one that is sent from the user or if this is a new request a
	// unique one as any response that comes back will contain it and can be used for correlation.
	AskResponse(message string, conversationToken *string, noInputPrompt []string) api.RootResponse
}

type defaultResponse struct {
	rootResponse api.RootResponse
}

// Create a response builder object used to respond to the request
func NewResponseBuilder() ResponseBuilder {
	rootResponse := api.RootResponse{}
	return defaultResponse{rootResponse: rootResponse}
}

func (r defaultResponse) TellResponse(message string, conversationToken *string) api.RootResponse {
	rootr := r.rootResponse
	rootr.ExpectUserResponse = false
	fResponse := api.FinalResponse{}
	sResponse := api.SpeechResponse{TextToSpeech: &message, SSML: nil}
	fResponse.SpeechResponse_ = sResponse
	rootr.FinalResponse_ = fResponse
	rootr.ConversationToken = conversationToken

	return rootr
}

func (r defaultResponse) AskResponse(message string, conversationToken *string, noInputPrompt []string) api.RootResponse {
	rootr := r.rootResponse
	rootr.ExpectUserResponse = true
	// if conversationToken is a blank string it'll still get omitted at the json
	// serialization layer which is what we want
	rootr.ConversationToken = conversationToken
	eInputs := api.ExpectedInput{}
	iPrompt := api.InputPrompt{}
	sResponse := api.SpeechResponse{TextToSpeech: &message, SSML: nil}
	sResponseSlice := make([]api.SpeechResponse, 1)
	sResponseSlice[0] = sResponse
	iPrompt.InitialPrompts = sResponseSlice

	if len(noInputPrompt) > 0 {
		noInPrompts := make([]api.SpeechResponse, len(noInputPrompt))
		for index, element := range noInputPrompt {
			noInPrompts[index] = api.SpeechResponse{TextToSpeech: &element, SSML: nil}
		}
		iPrompt.NoInputPrompts = noInPrompts
	}
	eInputs.InputPrompt_ = iPrompt

	eIntent := api.ExpectedIntent{Intent: api.TEXT_INTENT}
	eIntentSlice := make([]api.ExpectedIntent, 1)
	eIntentSlice[0] = eIntent
	eInputs.PossibleIntents = eIntentSlice

	return rootr
}
