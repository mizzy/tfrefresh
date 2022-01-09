.PHONY: build login deploy
tag = v0.0.3

build:
	GOOS=linux GOARCH=amd64 go build
	docker build . -t 887552608031.dkr.ecr.ap-northeast-1.amazonaws.com/tfrefresh:${tag}

login:
	aws ecr get-login-password | docker login --username AWS --password-stdin https://887552608031.dkr.ecr.ap-northeast-1.amazonaws.com

deploy: build login
	docker push 887552608031.dkr.ecr.ap-northeast-1.amazonaws.com/tfrefresh:${tag}
