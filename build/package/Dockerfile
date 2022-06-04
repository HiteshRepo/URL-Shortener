FROM golang:1.18-alpine AS builder

# Install all required dependencies for building
RUN apk update
RUN apk add git make build-base

WORKDIR /service

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make compile

FROM alpine:3.14.2
COPY --from=builder ./service/build/out/url_shortener url_shortener
COPY --from=builder ./service/configs/default.yaml default.yaml
ENTRYPOINT ["./url_shortener", "-configFile", "default.yaml"]