package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os/exec"
	"strings"
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

	err := run("git clone --depth=1 https://github.com/mizzy/tfrefresh.git")
	if err != nil {
		return err
	}


	return nil
}

func run(cmd string) error {
	args := strings.Split(cmd, " ")
	err := exec.Command(args[0], args[1:]...).Run()
	return err
}
