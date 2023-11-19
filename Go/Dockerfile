##
## Build
##

FROM golang:1.19-buster AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/bin

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /set-list-play /set-list-play

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/set-list-play"]
