BIN_PATH="bin/"
EXECUTABLE=game2048
WINDOWS=$(BIN_PATH)$(EXECUTABLE)_windows_amd64.exe
LINUX=$(BIN_PATH)$(EXECUTABLE)_linux_amd64
DARWIN=$(BIN_PATH)$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always --long --dirty)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)"  .

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)"  .

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)"  .

build: windows linux darwin ## Build binaries
	@echo version: $(VERSION)

benchmark:
	go test -bench=. ./...

test:
	go test ./...
