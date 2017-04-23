package go_action_sdk

import (
	"fmt"
	"github.com/wwsean08/go-action-sdk/api"
)

type MessageMode uint8

// Useful for response builders to distinguish between response modes and generate a response in a generic way
const (
	TEXT_MODE MessageMode = iota
	SSML_MODE
)

type ResponseBuilder interface {
	// Generates a new response object which ends the discussion, not asking the user for any sort of prompt.
	TellResponse(message string, conversationToken *string) (api.RootResponse, error)
	// Generates a new response object which ends the discussion using SSML in the final response, not asking the user for any sort of prompt.
	TellResponseSSML(message string, conversationToken *string) (api.RootResponse, error)
	// Generates a response object which asks the user to respond.  You can also provide a list of strings that
	// can be said to not provide a response.  A conversation token is key in making sure you do not lose the context of
	// the conversation, make sure this is either the same one that is sent from the user or if this is a new request a
	// unique one as any response that comes back will contain it and can be used for correlation.
	AskResponse(message string, conversationToken *string, noInputPrompt []string) (api.RootResponse, error)
	// Generates a response object which asks the user to respond using SSML.  You can also provide a list of strings that
	// can be said to not provide a response.  A conversation token is key in making sure you do not lose the context of
	// the conversation, make sure this is either the same one that is sent from the user or if this is a new request a
	// unique one as any response that comes back will contain it and can be used for correlation.
	AskResponseSSML(message string, conversationToken *string, noInputPrompt []string) (api.RootResponse, error)
}

type defaultResponse struct {
}

// Create a response builder object used to respond to the request
func NewResponseBuilder() ResponseBuilder {
	return defaultResponse{}
}

func (r defaultResponse) TellResponseSSML(message string, conversationToken *string) (api.RootResponse, error) {
	return r.tellResponse(message, conversationToken, SSML_MODE)
}

func (r defaultResponse) TellResponse(message string, conversationToken *string) (api.RootResponse, error) {
	return r.tellResponse(message, conversationToken, TEXT_MODE)
}

func (r defaultResponse) tellResponse(message string, conversationToken *string, mode MessageMode) (api.RootResponse, error) {
	rootr := api.RootResponse{ExpectUserResponse: false, ConversationToken: conversationToken}

	// Validate the message is ASCII per https://developers.google.com/actions/reference/conversation#SpeechResponse
	err := r.isAscii(message)
	if err != nil {
		return rootr, err
	}

	fResponse := api.FinalResponse{}
	sResponse := api.SpeechResponse{TextToSpeech: &message, SSML: nil}
	if mode == SSML_MODE {
		//Default to text and swap over to SSML if that's what the user wants to do
		sResponse.SSML = &message
		sResponse.TextToSpeech = nil
	}
	fResponse.SpeechResponse_ = sResponse
	rootr.FinalResponse_ = &fResponse

	return rootr, nil
}

func (r defaultResponse) AskResponseSSML(message string, conversationToken *string, noInputPrompt []string) (api.RootResponse, error) {
	return r.askResponse(message, conversationToken, noInputPrompt, SSML_MODE)
}

func (r defaultResponse) AskResponse(message string, conversationToken *string, noInputPrompt []string) (api.RootResponse, error) {
	return r.askResponse(message, conversationToken, noInputPrompt, TEXT_MODE)
}

func (r defaultResponse) askResponse(message string, conversationToken *string, noInputPrompt []string, mode MessageMode) (api.RootResponse, error) {
	rootr := api.RootResponse{ExpectUserResponse: true, ConversationToken: conversationToken}

	// Validate that noInputPrompt is no more than 3 elements long per https://developers.google.com/actions/reference/conversation#InputPrompt
	if len(noInputPrompt) > 3 {
		return rootr, fmt.Errorf("TooLongError: max length for noInputPromot is 3 but recieved a slice of length %d", len(noInputPrompt))
	}
	// Validate the message is ASCII per https://developers.google.com/actions/reference/conversation#SpeechResponse
	err := r.isAscii(message)
	if err != nil {
		return rootr, err
	}

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
	eInputsSlice := make([]api.ExpectedInput, 1)
	eIntentSlice[0] = eIntent
	eInputs.PossibleIntents = eIntentSlice
	eInputsSlice[0] = eInputs
	rootr.ExpectedInputs = eInputsSlice

	return rootr, nil
}

// Checks if the entire string is ascii, if not returns an error.  This validation is in place
// as all text in SpeechResponse objects MUST be ascii per https://developers.google.com/actions/reference/conversation#SpeechResponse
func (r defaultResponse) isAscii(message string) error {
	byte_message := []byte(message)
	for _, element := range byte_message {
		if int(element) > 127 {
			return fmt.Errorf("NonASCIIIError: Unexpected character found in message %s, found character %s.", message, string(element))
		}
	}
	return nil
}
