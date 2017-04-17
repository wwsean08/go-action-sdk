package go_action_sdk

import "github.com/wwsean08/go-action-sdk/api"

type ResponseBuilder interface {
	// Generates a response to the users query, no further questions are asked
	TellResponse(message string) api.RootResponse
	// Generates a response that will ask the user some sort of input.
	AskResponse(message string, conversationToken string, noInputPrompt []string) api.RootResponse
}

type defaultResponse struct {
	rootResponse api.RootResponse
}

// Create a response builder object used to respond to the request
func NewResponseBuilder() ResponseBuilder {
	rootResponse := api.RootResponse{}
	return defaultResponse{rootResponse: rootResponse}
}

func (r defaultResponse) TellResponse(message string) api.RootResponse {
	rootr := r.rootResponse
	rootr.ExpectUserResponse = false
	fResponse := api.FinalResponse{}
	sResponse := api.SpeechResponse{TextToSpeech: message, SSML: ""}
	fResponse.SpeechResponse_ = sResponse
	rootr.FinalResponse_ = fResponse

	return rootr
}

func (r defaultResponse) AskResponse(message string, conversationToken string, noInputPrompt []string) api.RootResponse {
	rootr := r.rootResponse
	rootr.ExpectUserResponse = true
	// if conversationToken is a blank string it'll still get omitted at the json
	// serialization layer which is what we want
	rootr.ConversationToken = conversationToken
	eInputs := api.ExpectedInput{}
	iPrompt := api.InputPrompt{}
	sResponse := api.SpeechResponse{TextToSpeech: message, SSML: ""}
	sResponseSlice := make([]api.SpeechResponse, 1)
	sResponseSlice[0] = sResponse
	iPrompt.InitialPrompts = sResponseSlice

	if len(noInputPrompt) > 0 {
		noInPrompts := make([]api.SpeechResponse, len(noInputPrompt))
		for index, element := range noInputPrompt {
			noInPrompts[index] = api.SpeechResponse{TextToSpeech: element, SSML: ""}
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
