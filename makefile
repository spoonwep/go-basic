build-mac-arm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build main.go
build-mac-amd:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
buildfile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
