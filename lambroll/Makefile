.PHONY: build deploy

build:
	GOOS=linux GOARCH=amd64 go build

deploy: build
	lambroll --tfstate ../terraform/.terraform/terraform.tfstate deploy
