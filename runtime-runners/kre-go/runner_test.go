package kre

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
)

type Input struct {
	Name string `json:"name"`
}

type Output struct {
	Greeting string `json:"greeting"`
}

func handlerInit(ctx *HandlerContext) {
	log.Println("[worker init]")
	ctx.SetValue("greeting", "Hello")
}

func handler(ctx *HandlerContext, data []byte) (interface{}, error) {
	log.Println("[worker handler]")

	input := Input{}
	err := json.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}

	greetingText := fmt.Sprintf("%s %s!", ctx.GetValue("greeting"), input.Name)
	log.Println(greetingText)

	// Saves metrics in MongoDB DB sending a message to the MongoWriter queue
	//ctx.SaveMetric(time.Now(), "class_x", "class_y")
	//ctx.SaveMetricError(ErrNewLabels)

	out := Output{}
	out.Greeting = greetingText
	return out, nil // Must be a serializable JSON
}

func setEnvVars(t *testing.T, envVars map[string]string) {
	for name, value := range envVars {
		err := os.Setenv(name, value)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestRunner(t *testing.T) {
	const inputSubject = "test-subject-input"
	setEnvVars(t, map[string]string{
		"KRT_VERSION":           "testVersion1",
		"KRT_NODE_NAME":         "nodeA",
		"KRT_BASE_PATH":         "./test",
		"KRT_NATS_SERVER":       "localhost:4222",
		"KRT_NATS_INPUT":        inputSubject,
		"KRT_NATS_OUTPUT":       "",
		"KRT_NATS_MONGO_WRITER": "mongo_writer",
	})

	go Start(handlerInit, handler)

	cfg := NewConfig()
	nc, err := nats.Connect(cfg.NATS.Server)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	input := Input{Name: "John"}
	inputData, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}

	msg, err := json.Marshal(Result{
		Reply: "",
		Data:  string(inputData),
	})
	if err != nil {
		t.Fatal(err)
	}
	res, err := nc.Request(inputSubject, msg, 1*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	expectedRes := "Hello John!"
	result := &Result{}
	err = json.Unmarshal(res.Data, result)
	if err != nil {
		t.Fatal(err)
	}

	resultData := &Output{}
	err = json.Unmarshal([]byte(result.Data), resultData)
	if err != nil {
		t.Fatal(err)
	}

	if expectedRes != resultData.Greeting {
		t.Fatalf("The greeting '%s' must be '%s'", resultData.Greeting, expectedRes)
	}
}
