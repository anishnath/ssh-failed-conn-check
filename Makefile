PROGRAM_NAME=ssh-failed-connections
GO_ARCH=amd64
build:
	GOARCH=$(GO_ARCH)  go build -o $(PROGRAM_NAME)

clean:
	rm -f $(PROGRAM_NAME)

.PHONY: build clean
