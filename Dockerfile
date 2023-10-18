FROM --platform=$BUILDPLATFORM golang:1.21 as build
WORKDIR /src
COPY . .
RUN go mod download
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o cli-template

FROM alpine:3.18
COPY --from=build /src/cli-template /bin/cli-template
ENTRYPOINT ["cli-template"]
