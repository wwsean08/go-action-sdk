package sample

import (
	"github.com/labstack/echo"
	"github.com/wwsean08/go-action-sdk"
	"github.com/wwsean08/go-action-sdk/api"
	"net/http"
)

var router = echo.New()

func init() {
	router.POST("/", indexHandler)
}

func indexHandler(c echo.Context) error {
	rootRequest := api.RootRequest{}
	err := c.Bind(&rootRequest)
	if err != nil {
		print(err)
		c.Error(err)
		return err
	}

	//Add our sdk header
	actionConfig := api.NewDefaultActionConfig()
	c.Response().Header().Add(actionConfig.ConversationAPIVersionHeader, actionConfig.ConversationAPIVersion)

	if rootRequest.Inputs == nil || *rootRequest.Inputs[0].Intent == api.MAIN_INTENT {
		return mainHandler(c)
	}
	return textHandler(c, rootRequest)
}
func textHandler(c echo.Context, request api.RootRequest) error {
	builder := go_action_sdk.NewResponseBuilder()
	query := request.Inputs[0].RawInputs[0].Query
	response, err := builder.TellResponse("You just said: "+*query, nil)
	if err != nil {
		c.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, response)
}

func mainHandler(c echo.Context) error {
	builder := go_action_sdk.NewResponseBuilder()
	response, err := builder.AskResponse("How's it going, please tell me something so I can repeat it back!", nil, nil)
	if err != nil {
		c.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, response)
}
