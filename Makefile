run: build
	@./bin/lifeofmarrow

templ:
	@templ generate -watch -proxy=http://localhost:8080

tailwind:
	@tailwindcss build -o static/css/tailwind.css --watch

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss @tailwindcss/aspect-ratio @tailwindcss/forms @tailwindcss/typography

build:
	npm run tailwindcss 
	@templ generate view
	@go build -o bin/lifeofmarrow main.go

test: 
	go test ./...