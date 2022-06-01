FROM golang:1.18.2-bullseye

WORKDIR /go/src/app

COPY . .

RUN go mod init

RUN go build -o /go/bin/app .

EXPOSE 8080

CMD [ "/go/bin/app" ]

