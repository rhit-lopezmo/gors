DEFAULT_GOAL: build

build: tidy
	go build -o build/game main.go
run:
	./build/game
tidy:
	go mod tidy
clean:
	rm build/*
