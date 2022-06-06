<!--
 * @Author: ChenLong longchen2008@126.com
 * @Date: 2022-06-05 10:07:03
 * @LastEditors: ChenLong longchen2008@126.com
 * @LastEditTime: 2022-06-05 11:43:26
 * @FilePath: \study\03_golang_tech_stack\special\cross\readme.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# Golang gin nginx Centos部署

将golang web应用，使用nginx部署到linux服务器下面，实现步骤如下：

## 创建gin测试web项目

```bash
go get -u github.com/gin-gonic/gin
```

```bash
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run()
}
```

## 交叉编译

### 1、查看自己的go env

在命令行输入：go env 结果如下

```bash
set GO111MODULE=on
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\52406\AppData\Local\go-build
set GOENV=C:\Users\52406\AppData\Roaming\go\env
set GOEXE=.exe
set GOEXPERIMENT=
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=C:\Users\52406\go\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=C:\Users\52406\go
set GOPRIVATE=
set GOPROXY=http://goproxy.cn
set GOROOT=C:\Program Files\Go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=C:\Program Files\Go\pkg\tool\windows_amd64
set GOVCS=
set GOVERSION=go1.17.3
set GCCGO=gccgo
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=NUL
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\52406\AppData\Local\Temp\go-build498288655=/tmp/go-build -gno-record-gcc-switches
```

### 2、修改环境变量

```bash
$env:GOOS = "linux"
$env:GOARCH = "amd64" 
```

### 3、重新查看环境变量，保存环境变量被修改

在命令行输入：go env 结果如下

```bash
set GO111MODULE=on
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\52406\AppData\Local\go-build
set GOENV=C:\Users\52406\AppData\Roaming\go\env
set GOEXE=
set GOEXPERIMENT=
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=C:\Users\52406\go\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=linux
set GOPATH=C:\Users\52406\go
set GOPRIVATE=
set GOPROXY=http://goproxy.cn
set GOROOT=C:\Program Files\Go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=C:\Program Files\Go\pkg\tool\windows_amd64
set GOVCS=
set GOVERSION=go1.17.3
set GCCGO=gccgo
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=0
set GOMOD=C:\Users\52406\Desktop\golang\golangprojects\duoke360.com\golangweb\go.mod
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-fPIC -m64 -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\52406\AppData\Local\Temp\go-build891061685=/tmp/go-build -gno-record-gcc-switches
```

### 4、使用go build 编译demo.go

```bash
go build main.go
```

会生成一个main可执行文件

### 5、将demo文件上传到linux环境下

### 6、修改demo的执行权限

```bash
chmod +x ./main
```

### 7、运行程序

```bash
nohup ./main &
```

## 配置nginx

```bash
server {
    listen 80;
    server_name  192.168.18.128;

    location / {
        proxy_pass http://localhost:8080;
    }
}
```

## 访问

```bash
http://192.168.18.128/ping
```

### 运行结果

```bash
pong
```
