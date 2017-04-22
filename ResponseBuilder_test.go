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
