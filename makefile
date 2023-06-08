build: 
	@go build -o bin/qascadeservice 
run: build
	@./bin/qascadeservice
fmt: 
	@gofmt -s -w .
tidy: 
	@go mod tidy