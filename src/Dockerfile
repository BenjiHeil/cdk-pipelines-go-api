FROM public.ecr.aws/docker/library/golang:1.20 AS builder
COPY . /src
WORKDIR /src
RUN make build
RUN make test && make coverage && make benchmark

FROM public.ecr.aws/lambda/go:1
COPY --from=builder /src/example-api ${LAMBDA_TASK_ROOT}
CMD [ "example-api" ] 