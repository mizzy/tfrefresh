FROM public.ecr.aws/lambda/go:1

RUN yum install -y git

RUN git clone https://github.com/tfutils/tfenv.git ~/.tfenv
ENV PATH ~/.tfenv/bin:$PATH

COPY tfrefresh ${LAMBDA_TASK_ROOT}

CMD [ "tfrefresh" ]
