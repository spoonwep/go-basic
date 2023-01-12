# 编译可执行文件
build-mac-arm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build main.go
build-mac-amd:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
buildfile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
# 数据库迁移工具
goose-create:
	cd migrations && goose sqlite3 ../assets/sqlite/sqlite.db create $(name) sql
goose-up:
	cd migrations && goose sqlite3 ../assets/sqlite/sqlite.db up
goose-down:
	cd migrations && goose sqlite3 ../assets/sqlite/sqlite.db down
# golangci-lint
check-install:
	@golangci-lint --version >/dev/null 2>&1 || { \
		echo "请先安装golangci-lint，如果用goland，可以在Preferences->Tools->File Watcher里面添加，会自动下载。否则可以点这里自行安装：https://golangci-lint.run/usage/install/#local-installation"; \
	}
lint: check-install
	golangci-lint run
# 单元测试
test:
	$(GOCMD) test -cover -race ./...