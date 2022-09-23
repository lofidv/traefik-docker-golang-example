FROM golang:alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./run .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/app/run .
COPY --from=builder /go/src/app/.env .

RUN chmod +x ./run

EXPOSE 8060

ENTRYPOINT [ "./run" ]
