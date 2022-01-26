all: deps build start

deps:
	go mod tidy

build:
	go build .

start:
	./mmorpg-bot