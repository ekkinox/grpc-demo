FROM golang:1.18-alpine as build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /src
COPY . .
RUN go mod download && go build -o /gateway

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /gateway /gateway
ENTRYPOINT ["/gateway"]
