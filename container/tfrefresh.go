package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os/exec"
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

	err := run("rm -rf /tmp/tfrefresh", "/tmp")
	if err != nil {
		return err
	}
	err = run("git clone --depth=1 https://github.com/mizzy/tfrefresh.git", "/tmp")
	if err != nil {
		return err
	}
	err = run("terraform init", "/tmp/tfrefresh/terraform")
	if err != nil {
		return err
	}

	err = run("terraform refresh -lock-timeout 10m", "/tmp/tfrefresh/terraform")
	if err != nil {
		return err
	}

	return nil
}

func run(c, dir string) error {
	cmd := exec.Command("/bin/sh", "-c", c)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Dir = dir
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	fmt.Println(cmd)
	fmt.Printf("stdout: %s\n", stdout.String())

	if err != nil {
		fmt.Printf("stder: %s\n", stderr.String())
		return err
	}

	return nil
}
