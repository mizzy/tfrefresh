package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fujiwara/tfstate-lookup/tfstate"
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
	if e.Detail.EventName == "CreateLogStream" {
		return nil
	}

	fmt.Println(e.Detail.EventName)

	tmp, err := ioutil.TempDir("/tmp", "")
	defer os.RemoveAll(tmp)

	err = run(fmt.Sprintf("curl -sLO %s", os.Getenv("TF_BACKEND_URL")), tmp)
	if err != nil {
		return err
	}
	err = run("terraform init", tmp)
	if err != nil {
		return err
	}

	state, err := tfstateRead(fmt.Sprintf("%s/.terraform/terraform.tfstate", tmp))
	if err != nil {
		return err
	}

	if e.Detail.EventName == "PutRetentionPolicy" {
		refreshLogGroup(e.Detail.RequestParameters["logGroupName"].(string), state, tmp)
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

func tfstateRead(stateLoc string) (*tfstate.TFState, error) {
	state, err := tfstate.ReadURL(stateLoc)
	return state, err
}

func refreshLogGroup(logGroup string, state *tfstate.TFState, dir string) error {
	list, err := state.List()
	if err != nil {
		return nil
	}

	for _, resource := range list {
		if strings.HasPrefix(resource, "aws_cloudwatch_log_group") {
			id, err := state.Lookup(fmt.Sprintf("%s.id", resource))
			if err != nil {
				return err
			}
			if id.String() == logGroup {
				run(fmt.Sprintf("terraform refresh -target %s", resource), dir)
			}
		}
	}

	return nil
}
