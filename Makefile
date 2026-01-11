DEFAULT_GOAL: build

build: tidy
	go build -o build/3d-render-engine main.go
tidy:
	go mod tidy
clean:
	rm build/*
