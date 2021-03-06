package main

import (
	"encoding/json"
	"fmt"

	"github.com/konstellation-io/kre/runners/kre-go"
)

type Input struct {
	Name string `json:"name"`
}

type Output struct {
	Greeting string `json:"greeting"`
}

func handlerInit(ctx *kre.HandlerContext) {
	ctx.Logger.Info("[worker init]")
	ctx.Set("greeting", "Hello")
}

func handler(ctx *kre.HandlerContext, data []byte) (interface{}, error) {
	ctx.Logger.Info("[worker handler]")

	input := Input{}
	err := json.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}

	greetingText := fmt.Sprintf("%s %s!", ctx.Get("greeting"), input.Name)
	ctx.Logger.Info(greetingText)

	out := Output{}
	out.Greeting = greetingText

	// Saving salutation into DB for later analysis or use
	err = ctx.DB.Save("go-greeter", out)
	if err != nil {
		return nil, err
	}

	return out, nil // Must be a serializable JSON
}

func main() {
	kre.Start(handlerInit, handler)
}
