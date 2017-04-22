package api

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewDefaultActionConfigProvidesTheProperInformation(t *testing.T) {
	config := NewDefaultActionConfig()
	assert.Equal(t, config.ErrorMessage, "Sorry, I am unable to process your request.")
	assert.Equal(t, config.ConversationAPIVersion, "v1")
	assert.Equal(t, config.ConversationAPIVersionHeader, "Google-Assistant-API-Version")
	assert.Equal(t, config.ResponseHeaders["Google-Assistant-API-Version"], "v1")
}
