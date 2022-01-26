FROM golang:1.16-alpine

WORKDIR /go/src/mmorpg
COPY . .

RUN go get -d -v ./...
# RUN go install -v ./...
RUN go build .

CMD ["./mmorpg-bot"]