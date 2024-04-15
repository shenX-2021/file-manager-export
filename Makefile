GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
CC = $(shell go env CC)

# 构建windows可执行文件
build-windows:
	go env -w GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc \
	&& go build -ldflags "-H=windowsgui" -o fm-export.exe main.go \
	&& go env -w GOOS=$(GOOS) GOARCH=$(GOARCH) CC=$(CC)