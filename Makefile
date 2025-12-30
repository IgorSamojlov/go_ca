.PHONY: run vet
run:
	air --build.cmd "go build -C cmd/http -o ../../build/http" --build.entrypoint "./build/http"
vet:
	go vet ./...
# lint:
