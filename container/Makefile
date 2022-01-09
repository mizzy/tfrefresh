.PHONY: build login deploy

login:
	aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
	aws ecr get-login-password | docker login --username AWS --password-stdin https://887552608031.dkr.ecr.ap-northeast-1.amazonaws.com

build: login
	GOOS=linux GOARCH=amd64 go build
	docker build . -t 887552608031.dkr.ecr.ap-northeast-1.amazonaws.com/tfrefresh:${TAG}

deploy: build
	docker push 887552608031.dkr.ecr.ap-northeast-1.amazonaws.com/tfrefresh:${TAG}

