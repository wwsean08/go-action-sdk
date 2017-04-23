package go_action_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultResponse_TellResponseReturnsExpectedOutput(t *testing.T) {
	builder := NewResponseBuilder()
	cToken := "42"
	message := "This Is A Test"
	response := builder.TellResponse(message, &cToken)
	assert.False(t, response.ExpectUserResponse)
	assert.Nil(t, response.ExpectedInputs)
	assert.Equal(t, &cToken, response.ConversationToken)
	assert.Equal(t, &message, response.FinalResponse_.SpeechResponse_.TextToSpeech)
	assert.Nil(t, response.FinalResponse_.SpeechResponse_.SSML)
}

func TestDefaultResponse_TellResponseSSMLReturnsExpectedOutput(t *testing.T) {
	builder := NewResponseBuilder()
	cToken := "42"
	message := "<s>This Is A Test</s>"
	response := builder.TellResponseSSML(message, &cToken)
	assert.False(t, response.ExpectUserResponse)
	assert.Nil(t, response.ExpectedInputs)
	assert.Equal(t, &cToken, response.ConversationToken)
	assert.Equal(t, &message, response.FinalResponse_.SpeechResponse_.SSML)
	assert.Nil(t, response.FinalResponse_.SpeechResponse_.TextToSpeech)
}

func TestDefaultResponse_AskResponseReturnsExpectedOutput(t *testing.T) {
	// nil noInputPrompt check
	builder := NewResponseBuilder()
	cToken := "42"
	message := "This Is A Test"
	response := builder.AskResponse(message, &cToken, nil)
	assert.True(t, response.ExpectUserResponse)
	assert.Nil(t, response.FinalResponse_)
	assert.NotNil(t, response.ExpectedInputs)
	assert.Len(t, response.ExpectedInputs, 1)
	assert.Len(t, response.ExpectedInputs[0].PossibleIntents, 1)
	assert.Equal(t, "assistant.intent.action.TEXT", response.ExpectedInputs[0].PossibleIntents[0].Intent)
	assert.Nil(t, response.ExpectedInputs[0].PossibleIntents[0].InputValueSpec_)
	assert.Len(t, response.ExpectedInputs[0].InputPrompt_.InitialPrompts, 1)
	assert.Nil(t, response.ExpectedInputs[0].InputPrompt_.InitialPrompts[0].SSML)
	assert.Equal(t, &message, response.ExpectedInputs[0].InputPrompt_.InitialPrompts[0].TextToSpeech)
	assert.Nil(t, response.ExpectedInputs[0].InputPrompt_.NoInputPrompts)

	// actual noInputPrompt check
	cToken = "42"
	message = "This Is A Test"
	noInputPromptStr := "Nevermind"
	noInputPromptSlice := make([]string, 1)
	noInputPromptSlice[0] = noInputPromptStr
	response = builder.AskResponse(message, &cToken, noInputPromptSlice)
	assert.True(t, response.ExpectUserResponse)
	assert.Nil(t, response.FinalResponse_)
	assert.NotNil(t, response.ExpectedInputs)
	assert.Len(t, response.ExpectedInputs, 1)
	assert.Len(t, response.ExpectedInputs[0].PossibleIntents, 1)
	assert.Equal(t, "assistant.intent.action.TEXT", response.ExpectedInputs[0].PossibleIntents[0].Intent)
	assert.Nil(t, response.ExpectedInputs[0].PossibleIntents[0].InputValueSpec_)
	assert.Len(t, response.ExpectedInputs[0].InputPrompt_.InitialPrompts, 1)
	assert.Nil(t, response.ExpectedInputs[0].InputPrompt_.InitialPrompts[0].SSML)
	assert.Equal(t, &message, response.ExpectedInputs[0].InputPrompt_.InitialPrompts[0].TextToSpeech)
	assert.Len(t, response.ExpectedInputs[0].InputPrompt_.NoInputPrompts, 1)
	assert.Equal(t, &noInputPromptStr, response.ExpectedInputs[0].InputPrompt_.NoInputPrompts[0].TextToSpeech)
	assert.Nil(t, response.ExpectedInputs[0].InputPrompt_.NoInputPrompts[0].SSML)
}

func TestDefaultResponse_AskResponseSSMLReturnsExpectedOutput(t *testing.T) {
	// nil noInputPrompt check
	builder := NewResponseBuilder()
	cToken := "42"
	message := "This Is A Test"
	response := builder.AskResponseSSML(message, &cToken, nil)
	assert.True(t, response.ExpectUserResponse)
	assert.Nil(t, response.FinalResponse_)
	assert.NotNil(t, response.ExpectedInputs)
	assert.Len(t, response.ExpectedInputs, 1)
	assert.Len(t, response.ExpectedInputs[0].PossibleIntents, 1)
	assert.Equal(t, "assistant.intent.action.TEXT", response.ExpectedInputs[0].PossibleIntents[0].Intent)
	assert.Nil(t, response.ExpectedInputs[0].PossibleIntents[0].InputValueSpec_)
	assert.Len(t, response.ExpectedInputs[0].InputPrompt_.InitialPrompts, 1)
	assert.Nil(t, response.ExpectedInputs[0].InputPrompt_.InitialPrompts[0].TextToSpeech)
	assert.Equal(t, &message, response.ExpectedInputs[0].InputPrompt_.InitialPrompts[0].SSML)
	assert.Nil(t, response.ExpectedInputs[0].InputPrompt_.NoInputPrompts)

	// actual noInputPrompt check
	cToken = "42"
	message = "This Is A Test"
	noInputPromptStr := "Nevermind"
	noInputPromptSlice := make([]string, 1)
	noInputPromptSlice[0] = noInputPromptStr
	response = builder.AskResponseSSML(message, &cToken, noInputPromptSlice)
	assert.True(t, response.ExpectUserResponse)
	assert.Nil(t, response.FinalResponse_)
	assert.NotNil(t, response.ExpectedInputs)
	assert.Len(t, response.ExpectedInputs, 1)
	assert.Len(t, response.ExpectedInputs[0].PossibleIntents, 1)
	assert.Equal(t, "assistant.intent.action.TEXT", response.ExpectedInputs[0].PossibleIntents[0].Intent)
	assert.Nil(t, response.ExpectedInputs[0].PossibleIntents[0].InputValueSpec_)
	assert.Len(t, response.ExpectedInputs[0].InputPrompt_.InitialPrompts, 1)
	assert.Nil(t, response.ExpectedInputs[0].InputPrompt_.InitialPrompts[0].TextToSpeech)
	assert.Equal(t, &message, response.ExpectedInputs[0].InputPrompt_.InitialPrompts[0].SSML)
	assert.Len(t, response.ExpectedInputs[0].InputPrompt_.NoInputPrompts, 1)
	assert.Equal(t, &noInputPromptStr, response.ExpectedInputs[0].InputPrompt_.NoInputPrompts[0].SSML)
	assert.Nil(t, response.ExpectedInputs[0].InputPrompt_.NoInputPrompts[0].TextToSpeech)
}
