package api

// RootResponse object which will be formatted to json and sent back to google and onto the user.
type RootResponse struct {
	ConversationToken  *string         `json:"conversation_token,omitempty"`
	ExpectUserResponse bool            `json:"expect_user_response"`
	ExpectedInputs     []ExpectedInput `json:"expected_inputs,omitempty"`
	FinalResponse_     *FinalResponse  `json:"final_response,omitempty"`
}

type ExpectedInput struct {
	PossibleIntents []ExpectedIntent `json:"possible_intents"`
	InputPrompt_    InputPrompt      `json:"input_prompt"`
}

type ExpectedIntent struct {
	Intent          string          `json:"intent"`
	InputValueSpec_ *InputValueSpec `json:"input_value_spec,omitempty"`
}

type FinalResponse struct {
	SpeechResponse_ SpeechResponse `json:"speech_response"`
}

type InputValueSpec struct {
	PermissionValueSpec_ PermissionValueSpec `json:"permission_value_spec"`
}

type PermissionValueSpec struct {
	OptContext  *string  `json:"opt_context"`
	Permissions []string `json:"permissions"`
}

type SpeechResponse struct {
	TextToSpeech *string `json:"text_to_speech,omitempty"`
	SSML         *string `json:"ssml,omitempty"`
}

type InputPrompt struct {
	InitialPrompts []SpeechResponse `json:"initial_prompts"`
	NoInputPrompts []SpeechResponse `json:"no_input_prompts,omitempty"`
}
