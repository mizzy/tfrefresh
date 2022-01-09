package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Detail Detail
}

type Detail struct {
	EventName         string
	RequestParameters map[string]interface{}
	ResponseElements  map[string]interface{}
}

func main() {
	lambda.Start(tfrefresh)
}

func tfrefresh(e Event) error {
	fmt.Printf("%#v\n", e)

	req, _ := json.Marshal(e.Detail.RequestParameters)
	fmt.Println(string(req))

	res, _ := json.Marshal(e.Detail.ResponseElements)
	fmt.Println(string(res))

	exec.Command("git clone https://github.com/tfutils/tfenv.git ~/.tfenv").Run()
	exec.Command("export PATH=$HOME/.tfenv/bin:$PATH")

	return nil
}
