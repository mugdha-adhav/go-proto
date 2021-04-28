package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
	sample "github.com/mugdha-adhav/go-proto/pkg/protobufs"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func createMessageFromDefinition() {
	data := &sample.MyMessage{
		Name:        "anakin-skywalker",
		Uuid:        1,
		Description: "i am a jedi master",
	}

	fmt.Println(data)
}

func createMessageFromYAML() error {
	messageFilename := "data/sample.yaml"
	bytes, err := ioutil.ReadFile(messageFilename)
	if err != nil {
		return err
	}

	messageType, _ := protoregistry.GlobalTypes.FindMessageByURL("mugdha.messages.sample.myMessage")
	data := messageType.New().Interface().(proto.Message)

	if err := yaml.Unmarshal(bytes, &data); err != nil {
		return err
	}

	fmt.Println(data)

	return nil
}

func main() {
	createMessageFromDefinition()
	if err := createMessageFromYAML(); err != nil {
		fmt.Println(err.Error())
	}
}
