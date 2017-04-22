package go_action_sdk

import "github.com/wwsean08/go-action-sdk/api"

type MessageMode uint8

// Useful for custom response handlers to distinguish between response modes and generate a response in a generic way
const (
	TEXT_MODE MessageMode = iota
	SSML_MODE
)

type ResponseBuilder interface {
	// Generates a new response object which ends the discussion, not asking the user for any sort of prompt.
	TellResponse(message string, conversationToken *string) api.RootResponse
	// Generates a new response object which ends the discussion using SSML in the final response, not asking the user for any sort of prompt.
	TellResponseSSML(message string, conversationToken *string) api.RootResponse
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

func (r defaultResponse) TellResponseSSML(message string, conversationToken *string) api.RootResponse {
	return r.tellResponse(message, conversationToken, SSML_MODE)
}

func (r defaultResponse) TellResponse(message string, conversationToken *string) api.RootResponse {
	return r.tellResponse(message, conversationToken, TEXT_MODE)
}

func (r defaultResponse) tellResponse(message string, conversationToken *string, mode MessageMode) api.RootResponse {
	rootr := r.rootResponse
	rootr.ExpectUserResponse = false
	fResponse := api.FinalResponse{}
	sResponse := api.SpeechResponse{TextToSpeech: &message, SSML: nil}
	if mode == SSML_MODE {
		//Default to text and swap over to SSML if that's what the user wants to do
		sResponse.SSML = &message
		sResponse.TextToSpeech = nil
	}
	fResponse.SpeechResponse_ = sResponse
	rootr.FinalResponse_ = fResponse
	rootr.ConversationToken = conversationToken

	return rootr
}

func (r defaultResponse) AskResponseSSML(message string, conversationToken *string, noInputPrompt []string) api.RootResponse {
	return r.askResponse(message, conversationToken, noInputPrompt, SSML_MODE)
}

func (r defaultResponse) AskResponse(message string, conversationToken *string, noInputPrompt []string) api.RootResponse {
	return r.askResponse(message, conversationToken, noInputPrompt, TEXT_MODE)
}

func (r defaultResponse) askResponse(message string, conversationToken *string, noInputPrompt []string, mode MessageMode) api.RootResponse {
	rootr := r.rootResponse
	rootr.ExpectUserResponse = true
	// if conversationToken is a blank string it'll still get omitted at the json
	// serialization layer which is what we want
	rootr.ConversationToken = conversationToken
	eInputs := api.ExpectedInput{}
	iPrompt := api.InputPrompt{}
	sResponse := api.SpeechResponse{TextToSpeech: &message, SSML: nil}
	if mode == SSML_MODE {
		//Default to text and swap over to SSML if that's what the user wants to do
		sResponse.TextToSpeech = nil
		sResponse.SSML = &message
	}
	sResponseSlice := make([]api.SpeechResponse, 1)
	sResponseSlice[0] = sResponse
	iPrompt.InitialPrompts = sResponseSlice

	if len(noInputPrompt) > 0 {
		noInPrompts := make([]api.SpeechResponse, len(noInputPrompt))
		for index, element := range noInputPrompt {
			// Set the speech response correctly the first time as this could be long
			if mode == TEXT_MODE {
				noInPrompts[index] = api.SpeechResponse{TextToSpeech: &element, SSML: nil}
			} else {
				noInPrompts[index] = api.SpeechResponse{TextToSpeech: nil, SSML: &element}
			}
		}
		iPrompt.NoInputPrompts = noInPrompts
	}
	eInputs.InputPrompt_ = iPrompt

	eIntent := api.ExpectedIntent{Intent: api.TEXT_INTENT}
	// TODO: Change this limit, currently the API only allows one response per https://developers.google.com/actions/reference/conversation
	eIntentSlice := make([]api.ExpectedIntent, 1)
	eIntentSlice[0] = eIntent
	eInputs.PossibleIntents = eIntentSlice

	return rootr
}
