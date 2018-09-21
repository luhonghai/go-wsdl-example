FROM golang:1.11-alpine3.7 AS build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on
RUN apk add --no-cache \
    git \
    musl-dev

WORKDIR /go/src/github.com/luhonghai/wsdl-example
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go install -ldflags '-s -w'

# ---

FROM scratch
COPY --from=build /go/bin/wsdl-example /wsdl-example
ENTRYPOINT [ "/wsdl-example" ]
CMD [ "--help" ]