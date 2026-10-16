FROM golang:1.27.2

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

RUN go test -v

CMD ["/docker-gs-ping"]