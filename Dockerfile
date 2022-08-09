FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY assets/ assets/
COPY templates/ templates/

RUN go build -o ./go-web-app

EXPOSE 8083

CMD [ "./go-web-app" ]