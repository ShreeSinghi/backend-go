
all: 
	go build -o mvc ./cmd/main.go
	./mvc
	
clean:
	rm mvc

setup:
	chmod 777 ./scripts/setup.sh
	./scripts/setup.sh
