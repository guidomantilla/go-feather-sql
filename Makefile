.PHONY: phony
phony-goal: ; @echo $@

validate: generate sort-import format vet lint coverage

generate:
	go generate ./...

sort-import:
	goimports-reviser -rm-unused -set-alias -format -recursive ./...

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test:
	go test -covermode count -coverprofile coverage.out.tmp.01 ./pkg/...
	cat coverage.out.tmp.01 | grep -v "mocks.go" > coverage.out

coverage: test
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

sonarqube: coverage
	sonar-scanner

update-dependencies:
	go get -u ./...
	go get -t -u ./...
	go mod tidy

prepare:
	go install -v github.com/incu6us/goimports-reviser/v3@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/golang/mock/mockgen@latest
	go mod download
	go mod tidy