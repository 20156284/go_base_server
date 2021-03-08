BINARY="go_base_server"
GBS = "gbs"

all: check gbs run

gbs:
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.io,direct
	go build -o ${GBS} cmd/main.go
	@if [ -f ${GBS} ] ; then ./${GBS} initdb && rm ${GBS} ; fi

linux-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

windows-build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe

mac-build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BINARY}

run:
	@go run main.go

swagger:
	@gf swagger

check:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	@if [ -f ${GBS} ] ; then rm ${GBS} ; fi

help:
	@echo "make - 构建gfva终端工具并初始化数据,初始化数据后删除gfva终端工具,启动server项目"
	@echo "make gbs - 构建gfva终端工具 并初始化数据"
	@echo "make linux-build - 编译 Go 代码, 生成Linux系统的二进制文件"
	@echo "make windows-build - 编译 Go 代码, 生成Windows系统的exe文件"
	@echo "make mac-build - 编译 Go 代码, 生成Mac系统的二进制文件"
	@echo "make run - 直接运行 main.go"
	@echo "make clean - 移除二进制文件"
	@echo "make check - 运行 Go 工具 'fmt' and 'vet'"