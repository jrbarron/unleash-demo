package main

import (
	"github.com/Unleash/unleash-client-go/v3/context"
	"regexp"
)

type UserAgentStrategy struct{}

func (s UserAgentStrategy) Name() string {
	return "UserAgentStrategy"
}

func (s UserAgentStrategy) IsEnabled(
	params map[string]interface{},
	ctx *context.Context,
) bool {

	if ctx == nil {
		return false
	}
	value, found := params["userAgent"]
	if !found {
		return false
	}

	userAgent, ok := value.(string)
	if !ok {
		return false
	}

	re, err := regexp.Compile(userAgent)
	if err != nil {
		return false
	}

	return re.MatchString(ctx.Properties["userAgent"])
}
