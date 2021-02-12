all: clean build

build: main.go
	go build -o bin/HelloGoService main.go

clean:
	rm -rf bin/HelloGoService/*
