FROM public.ecr.aws/lambda/go:1

RUN yum install -y git
RUN curl -slO https://releases.hashicorp.com/terraform/1.1.3/terraform_1.1.3_linux_amd64.zip
RUN unzip terraform_1.1.3_linux_amd64.zip terraform -d /usr/local/bin

COPY tfrefresh ${LAMBDA_TASK_ROOT}

CMD [ "tfrefresh" ]
