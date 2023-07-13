FROM public.ecr.aws/bitnami/golang:1.19.11-debian-11-r0 AS builder
COPY . /src
WORKDIR /src
RUN GOOS=linux GOARCH=amd64 go build -o main main.go

FROM public.ecr.aws/lambda/go:1
COPY --from=builder /src/main ${LAMBDA_TASK_ROOT}
CMD [ "main" ] 