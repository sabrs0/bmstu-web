FROM golang:1.21.1-alpine as builder

RUN apk --no-cache add gcc g++ make ca-certificates
RUN mkdir -p  /go/src/app

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/cmd/foundationer/main.go

FROM ubuntu:22.10

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

COPY --from=builder /go/src/app .

CMD [ "./main" ]
