package api

// RootRequest object which should contain all the request data sent via google
type RootRequest struct {
	User_         *User         `json:"user"`
	Device_       *Device       `json:"device"`
	Conversation_ *Conversation `json:"conversation"`
	Inputs        []Input       `json:"inputs"`
}

type Conversation struct {
	ConversationID    *string `json:"conversation_id"`
	ConversationToken *string `json:"conversation_token"`
	Type              *string `json:"type"`
	//Type Was an enum in the java version, may need to be changed
}

type Time struct {
	Seconds int `json:"seconds"`
	Nanos   int `json:"nanos"`
}

type Input struct {
	Intent    *string    `json:"intent"`
	RawInputs []RawInput `json:"raw_inputs"`
	Arguments []Argument `json:"arguments"`
}

type Argument struct {
	Name             *string   `json:"name"`
	RawText          *string   `json:"raw_text"`
	IntValue         *string   `json:"int_value"`
	BoolValue        *string   `json:"bool_value"`
	TextValue        *string   `json:"text_value"`
	DateValue        *string   `json:"date_value"`
	TimeValue        *string   `json:"time_value"`
	LocationValue    *Location `json:"location_value"`
	FormattedAddress *string   `json:"formatted_address"`
}

type RawInput struct {
	CreateTime *Time   `json:"create_time"`
	Query      *string `json:"query"`
	InputType  *string `json:"input_type"`
	//InputType was an enuim in java version, may need to be changed
}

type Device struct {
	Location_ *Location `json:"location"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Location struct {
	City             *string      `json:"city"`
	ZipCode          *string      `json:"zip_code"`
	FormattedAddress *string      `json:"formatted_address"`
	Coordinates_     *Coordinates `json:"coordinates"`
}

type UserProfile struct {
	GivenName   *string `json:"given_name"`
	FamilyName  *string `json:"family_name"`
	DisplayName *string `json:"display_name"`
}

type User struct {
	UserId      *string      `json:"user_id"`
	AccessToken *string      `json:"access_token"`
	Profile     *UserProfile `json:"profile"`
}
