FROM public.ecr.aws/docker/library/golang:1.20 AS build
COPY . /src
WORKDIR /src
RUN make -s build

FROM build AS stage
WORKDIR /src
RUN make -s ci

FROM public.ecr.aws/lambda/go:1
COPY --from=build /src/example-api ${LAMBDA_TASK_ROOT}
CMD [ "example-api" ] 