BINARY=retroflag-powerd

build:
	go build -o $(BINARY) ./cmd/retroflag-powerd

run:
	go run ./cmd/retroflag-powerd

clean:
	rm -f $(BINARY)
