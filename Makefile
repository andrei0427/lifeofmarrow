build: 
	go build -o ./bin/lifeofmarrow

run:
	air

css:
	yarn watch

test: 
	go test ./... -v