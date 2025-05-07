tidy:
	@go mod tidy
	@go mod vendor

run:
	@go run cmd/main.go

build:
	@go build -o bin/main.exe cmd/main.go

git:
	@git add .
	@git commit -m "Changed"
	@git push

build:
	@go build -o bin/main.exe cmd/main.go