build:
	@go build -o bin/cache

run: build
	@./bin/cache

clean:
	@rm -rf bin/