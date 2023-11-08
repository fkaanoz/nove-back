FROM golang:1.20 AS BUILD
COPY . /var/app/shtil/

WORKDIR /var/app/shtil

RUN make tidy

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.ServiceVersion=1.1.0'" app/shtil/main.go

FROM alpine:3.18
COPY --from=BUILD /var/app/shtil/main /var/app/
